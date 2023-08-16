// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import {FunctionsBillingRegistry} from "../dev/functions/FunctionsBillingRegistry.sol";

contract FunctionsBillingRegistryWithInit is FunctionsBillingRegistry {
  constructor(
    address pli,
    address pliEthFeed,
    address oracle
  ) {
    initialize(pli, pliEthFeed, oracle);
  }
}
