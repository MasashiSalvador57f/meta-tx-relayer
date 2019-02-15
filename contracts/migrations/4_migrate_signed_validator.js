var SignValidator = artifacts.require("./SignValidator.sol");

module.exports = function(deployer) {
    deployer.deploy(SignValidator);
};
