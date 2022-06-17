// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

// re:register; ac: account; t: token; p: provider; k: keeper; u: user; pp: pledgepool
interface IRoleSetter {
    event ReAcc(address addr, uint64 index); // to get all registered account by filter logs
    event ReRole(uint8 indexed rType, uint64 index); // to get all users/keepers/providers by filter logs
    event SetG(uint64 indexed gi, uint64 index); // to get all accounts in one group by filter logs
    event AlterPayee(address indexed payee, address addr);
    event ConfirmPayee(address indexed payee, address addr);

    // called by admin
    function ban(uint64 _i, bool _ban) external;

    function reAcc(address _a) external;

    function reRole(
        uint64 _i,
        uint8 _rtype,
        bytes memory extra
    ) external;

    function alterPayee(uint64 _index, address _p) external;

    function confirmPayee(uint64 _index, address _p) external;

    function setG(uint64 _i, uint64 _gi) external;

    function activate(uint64 _i) external returns (uint64);
}
