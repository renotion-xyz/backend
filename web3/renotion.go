package web3

import (
	"os"
	"reflect"

	"github.com/chenzhijie/go-web3"
	"github.com/chenzhijie/go-web3/eth"
)

type Renotion struct {
	web3     *web3.Web3
	contract *eth.Contract
}

func (rnt *Renotion) GetDomainMetadata(tokenID string) (*DomainMetadata, error) {
	result, err := rnt.contract.Call("metadataFor", uint256ToBytes32(tokenID))
	if err != nil {
		return nil, err
	}
	value := reflect.ValueOf(result)
	metadata := DomainMetadata{
		Hostname: value.FieldByName("Hostname").String(),
		Page:     value.FieldByName("Page").String(),
	}
	return &metadata, nil
}

func NewRenotion(rpcUrl string, contractAddress string) (*Renotion, error) {
	web3, err := web3.NewWeb3(rpcUrl)
	if err != nil {
		return nil, err
	}
	contents, err := os.ReadFile("./artifacts/Renotion.abi.json")
	if err != nil {
		return nil, err
	}
	abiString := string(contents)
	contract, err := web3.Eth.NewContract(abiString, contractAddress)
	if err != nil {
		return nil, err
	}
	return &Renotion{web3, contract}, nil
}
