// SPDX-License-Identifier: MIT

pragma solidity 0.8.6;

import "../dev/VRFSubscriptionBalanceMonitor.sol";

contract VRFSubscriptionBalanceMonitorExposed is VRFSubscriptionBalanceMonitor {
  constructor(
    address pliTokenAddress,
    address coordinatorAddress,
    address keeperRegistryAddress,
    uint256 minWaitPeriodSeconds
  ) VRFSubscriptionBalanceMonitor(pliTokenAddress, coordinatorAddress, keeperRegistryAddress, minWaitPeriodSeconds) {}

  function setLastTopUpXXXTestOnly(uint64 target, uint56 lastTopUpTimestamp) external {
    s_targets[target].lastTopUpTimestamp = lastTopUpTimestamp;
  }
}
