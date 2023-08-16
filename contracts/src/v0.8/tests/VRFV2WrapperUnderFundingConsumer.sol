// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../ConfirmedOwner.sol";
import "../interfaces/PliTokenInterface.sol";
import "../interfaces/VRFV2WrapperInterface.sol";

contract VRFV2WrapperUnderFundingConsumer is ConfirmedOwner {
  PliTokenInterface internal immutable PLI;
  VRFV2WrapperInterface internal immutable VRF_V2_WRAPPER;

  constructor(address _pli, address _vrfV2Wrapper) ConfirmedOwner(msg.sender) {
    PLI = PliTokenInterface(_pli);
    VRF_V2_WRAPPER = VRFV2WrapperInterface(_vrfV2Wrapper);
  }

  function makeRequest(
    uint32 _callbackGasLimit,
    uint16 _requestConfirmations,
    uint32 _numWords
  ) external onlyOwner {
    PLI.transferAndCall(
      address(VRF_V2_WRAPPER),
      // Pay less than the needed amount
      VRF_V2_WRAPPER.calculateRequestPrice(_callbackGasLimit) - 1,
      abi.encode(_callbackGasLimit, _requestConfirmations, _numWords)
    );
  }
}
