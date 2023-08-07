pragma solidity 0.8.6;

import "@openzeppelin/contracts/security/Pausable.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import "../../vendor/@arbitrum/nitro-contracts/src/precompiles/ArbGasInfo.sol";
import "../../vendor/@eth-optimism/contracts/0.8.6/contracts/L2/predeploys/OVM_GasPriceOracle.sol";
import "../ExecutionPrevention.sol";
import {Config, State, Upkeep} from "../../interfaces/automation/1_3/AutomationRegistryInterface1_3.sol";
import "../../ConfirmedOwner.sol";
import "../../interfaces/AggregatorV3Interface.sol";
import "../../interfaces/PliTokenInterface.sol";
import "../../interfaces/automation/KeeperCompatibleInterface.sol";
import "../../interfaces/automation/UpkeepTranscoderInterface.sol";

/**
 * @notice Base Keeper Registry contract, contains shared logic between
 * KeeperRegistry and KeeperRegistryLogic
 */
abstract contract KeeperRegistryBase1_3 is ConfirmedOwner, ExecutionPrevention, ReentrancyGuard, Pausable {
  address internal constant ZERO_ADDRESS = address(0);
  address internal constant IGNORE_ADDRESS = 0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF;
  bytes4 internal constant CHECK_SELECTOR = KeeperCompatibleInterface.checkUpkeep.selector;
  bytes4 internal constant PERFORM_SELECTOR = KeeperCompatibleInterface.performUpkeep.selector;
  uint256 internal constant PERFORM_GAS_MIN = 2_300;
  uint256 internal constant CANCELLATION_DELAY = 50;
  uint256 internal constant PERFORM_GAS_CUSHION = 5_000;
  uint256 internal constant PPB_BASE = 1_000_000_000;
  uint32 internal constant UINT32_MAX = type(uint32).max;
  uint96 internal constant PLI_TOTAL_SUPPLY = 1e27;
  UpkeepFormat internal constant UPKEEP_TRANSCODER_VERSION_BASE = UpkeepFormat.V2;
  // L1_FEE_DATA_PADDING includes 35 bytes for L1 data padding for Optimism
  bytes internal constant L1_FEE_DATA_PADDING =
    "0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff";
  // MAX_INPUT_DATA represents the estimated max size of the sum of L1 data padding and msg.data in performUpkeep
  // function, which includes 4 bytes for function selector, 32 bytes for upkeep id, 35 bytes for data padding, and
  // 64 bytes for estimated perform data
  bytes internal constant MAX_INPUT_DATA =
    "0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff";

  address[] internal s_keeperList;
  EnumerableSet.UintSet internal s_upkeepIDs;
  mapping(uint256 => Upkeep) internal s_upkeep;
  mapping(address => KeeperInfo) internal s_keeperInfo;
  mapping(address => address) internal s_proposedPayee;
  mapping(uint256 => address) internal s_proposedAdmin;
  mapping(uint256 => bytes) internal s_checkData;
  mapping(address => MigrationPermission) internal s_peerRegistryMigrationPermission;
  Storage internal s_storage;
  uint256 internal s_fallbackGasPrice; // not in config object for gas savings
  uint256 internal s_fallbackPliPrice; // not in config object for gas savings
  uint96 internal s_ownerPliBalance;
  uint256 internal s_expectedPliBalance;
  address internal s_transcoder;
  address internal s_registrar;

  PliTokenInterface public immutable PLI;
  AggregatorV3Interface public immutable PLI_ETH_FEED;
  AggregatorV3Interface public immutable FAST_GAS_FEED;
  OVM_GasPriceOracle public immutable OPTIMISM_ORACLE = OVM_GasPriceOracle(0x420000000000000000000000000000000000000F);
  ArbGasInfo public immutable ARB_NITRO_ORACLE = ArbGasInfo(0x000000000000000000000000000000000000006C);
  PaymentModel public immutable PAYMENT_MODEL;
  uint256 public immutable REGISTRY_GAS_OVERHEAD;

  error ArrayHasNoEntries();
  error CannotCancel();
  error DuplicateEntry();
  error EmptyAddress();
  error GasLimitCanOnlyIncrease();
  error GasLimitOutsideRange();
  error IndexOutOfRange();
  error InsufficientFunds();
  error InvalidDataLength();
  error InvalidPayee();
  error InvalidRecipient();
  error KeepersMustTakeTurns();
  error MigrationNotPermitted();
  error NotAContract();
  error OnlyActiveKeepers();
  error OnlyCallableByAdmin();
  error OnlyCallableByPLIToken();
  error OnlyCallableByOwnerOrAdmin();
  error OnlyCallableByOwnerOrRegistrar();
  error OnlyCallableByPayee();
  error OnlyCallableByProposedAdmin();
  error OnlyCallableByProposedPayee();
  error OnlyPausedUpkeep();
  error OnlyUnpausedUpkeep();
  error ParameterLengthError();
  error PaymentGreaterThanAllPLI();
  error TargetCheckReverted(bytes reason);
  error TranscoderNotSet();
  error UpkeepCancelled();
  error UpkeepNotCanceled();
  error UpkeepNotNeeded();
  error ValueNotChanged();

  enum MigrationPermission {
    NONE,
    OUTGOING,
    INCOMING,
    BIDIRECTIONAL
  }

  enum PaymentModel {
    DEFAULT,
    ARBITRUM,
    OPTIMISM
  }

  /**
   * @notice storage of the registry, contains a mix of config and state data
   */
  struct Storage {
    uint32 paymentPremiumPPB;
    uint32 flatFeeMicroPli;
    uint24 blockCountPerTurn;
    uint32 checkGasLimit;
    uint24 stalenessSeconds;
    uint16 gasCeilingMultiplier;
    uint96 minUpkeepSpend; // 1 full evm word
    uint32 maxPerformGas;
    uint32 nonce;
  }

  struct KeeperInfo {
    address payee;
    uint96 balance;
    bool active;
  }

  struct PerformParams {
    address from;
    uint256 id;
    bytes performData;
    uint256 maxPliPayment;
    uint256 gasLimit;
    uint256 fastGasWei;
    uint256 pliEth;
  }

  event ConfigSet(Config config);
  event FundsAdded(uint256 indexed id, address indexed from, uint96 amount);
  event FundsWithdrawn(uint256 indexed id, uint256 amount, address to);
  event KeepersUpdated(address[] keepers, address[] payees);
  event OwnerFundsWithdrawn(uint96 amount);
  event PayeeshipTransferRequested(address indexed keeper, address indexed from, address indexed to);
  event PayeeshipTransferred(address indexed keeper, address indexed from, address indexed to);
  event PaymentWithdrawn(address indexed keeper, uint256 indexed amount, address indexed to, address payee);
  event UpkeepAdminTransferRequested(uint256 indexed id, address indexed from, address indexed to);
  event UpkeepAdminTransferred(uint256 indexed id, address indexed from, address indexed to);
  event UpkeepCanceled(uint256 indexed id, uint64 indexed atBlockHeight);
  event UpkeepCheckDataUpdated(uint256 indexed id, bytes newCheckData);
  event UpkeepGasLimitSet(uint256 indexed id, uint96 gasLimit);
  event UpkeepMigrated(uint256 indexed id, uint256 remainingBalance, address destination);
  event UpkeepPaused(uint256 indexed id);
  event UpkeepPerformed(
    uint256 indexed id,
    bool indexed success,
    address indexed from,
    uint96 payment,
    bytes performData
  );
  event UpkeepReceived(uint256 indexed id, uint256 startingBalance, address importedFrom);
  event UpkeepUnpaused(uint256 indexed id);
  event UpkeepRegistered(uint256 indexed id, uint32 executeGas, address admin);

  /**
   * @param paymentModel the payment model of default, Arbitrum, or Optimism
   * @param registryGasOverhead the gas overhead used by registry in performUpkeep
   * @param pli address of the PLI Token
   * @param pliEthFeed address of the PLI/ETH price feed
   * @param fastGasFeed address of the Fast Gas price feed
   */
  constructor(
    PaymentModel paymentModel,
    uint256 registryGasOverhead,
    address pli,
    address pliEthFeed,
    address fastGasFeed
  ) ConfirmedOwner(msg.sender) {
    PAYMENT_MODEL = paymentModel;
    REGISTRY_GAS_OVERHEAD = registryGasOverhead;
    if (ZERO_ADDRESS == pli || ZERO_ADDRESS == pliEthFeed || ZERO_ADDRESS == fastGasFeed) {
      revert EmptyAddress();
    }
    PLI = PliTokenInterface(pli);
    PLI_ETH_FEED = AggregatorV3Interface(pliEthFeed);
    FAST_GAS_FEED = AggregatorV3Interface(fastGasFeed);
  }

  /**
   * @dev retrieves feed data for fast gas/eth and pli/eth prices. if the feed
   * data is stale it uses the configured fallback price. Once a price is picked
   * for gas it takes the min of gas price in the transaction or the fast gas
   * price in order to reduce costs for the upkeep clients.
   */
  function _getFeedData() internal view returns (uint256 gasWei, uint256 pliEth) {
    uint32 stalenessSeconds = s_storage.stalenessSeconds;
    bool staleFallback = stalenessSeconds > 0;
    uint256 timestamp;
    int256 feedValue;
    (, feedValue, , timestamp, ) = FAST_GAS_FEED.latestRoundData();
    if ((staleFallback && stalenessSeconds < block.timestamp - timestamp) || feedValue <= 0) {
      gasWei = s_fallbackGasPrice;
    } else {
      gasWei = uint256(feedValue);
    }
    (, feedValue, , timestamp, ) = PLI_ETH_FEED.latestRoundData();
    if ((staleFallback && stalenessSeconds < block.timestamp - timestamp) || feedValue <= 0) {
      pliEth = s_fallbackPliPrice;
    } else {
      pliEth = uint256(feedValue);
    }
    return (gasWei, pliEth);
  }

  /**
   * @dev calculates PLI paid for gas spent plus a configure premium percentage
   * @param gasLimit the amount of gas used
   * @param fastGasWei the fast gas price
   * @param pliEth the exchange ratio between PLI and ETH
   * @param isExecution if this is triggered by a perform upkeep function
   */
  function _calculatePaymentAmount(
    uint256 gasLimit,
    uint256 fastGasWei,
    uint256 pliEth,
    bool isExecution
  ) internal view returns (uint96 payment) {
    Storage memory store = s_storage;
    uint256 gasWei = fastGasWei * store.gasCeilingMultiplier;
    // in case it's actual execution use actual gas price, capped by fastGasWei * gasCeilingMultiplier
    if (isExecution && tx.gasprice < gasWei) {
      gasWei = tx.gasprice;
    }

    uint256 weiForGas = gasWei * (gasLimit + REGISTRY_GAS_OVERHEAD);
    uint256 premium = PPB_BASE + store.paymentPremiumPPB;
    uint256 l1CostWei = 0;
    if (PAYMENT_MODEL == PaymentModel.OPTIMISM) {
      bytes memory txCallData = new bytes(0);
      if (isExecution) {
        txCallData = bytes.concat(msg.data, L1_FEE_DATA_PADDING);
      } else {
        txCallData = MAX_INPUT_DATA;
      }
      l1CostWei = OPTIMISM_ORACLE.getL1Fee(txCallData);
    } else if (PAYMENT_MODEL == PaymentModel.ARBITRUM) {
      l1CostWei = ARB_NITRO_ORACLE.getCurrentTxL1GasFees();
    }
    // if it's not performing upkeeps, use gas ceiling multiplier to estimate the upper bound
    if (!isExecution) {
      l1CostWei = store.gasCeilingMultiplier * l1CostWei;
    }

    uint256 total = ((weiForGas + l1CostWei) * 1e9 * premium) / pliEth + uint256(store.flatFeeMicroPli) * 1e12;
    if (total > PLI_TOTAL_SUPPLY) revert PaymentGreaterThanAllPLI();
    return uint96(total); // PLI_TOTAL_SUPPLY < UINT96_MAX
  }

  /**
   * @dev ensures all required checks are passed before an upkeep is performed
   */
  function _prePerformUpkeep(
    Upkeep memory upkeep,
    address from,
    uint256 maxPliPayment
  ) internal view {
    if (upkeep.paused) revert OnlyUnpausedUpkeep();
    if (!s_keeperInfo[from].active) revert OnlyActiveKeepers();
    if (upkeep.balance < maxPliPayment) revert InsufficientFunds();
    if (upkeep.lastKeeper == from) revert KeepersMustTakeTurns();
  }

  /**
   * @dev ensures the upkeep is not cancelled and the caller is the upkeep admin
   */
  function requireAdminAndNotCancelled(Upkeep memory upkeep) internal view {
    if (msg.sender != upkeep.admin) revert OnlyCallableByAdmin();
    if (upkeep.maxValidBlocknumber != UINT32_MAX) revert UpkeepCancelled();
  }

  /**
   * @dev generates a PerformParams struct for use in _performUpkeepWithParams()
   */
  function _generatePerformParams(
    address from,
    uint256 id,
    bytes memory performData,
    bool isExecution
  ) internal view returns (PerformParams memory) {
    uint256 gasLimit = s_upkeep[id].executeGas;
    (uint256 fastGasWei, uint256 pliEth) = _getFeedData();
    uint96 maxPliPayment = _calculatePaymentAmount(gasLimit, fastGasWei, pliEth, isExecution);

    return
      PerformParams({
        from: from,
        id: id,
        performData: performData,
        maxPliPayment: maxPliPayment,
        gasLimit: gasLimit,
        fastGasWei: fastGasWei,
        pliEth: pliEth
      });
  }
}
