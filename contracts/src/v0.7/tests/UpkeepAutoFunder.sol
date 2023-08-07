// SPDX-License-Identifier: MIT
pragma solidity ^0.7.0;

import "../KeeperCompatible.sol";
import "../interfaces/PliTokenInterface.sol";
import "../interfaces/KeeperRegistryInterface.sol";
import "../ConfirmedOwner.sol";

contract UpkeepAutoFunder is KeeperCompatible, ConfirmedOwner {
  bool public s_isEligible;
  bool public s_shouldCancel;
  uint256 public s_upkeepId;
  uint96 public s_autoFundPli;
  PliTokenInterface public immutable PLI;
  KeeperRegistryBaseInterface public immutable s_keeperRegistry;

  constructor(address pliAddress, address registryAddress) ConfirmedOwner(msg.sender) {
    PLI = PliTokenInterface(pliAddress);
    s_keeperRegistry = KeeperRegistryBaseInterface(registryAddress);

    s_isEligible = false;
    s_shouldCancel = false;
    s_upkeepId = 0;
    s_autoFundPli = 0;
  }

  function setShouldCancel(bool value) external onlyOwner {
    s_shouldCancel = value;
  }

  function setIsEligible(bool value) external onlyOwner {
    s_isEligible = value;
  }

  function setAutoFundPli(uint96 value) external onlyOwner {
    s_autoFundPli = value;
  }

  function setUpkeepId(uint256 value) external onlyOwner {
    s_upkeepId = value;
  }

  function checkUpkeep(bytes calldata data)
    external
    override
    cannotExecute
    returns (bool callable, bytes calldata executedata)
  {
    return (s_isEligible, data);
  }

  function performUpkeep(bytes calldata data) external override {
    require(s_isEligible, "Upkeep should be eligible");
    s_isEligible = false; // Allow upkeep only once until it is set again

    // Topup upkeep so it can be called again
    PLI.transferAndCall(address(s_keeperRegistry), s_autoFundPli, abi.encode(s_upkeepId));

    if (s_shouldCancel) {
      s_keeperRegistry.cancelUpkeep(s_upkeepId);
    }
  }
}
