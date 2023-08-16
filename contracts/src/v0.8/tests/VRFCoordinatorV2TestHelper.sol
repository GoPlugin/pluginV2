// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../VRFCoordinatorV2.sol";

contract VRFCoordinatorV2TestHelper is VRFCoordinatorV2 {
  uint96 s_paymentAmount;
  uint256 s_gasStart;

  constructor(
    address pli,
    address blockhashStore,
    address pliEthFeed
  )
    // solhint-disable-next-line no-empty-blocks
    VRFCoordinatorV2(pli, blockhashStore, pliEthFeed)
  {
    /* empty */
  }

  function calculatePaymentAmountTest(
    uint256 gasAfterPaymentCalculation,
    uint32 fulfillmentFlatFeePliPPM,
    uint256 weiPerUnitGas
  ) external {
    s_paymentAmount = calculatePaymentAmount(
      gasleft(),
      gasAfterPaymentCalculation,
      fulfillmentFlatFeePliPPM,
      weiPerUnitGas
    );
  }

  function getPaymentAmount() public view returns (uint96) {
    return s_paymentAmount;
  }

  function getGasStart() public view returns (uint256) {
    return s_gasStart;
  }
}
