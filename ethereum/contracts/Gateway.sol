pragma solidity >=0.4.22 <0.6.3;

contract Gateway {
    enum Status {
        None,
        New,
        Rejected,
        Accepted,
        Success,
        Returned
    }
    enum Type {
        Burn,
        Mint
    }
    struct Request {
        Status status;
        Type rType;
        address owner;
        string target;
        bytes secret;
        bytes32 secretHash;
        uint256 tokenAmount;
        uint expireHeight;
    }

    address[] admins;
    address tokenAddress;
    constructor(address[5] memory newAdmins, address noewTokenAddress) public {
        admins = newAdmins;
        tokenAddress = noewTokenAddress;
    }

    event NewRequest(bytes32 requestHash, address owner, string target, bytes32 secretHash, uint expireHeight);
    event StatusChanged(bytes32 requestHash, Status status);

    mapping(bytes32 => Request) public requests;
    function returnToSender(bytes32 requestHash) public {
        require(block.number > requests[requestHash].expireHeight, "timeout is not over yet");
        require(requests[requestHash].rType != Type.Burn, "type is not burn");
        require(requests[requestHash].status != Status.Success, "status is success");
        require(requests[requestHash].status != Status.Returned, "status is returned");

        msg.sender.transfer(requests[requestHash].tokenAmount);
        requests[requestHash].status = Status.Returned;
    }

    function confirmMint(bytes32 requestHash, bytes memory secret) public {
        require(block.number <= requests[requestHash].expireHeight, "timeout is over yet");
        require(requests[requestHash].rType == Type.Mint, "type is not mint");
        require(requests[requestHash].status == Status.Accepted, "status is not accepted");
        requests[requestHash].secret = secret;
        requests[requestHash].status = Status.Success;
        emit StatusChanged(requestHash, Status.Success);
    }

    function createBurnRequest(string memory recipient, bytes32 secretHash, uint expireHeight) public payable {
        bytes32 requestHash = keccak256(abi.encodePacked(this, msg.sender, recipient, secretHash));
        require(msg.value > 0, "value less or equals 0");
        require(expireHeight > 0, "expire height less or equals 0");
        require(requests[requestHash].status == Status.None, "swap exist");

        bytes memory empty;
        requests[requestHash] = Request(
            {status: Status.New, rType: Type.Burn, owner: msg.sender, target: recipient, secret: empty, secretHash: secretHash, tokenAmount: msg.value, expireHeight: expireHeight });
        emit NewRequest(requestHash, msg.sender, recipient, secretHash, expireHeight);
    }

    function createMintRequest(string memory sender, address recipient, uint256 value, bytes32 secretHash, uint expireHeight,
            uint8[5] memory v, bytes32[5] memory r, bytes32[5] memory s) public payable {
        bytes32 requestHash = keccak256(abi.encodePacked(this, msg.sender, recipient, secretHash));
        require(msg.value > 0, "value less or equals 0");
        require(expireHeight > 0, "expire height less or equals 0");
        require(requests[requestHash].status == Status.None, "swap exist");

        int count = 0;
        for(uint i = 0; i < 5; i++) {
            count += ecrecover(keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", requestHash)),
                v[i], r[i], s[i]) == admins[0] ? 1 : 0;
        }

        require(count >= 3, "admins vote count is less 3");
        bytes memory empty;
        requests[requestHash] = Request(
            {status: Status.Accepted, rType: Type.Burn, owner: recipient, target: sender, secret: empty, secretHash: secretHash, tokenAmount: value, expireHeight: expireHeight });
        emit NewRequest(requestHash, recipient, sender, secretHash, expireHeight);
    }

    function acceptedOrReject(bytes32 requestHash, uint8[5] memory v, bytes32[5] memory r, bytes32[5] memory s, bool isAccepted) public payable {
        require(block.number <= requests[requestHash].expireHeight, "timeout is not over yet");
        require(requests[requestHash].status == Status.New, "status is now new");

        int count = 0;
        for(uint i = 0; i < 5; i++) {
            count += ecrecover(keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", requestHash, (isAccepted ? "true" : "false"))),
                v[i], r[i], s[i]) == admins[0] ? 1 : 0;
        }

        Status status = isAccepted ? Status.Accepted : Status.Rejected;
        require(count >= 3, "admins vote count is less 3");
        requests[requestHash].status = status;
        //mint
        emit StatusChanged(requestHash, status);
    }
}