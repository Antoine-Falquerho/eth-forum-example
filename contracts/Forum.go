// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ForumUser is an auto generated low-level Go binding around an user-defined struct.
type ForumUser struct {
	Name      string
	Karma     *big.Int
	PostCount *big.Int
}

// ForumABI is the input ABI used to generate the binding from.
const ForumABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"}],\"name\":\"addPost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_post_id\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"_vote\",\"type\":\"int256\"}],\"name\":\"addVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getUser\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"karma\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"post_count\",\"type\":\"uint256\"}],\"internalType\":\"structForum.User\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"last_post_id\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"postVotes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"vote\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"posts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"karma\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"postedAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userPosts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"karma\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"postedAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ForumBin is the compiled bytecode used for deploying new contracts.
var ForumBin = "0x608060405234801561001057600080fd5b50610c2f806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80638064d1491161005b5780638064d1491461014c578063a7943a1d1461015f578063c264ea4314610172578063f78f838f1461018557600080fd5b80630b1e7f831461008d5780633121db1c146100ba57806343699c5c146100cf5780636f77926b1461012c575b600080fd5b6100a061009b366004610a20565b61019c565b6040516100b1959493929190610ab2565b60405180910390f35b6100cd6100c8366004610901565b6102e5565b005b61010d6100dd366004610a39565b6004602090815260009283526040808420909152908252902080546001909101546001600160a01b039091169082565b604080516001600160a01b0390931683526020830191909152016100b1565b61013f61013a3660046108df565b610313565b6040516100b19190610afd565b6100a061015a3660046109c3565b610406565b6100cd61016d3660046109ed565b61043d565b6100cd61018036600461094f565b610657565b61018e60005481565b6040519081526020016100b1565b600260205260009081526040902080546001820180546001600160a01b0390921692916101c890610b92565b80601f01602080910402602001604051908101604052809291908181526020018280546101f490610b92565b80156102415780601f1061021657610100808354040283529160200191610241565b820191906000526020600020905b81548152906001019060200180831161022457829003601f168201915b50505050509080600201805461025690610b92565b80601f016020809104026020016040519081016040528092919081815260200182805461028290610b92565b80156102cf5780601f106102a4576101008083540402835291602001916102cf565b820191906000526020600020905b8154815290600101906020018083116102b257829003601f168201915b5050505050908060030154908060040154905085565b6001600160a01b0382166000908152600160209081526040909120825161030e9284019061079d565b505050565b61033760405180606001604052806060815260200160008152602001600081525090565b6001600160a01b0382166000908152600160205260409081902081516060810190925280548290829061036990610b92565b80601f016020809104026020016040519081016040528092919081815260200182805461039590610b92565b80156103e25780601f106103b7576101008083540402835291602001916103e2565b820191906000526020600020905b8154815290600101906020018083116103c557829003601f168201915b50505050508152602001600182015481526020016002820154815250509050919050565b6003602090815260009283526040808420909152908252902080546001820180546001600160a01b0390921692916101c890610b92565b6040805180820182523381526020808201849052600085815260028252838120845160a0810190955280546001600160a01b0316855260018101805494959294929391929184019161048e90610b92565b80601f01602080910402602001604051908101604052809291908181526020018280546104ba90610b92565b80156105075780601f106104dc57610100808354040283529160200191610507565b820191906000526020600020905b8154815290600101906020018083116104ea57829003601f168201915b5050505050815260200160028201805461052090610b92565b80601f016020809104026020016040519081016040528092919081815260200182805461054c90610b92565b80156105995780601f1061056e57610100808354040283529160200191610599565b820191906000526020600020905b81548152906001019060200180831161057c57829003601f168201915b505050918352505060038281015460208084019190915260049384015460409384015260008981529381528284206001600160a01b038b81168652908252838520885181546001600160a01b03191692169190911781558782015160019091015588845260029052908220018054929350859290919061061a908490610b39565b909155505080516001600160a01b03166000908152600160208190526040822001805485929061064b908490610b39565b90915550505050505050565b6040805160a0810182526001600160a01b0385811682526020808301868152838501869052600060608501819052426080860152805481526002835294909420835181546001600160a01b031916931692909217825592518051929384936106c5926001850192019061079d565b50604082015180516106e191600284019160209091019061079d565b5060608201516003828101919091556080909201516004909101556001600160a01b038581166000908152600160208181526040808420600201549582528084208685528252909220855181546001600160a01b031916941693909317835584820151805186949361075793850192019061079d565b506040820151805161077391600284019160209091019061079d565b506060820151816003015560808201518160040155905050600160008082825461064b9190610b7a565b8280546107a990610b92565b90600052602060002090601f0160209004810192826107cb5760008555610811565b82601f106107e457805160ff1916838001178555610811565b82800160010185558215610811579182015b828111156108115782518255916020019190600101906107f6565b5061081d929150610821565b5090565b5b8082111561081d5760008155600101610822565b80356001600160a01b038116811461084d57600080fd5b919050565b600082601f83011261086357600080fd5b813567ffffffffffffffff8082111561087e5761087e610be3565b604051601f8301601f19908116603f011681019082821181831017156108a6576108a6610be3565b816040528381528660208588010111156108bf57600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000602082840312156108f157600080fd5b6108fa82610836565b9392505050565b6000806040838503121561091457600080fd5b61091d83610836565b9150602083013567ffffffffffffffff81111561093957600080fd5b61094585828601610852565b9150509250929050565b60008060006060848603121561096457600080fd5b61096d84610836565b9250602084013567ffffffffffffffff8082111561098a57600080fd5b61099687838801610852565b935060408601359150808211156109ac57600080fd5b506109b986828701610852565b9150509250925092565b600080604083850312156109d657600080fd5b6109df83610836565b946020939093013593505050565b600080600060608486031215610a0257600080fd5b610a0b84610836565b95602085013595506040909401359392505050565b600060208284031215610a3257600080fd5b5035919050565b60008060408385031215610a4c57600080fd5b82359150610a5c60208401610836565b90509250929050565b6000815180845260005b81811015610a8b57602081850181015186830182015201610a6f565b81811115610a9d576000602083870101525b50601f01601f19169290920160200192915050565b6001600160a01b038616815260a060208201819052600090610ad690830187610a65565b8281036040840152610ae88187610a65565b60608401959095525050608001529392505050565b602081526000825160606020840152610b196080840182610a65565b905060208401516040840152604084015160608401528091505092915050565b600080821280156001600160ff1b0384900385131615610b5b57610b5b610bcd565b600160ff1b8390038412811615610b7457610b74610bcd565b50500190565b60008219821115610b8d57610b8d610bcd565b500190565b600181811c90821680610ba657607f821691505b60208210811415610bc757634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fdfea26469706673582212204351626652ab72fe8a3f2497dde753e2cc5fc90fba3e7744a18fe65d214be97064736f6c63430008060033"

