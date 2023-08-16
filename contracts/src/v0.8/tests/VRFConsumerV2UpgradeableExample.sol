// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../interfaces/PliTokenInterface.sol";
import "../interfaces/VRFCoordinatorV2Interface.sol";
import "../dev/VRFConsumerBaseV2Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable-4.7.3/proxy/utils/Initializable.sol";

contract VRFConsumerV2UpgradeableExample is Initializable, VRFConsumerBaseV2Upgradeable {
  uint256[] public s_randomWords;
  uint256 public s_requestId;
  VRFCoordinatorV2Interface public COORDINATOR;
  PliTokenInterface public PLITOKEN;
  uint64 public s_subId;
  uint256 public s_gasAvailable;

  function initialize(address _vrfCoordinator, address _pli) public initializer {
    __VRFConsumerBaseV2_init(_vrfCoordinator);
    COORDINATOR = VRFCoordinatorV2Interface(_vrfCoordinator);
    PLITOKEN = PliTokenInterface(_pli);
  }

  function fulfillRandomWords(uint256 requestId, uint256[] memory randomWords) internal override {
    require(requestId == s_requestId, "request ID is incorrect");

    s_gasAvailable = gasleft();
    s_randomWords = randomWords;
  }

  function createSubscriptionAndFund(uint96 amount) external {
    if (s_subId == 0) {
      s_subId = COORDINATOR.createSubscription();
      COORDINATOR.addConsumer(s_subId, address(this));
    }
    // Approve the pli transfer.
    PLITOKEN.transferAndCall(address(COORDINATOR), amount, abi.encode(s_subId));
  }

  function topUpSubscription(uint96 amount) external {
    require(s_subId != 0, "sub not set");
    // Approve the pli transfer.
    PLITOKEN.transferAndCall(address(COORDINATOR), amount, abi.encode(s_subId));
  }

  function updateSubscription(address[] memory consumers) external {
    require(s_subId != 0, "subID not set");
    for (uint256 i = 0; i < consumers.length; i++) {
      COORDINATOR.addConsumer(s_subId, consumers[i]);
    }
  }

  function requestRandomness(
    bytes32 keyHash,
    uint64 subId,
    uint16 minReqConfs,
    uint32 callbackGasLimit,
    uint32 numWords
  ) external returns (uint256) {
    s_requestId = COORDINATOR.requestRandomWords(keyHash, subId, minReqConfs, callbackGasLimit, numWords);
    return s_requestId;
  }
}
