pragma solidity 0.5.0;

contract TxRelay {

    // version number for each sender.
    mapping(address => uint) public nonce;
    // Asking for relaying is allowed only for registered users;
    mapping(address => bool) registeredAddresses;
    // relay destinations should be whitelisted.
    mapping(address => bool) destinationWhitelist;

    function relayMetaTx(
        uint8 sigV,
        bytes32 sigR,
        bytes32 sigS,
        address destination,
        bytes memory data,
        address originalSigner
    ) public {
        // TODO: this should be managed in registry contract.
        require(registeredAddresses[msg.sender], "msg sender should be registered");

        // TODO: this also should be managed in whitelist contract.
        require(destinationWhitelist[destination], "destination should be registered");

        // use EIP 191
        // 0x19 :: version :: relay :: whitelistOwner :: nonce :: destination :: data
        bytes32 h = keccak256(abi.encodePacked(byte(0x19), byte(0x45), address(this), originalSigner, nonce[originalSigner], destination, data));

        address addressFromSig = ecrecover(h, sigV, sigR, sigS);

        require(originalSigner == addressFromSig, "signature should be verified");

        // increment local version number.
        nonce[originalSigner]++;

        // solhint-disable-next-line
        (bool success, ) = destination.call(data);
        if (!success) {
            revert("transaction execution failed");
        }
    }

    function addToWhitelist(address addr) public {
        registeredAddresses[addr] = true;
    }

    function removeFromWhitelist(address addr) public {
        registeredAddresses[addr] = false;
    }

    // handle contract or not.
    function addDestination(address addr) public {
        destinationWhitelist[addr] = true;
    }

    function removeDestination(address addr) public {
        destinationWhitelist[addr] = false;
    }
}