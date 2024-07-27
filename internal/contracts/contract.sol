pragma solidity ^0.8.2;

contract Wallet {
    uint256 balance;
    address public admin;

    constructor() {
        admin = msg.sender;
        balance = 0;
        updateBalance();
    }

    function updateBalance() internal {
        balance += msg.value;
    }

    function Withdraw(uint256 _amt) public {
        require(msg.sender == admin);
        balance = balance - _amt;
    }

    function Deposit(uint256 amt) public returns (uint256) {
        return balance = balance + amt;
    }

    function Balance() public view returns(uint256) {
        return balance;
    }
}

// solcjs --optimize --abi ./internal/contracts/contract.sol -o internal/contracts/build
// solcjs --optimize --bin ./internal/contracts/contract.sol -o internal/contracts/build
// abigen --abi=./internal/contracts/build/internal_contracts_contract_sol_Wallet.abi --bin=./internal/contracts/build/internal_contracts_contract_sol_Wallet.bin --pkg=api --out=./internal/contracts/contract.go