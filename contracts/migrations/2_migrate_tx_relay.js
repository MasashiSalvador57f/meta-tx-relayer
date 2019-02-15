var TxRelay = artifacts.require("./TxRelay.sol");

module.exports = function(deployer) {
    deployer.deploy(TxRelay);
};
