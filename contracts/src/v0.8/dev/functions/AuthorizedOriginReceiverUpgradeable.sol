// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import {EnumerableSet} from "../vendor/openzeppelin-solidity/v.4.8.0/contracts/utils/structs/EnumerableSet.sol";
import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

/**
 * @notice Modified AuthorizedReciever abstract for use on the FunctionsOracle contract to limit usage
 * @notice Uses tx.origin instead of msg.sender because the client contract sends messages to the Oracle contract
 */

abstract contract AuthorizedOriginReceiverUpgradeable is Initializable {
  using EnumerableSet for EnumerableSet.AddressSet;

  event AuthorizedSendersChanged(address[] senders, address changedBy);
  event AuthorizedSendersActive(address account);
  event AuthorizedSendersDeactive(address account);

  error EmptySendersList();
  error UnauthorizedSender();
  error NotAllowedToSetSenders();
  error AlreadySet();

  bool private s_active;
  EnumerableSet.AddressSet private s_authorizedSenders;
  address[] private s_authorizedSendersList;

  /**
   * @dev Initializes the contract in active state.
   */
  function __AuthorizedOriginReceiver_initialize(bool active) internal onlyInitializing {
    s_active = active;
  }

  /**
   * @dev Returns true if the contract is paused, and false otherwise.
   */
  function authorizedReceiverActive() public view virtual returns (bool) {
    return s_active;
  }

  /**
   * @dev Triggers AuthorizedOriginReceiver usage to block unuthorized senders.
   *
   * Requirements:
   *
   * - The contract must not be deactive.
   */
  function activateAuthorizedReceiver() external validateAuthorizedSenderSetter {
    if (authorizedReceiverActive()) {
      revert AlreadySet();
    }
    s_active = true;
    emit AuthorizedSendersActive(msg.sender);
  }

  /**
   * @dev Triggers AuthorizedOriginReceiver usage to allow all senders.
   *
   * Requirements:
   *
   * - The contract must be active.
   */
  function deactivateAuthorizedReceiver() external validateAuthorizedSenderSetter {
    if (!authorizedReceiverActive()) {
      revert AlreadySet();
    }
    s_active = false;
    emit AuthorizedSendersDeactive(msg.sender);
  }

  /**
   * @notice Sets the permission to request for the given wallet(s).
   * @param senders The addresses of the wallet addresses to grant access
   */
  function addAuthorizedSenders(address[] calldata senders) external validateAuthorizedSenderSetter {
    if (senders.length == 0) {
      revert EmptySendersList();
    }
    for (uint256 i = 0; i < senders.length; i++) {
      bool success = s_authorizedSenders.add(senders[i]);
      if (success) {
        s_authorizedSendersList.push(senders[i]);
      }
    }
    emit AuthorizedSendersChanged(senders, msg.sender);
  }

  /**
   * @notice Remove the permission to request for the given wallet(s).
   * @param senders The addresses of the wallet addresses to revoke access
   */
  function removeAuthorizedSenders(address[] calldata senders) external validateAuthorizedSenderSetter {
    if (senders.length == 0) {
      revert EmptySendersList();
    }
    for (uint256 i = 0; i < senders.length; i++) {
      bool success = s_authorizedSenders.remove(senders[i]);
      if (success) {
        // Remove from s_authorizedSendersList
        for (uint256 j = 0; j < s_authorizedSendersList.length; j++) {
          if (s_authorizedSendersList[j] == senders[i]) {
            address last = s_authorizedSendersList[s_authorizedSendersList.length - 1];
            // Copy last element and overwrite senders[i] to be deleted with it
            s_authorizedSendersList[i] = last;
            s_authorizedSendersList.pop();
          }
        }
      }
    }
    emit AuthorizedSendersChanged(senders, msg.sender);
  }

  /**
   * @notice Retrieve a list of authorized senders
   * @return array of addresses
   */
  function getAuthorizedSenders() public view returns (address[] memory) {
    return s_authorizedSendersList;
  }

  /**
   * @notice Use this to check if a node is authorized for fulfilling requests
   * @param sender The address of the Plugin node
   * @return The authorization status of the node
   */
  function isAuthorizedSender(address sender) public view returns (bool) {
    if (!authorizedReceiverActive()) {
      return true;
    }
    return s_authorizedSenders.contains(sender);
  }

  /**
   * @notice customizable guard of who can update the authorized sender list
   * @return bool whether sender can update authorized sender list
   */
  function _canSetAuthorizedSenders() internal virtual returns (bool);

  /**
   * @notice validates the sender is an authorized sender
   */
  function _validateIsAuthorizedSender() internal view {
    if (!isAuthorizedSender(tx.origin)) {
      revert UnauthorizedSender();
    }
  }

  /**
   * @notice prevents non-authorized addresses from calling this method
   */
  modifier validateAuthorizedSender() {
    _validateIsAuthorizedSender();
    _;
  }

  /**
   * @notice prevents non-authorized addresses from calling this method
   */
  modifier validateAuthorizedSenderSetter() {
    if (!_canSetAuthorizedSenders()) {
      revert NotAllowedToSetSenders();
    }
    _;
  }
}