// DeployForum deploys a new Ethereum contract, binding an instance of Forum to it.
func DeployForum(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Forum, error) {
	parsed, err := abi.JSON(strings.NewReader(ForumABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ForumBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Forum{ForumCaller: ForumCaller{contract: contract}, ForumTransactor: ForumTransactor{contract: contract}, ForumFilterer: ForumFilterer{contract: contract}}, nil
}

// Forum is an auto generated Go binding around an Ethereum contract.
type Forum struct {
	ForumCaller     // Read-only binding to the contract
	ForumTransactor // Write-only binding to the contract
	ForumFilterer   // Log filterer for contract events
}

// ForumCaller is an auto generated read-only Go binding around an Ethereum contract.
type ForumCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForumTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ForumTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForumFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ForumFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForumSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ForumSession struct {
	Contract     *Forum            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ForumCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ForumCallerSession struct {
	Contract *ForumCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ForumTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ForumTransactorSession struct {
	Contract     *ForumTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ForumRaw is an auto generated low-level Go binding around an Ethereum contract.
type ForumRaw struct {
	Contract *Forum // Generic contract binding to access the raw methods on
}

// ForumCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ForumCallerRaw struct {
	Contract *ForumCaller // Generic read-only contract binding to access the raw methods on
}

// ForumTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ForumTransactorRaw struct {
	Contract *ForumTransactor // Generic write-only contract binding to access the raw methods on
}

// NewForum creates a new instance of Forum, bound to a specific deployed contract.
func NewForum(address common.Address, backend bind.ContractBackend) (*Forum, error) {
	contract, err := bindForum(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Forum{ForumCaller: ForumCaller{contract: contract}, ForumTransactor: ForumTransactor{contract: contract}, ForumFilterer: ForumFilterer{contract: contract}}, nil
}

// NewForumCaller creates a new read-only instance of Forum, bound to a specific deployed contract.
func NewForumCaller(address common.Address, caller bind.ContractCaller) (*ForumCaller, error) {
	contract, err := bindForum(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ForumCaller{contract: contract}, nil
}

// NewForumTransactor creates a new write-only instance of Forum, bound to a specific deployed contract.
func NewForumTransactor(address common.Address, transactor bind.ContractTransactor) (*ForumTransactor, error) {
	contract, err := bindForum(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ForumTransactor{contract: contract}, nil
}

// NewForumFilterer creates a new log filterer instance of Forum, bound to a specific deployed contract.
func NewForumFilterer(address common.Address, filterer bind.ContractFilterer) (*ForumFilterer, error) {
	contract, err := bindForum(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ForumFilterer{contract: contract}, nil
}

// bindForum binds a generic wrapper to an already deployed contract.
func bindForum(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ForumABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Forum *ForumRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Forum.Contract.ForumCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Forum *ForumRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Forum.Contract.ForumTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Forum *ForumRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Forum.Contract.ForumTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Forum *ForumCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Forum.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Forum *ForumTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Forum.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Forum *ForumTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Forum.Contract.contract.Transact(opts, method, params...)
}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address _address) view returns((string,int256,uint256))
func (_Forum *ForumCaller) GetUser(opts *bind.CallOpts, _address common.Address) (ForumUser, error) {
	var out []interface{}
	err := _Forum.contract.Call(opts, &out, "getUser", _address)

	if err != nil {
		return *new(ForumUser), err
	}

	out0 := *abi.ConvertType(out[0], new(ForumUser)).(*ForumUser)

	return out0, err

}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address _address) view returns((string,int256,uint256))
func (_Forum *ForumSession) GetUser(_address common.Address) (ForumUser, error) {
	return _Forum.Contract.GetUser(&_Forum.CallOpts, _address)
}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address _address) view returns((string,int256,uint256))
func (_Forum *ForumCallerSession) GetUser(_address common.Address) (ForumUser, error) {
	return _Forum.Contract.GetUser(&_Forum.CallOpts, _address)
}

// LastPostId is a free data retrieval call binding the contract method 0xf78f838f.
//
// Solidity: function last_post_id() view returns(uint256)
func (_Forum *ForumCaller) LastPostId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Forum.contract.Call(opts, &out, "last_post_id")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastPostId is a free data retrieval call binding the contract method 0xf78f838f.
//
// Solidity: function last_post_id() view returns(uint256)
func (_Forum *ForumSession) LastPostId() (*big.Int, error) {
	return _Forum.Contract.LastPostId(&_Forum.CallOpts)
}

// LastPostId is a free data retrieval call binding the contract method 0xf78f838f.
//
// Solidity: function last_post_id() view returns(uint256)
func (_Forum *ForumCallerSession) LastPostId() (*big.Int, error) {
	return _Forum.Contract.LastPostId(&_Forum.CallOpts)
}

// PostVotes is a free data retrieval call binding the contract method 0x43699c5c.
//
// Solidity: function postVotes(uint256 , address ) view returns(address owner, int256 vote)
func (_Forum *ForumCaller) PostVotes(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Owner common.Address
	Vote  *big.Int
}, error) {
	var out []interface{}
	err := _Forum.contract.Call(opts, &out, "postVotes", arg0, arg1)

	outstruct := new(struct {
		Owner common.Address
		Vote  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Vote = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PostVotes is a free data retrieval call binding the contract method 0x43699c5c.
//
// Solidity: function postVotes(uint256 , address ) view returns(address owner, int256 vote)
func (_Forum *ForumSession) PostVotes(arg0 *big.Int, arg1 common.Address) (struct {
	Owner common.Address
	Vote  *big.Int
}, error) {
	return _Forum.Contract.PostVotes(&_Forum.CallOpts, arg0, arg1)
}

// PostVotes is a free data retrieval call binding the contract method 0x43699c5c.
//
// Solidity: function postVotes(uint256 , address ) view returns(address owner, int256 vote)
func (_Forum *ForumCallerSession) PostVotes(arg0 *big.Int, arg1 common.Address) (struct {
	Owner common.Address
	Vote  *big.Int
}, error) {
	return _Forum.Contract.PostVotes(&_Forum.CallOpts, arg0, arg1)
}

// Posts is a free data retrieval call binding the contract method 0x0b1e7f83.
//
// Solidity: function posts(uint256 ) view returns(address owner, string title, string content, int256 karma, uint256 postedAt)
func (_Forum *ForumCaller) Posts(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Owner    common.Address
	Title    string
	Content  string
	Karma    *big.Int
	PostedAt *big.Int
}, error) {
	var out []interface{}
	err := _Forum.contract.Call(opts, &out, "posts", arg0)

	outstruct := new(struct {
		Owner    common.Address
		Title    string
		Content  string
		Karma    *big.Int
		PostedAt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Title = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Content = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Karma = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.PostedAt = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Posts is a free data retrieval call binding the contract method 0x0b1e7f83.
//
// Solidity: function posts(uint256 ) view returns(address owner, string title, string content, int256 karma, uint256 postedAt)
func (_Forum *ForumSession) Posts(arg0 *big.Int) (struct {
	Owner    common.Address
	Title    string
	Content  string
	Karma    *big.Int
	PostedAt *big.Int
}, error) {
	return _Forum.Contract.Posts(&_Forum.CallOpts, arg0)
}

// Posts is a free data retrieval call binding the contract method 0x0b1e7f83.
//
// Solidity: function posts(uint256 ) view returns(address owner, string title, string content, int256 karma, uint256 postedAt)
func (_Forum *ForumCallerSession) Posts(arg0 *big.Int) (struct {
	Owner    common.Address
	Title    string
	Content  string
	Karma    *big.Int
	PostedAt *big.Int
}, error) {
	return _Forum.Contract.Posts(&_Forum.CallOpts, arg0)
}

// UserPosts is a free data retrieval call binding the contract method 0x8064d149.
//
// Solidity: function userPosts(address , uint256 ) view returns(address owner, string title, string content, int256 karma, uint256 postedAt)
func (_Forum *ForumCaller) UserPosts(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Owner    common.Address
	Title    string
	Content  string
	Karma    *big.Int
	PostedAt *big.Int
}, error) {
	var out []interface{}
	err := _Forum.contract.Call(opts, &out, "userPosts", arg0, arg1)

	outstruct := new(struct {
		Owner    common.Address
		Title    string
		Content  string
		Karma    *big.Int
		PostedAt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Title = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Content = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Karma = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.PostedAt = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserPosts is a free data retrieval call binding the contract method 0x8064d149.
//
// Solidity: function userPosts(address , uint256 ) view returns(address owner, string title, string content, int256 karma, uint256 postedAt)
func (_Forum *ForumSession) UserPosts(arg0 common.Address, arg1 *big.Int) (struct {
	Owner    common.Address
	Title    string
	Content  string
	Karma    *big.Int
	PostedAt *big.Int
}, error) {
	return _Forum.Contract.UserPosts(&_Forum.CallOpts, arg0, arg1)
}

// UserPosts is a free data retrieval call binding the contract method 0x8064d149.
//
// Solidity: function userPosts(address , uint256 ) view returns(address owner, string title, string content, int256 karma, uint256 postedAt)
func (_Forum *ForumCallerSession) UserPosts(arg0 common.Address, arg1 *big.Int) (struct {
	Owner    common.Address
	Title    string
	Content  string
	Karma    *big.Int
	PostedAt *big.Int
}, error) {
	return _Forum.Contract.UserPosts(&_Forum.CallOpts, arg0, arg1)
}

// AddPost is a paid mutator transaction binding the contract method 0xc264ea43.
//
// Solidity: function addPost(address _owner, string title, string content) returns()
func (_Forum *ForumTransactor) AddPost(opts *bind.TransactOpts, _owner common.Address, title string, content string) (*types.Transaction, error) {
	return _Forum.contract.Transact(opts, "addPost", _owner, title, content)
}

// AddPost is a paid mutator transaction binding the contract method 0xc264ea43.
//
// Solidity: function addPost(address _owner, string title, string content) returns()
func (_Forum *ForumSession) AddPost(_owner common.Address, title string, content string) (*types.Transaction, error) {
	return _Forum.Contract.AddPost(&_Forum.TransactOpts, _owner, title, content)
}

// AddPost is a paid mutator transaction binding the contract method 0xc264ea43.
//
// Solidity: function addPost(address _owner, string title, string content) returns()
func (_Forum *ForumTransactorSession) AddPost(_owner common.Address, title string, content string) (*types.Transaction, error) {
	return _Forum.Contract.AddPost(&_Forum.TransactOpts, _owner, title, content)
}

// AddVote is a paid mutator transaction binding the contract method 0xa7943a1d.
//
// Solidity: function addVote(address _owner, uint256 _post_id, int256 _vote) returns()
func (_Forum *ForumTransactor) AddVote(opts *bind.TransactOpts, _owner common.Address, _post_id *big.Int, _vote *big.Int) (*types.Transaction, error) {
	return _Forum.contract.Transact(opts, "addVote", _owner, _post_id, _vote)
}

// AddVote is a paid mutator transaction binding the contract method 0xa7943a1d.
//
// Solidity: function addVote(address _owner, uint256 _post_id, int256 _vote) returns()
func (_Forum *ForumSession) AddVote(_owner common.Address, _post_id *big.Int, _vote *big.Int) (*types.Transaction, error) {
	return _Forum.Contract.AddVote(&_Forum.TransactOpts, _owner, _post_id, _vote)
}

// AddVote is a paid mutator transaction binding the contract method 0xa7943a1d.
//
// Solidity: function addVote(address _owner, uint256 _post_id, int256 _vote) returns()
func (_Forum *ForumTransactorSession) AddVote(_owner common.Address, _post_id *big.Int, _vote *big.Int) (*types.Transaction, error) {
	return _Forum.Contract.AddVote(&_Forum.TransactOpts, _owner, _post_id, _vote)
}

// SetName is a paid mutator transaction binding the contract method 0x3121db1c.
//
// Solidity: function setName(address _owner, string _name) returns()
func (_Forum *ForumTransactor) SetName(opts *bind.TransactOpts, _owner common.Address, _name string) (*types.Transaction, error) {
	return _Forum.contract.Transact(opts, "setName", _owner, _name)
}

// SetName is a paid mutator transaction binding the contract method 0x3121db1c.
//
// Solidity: function setName(address _owner, string _name) returns()
func (_Forum *ForumSession) SetName(_owner common.Address, _name string) (*types.Transaction, error) {
	return _Forum.Contract.SetName(&_Forum.TransactOpts, _owner, _name)
}

// SetName is a paid mutator transaction binding the contract method 0x3121db1c.
//
// Solidity: function setName(address _owner, string _name) returns()
func (_Forum *ForumTransactorSession) SetName(_owner common.Address, _name string) (*types.Transaction, error) {
	return _Forum.Contract.SetName(&_Forum.TransactOpts, _owner, _name)
}
