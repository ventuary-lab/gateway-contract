const Supersymmetry = artifacts.require("./Supersymmetry.sol");

module.exports = async function(deployer, network, accounts) {
    await deployer.deploy(Supersymmetry, [accounts[0], accounts[1],accounts[2],accounts[3],accounts[4]]);
};