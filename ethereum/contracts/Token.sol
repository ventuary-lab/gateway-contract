pragma solidity >=0.4.22 <0.6.3;

import "../../node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "../../node_modules/@openzeppelin/contracts/access/roles/MinterRole.sol";

contract Token is ERC20, MinterRole {
   
  
    function mint(address account, uint256 amount) public onlyMinter returns (bool) {
        _mint(account, amount);
        return true;
    }
}
