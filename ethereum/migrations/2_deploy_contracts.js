const Gateway = artifacts.require("./Gateway.sol");
module.exports = function(deployer, network, accounts) {
    deployer.deploy(Gateway);
};