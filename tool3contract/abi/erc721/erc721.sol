pragma solidity ^0.8.16;

interface ERC721 {
    function balanceOf(address owner)  external  view returns (uint256);
    function ownerOf(uint256 tokenId)  external view  returns (address);
    function name()  external  view returns (string memory);
    function symbol()   external  view returns (string memory);
    function tokenURI(uint256 tokenId) external view  returns (string memory);
    function totalSupply() external view returns (uint256);
}