pragma solidity ^0.8.16;

//  声明存储
contract Store {
    mapping(string => string) public data_contract;
    string public version_;

    constructor(string memory vers) {
        version_ = vers;
    }

    function setItem(string memory key, string memory value) public {
        data_contract[key] = value;
    }

    function getItem(string memory key) public view returns (string memory) {
        return data_contract[key];
    }
    function versionContract() public view returns (string memory) {
        return version_;
    }
}