package tool3contract

import (
	"context"
	"encoding/json"
	"math/big"
	"time"
	"tool3/tool3contract/abi/erc721"
	"tool3/tool3storage"

	"github.com/spf13/cast"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
)

type Erc721Contract struct {
	Address  ethcommon.Address
	contract *erc721.Erc721
}

type Erc721Meta struct {
	RawMeta    []byte `json:"-"`
	Image      string `json:"image"`
	Attributes []struct {
		TraitType string `json:"trait_type"`
		Value     string `json:"value"`
	} `json:"attributes"`
}

func NewErc721Contract(rpcUrl string, contractAddress string) (e *Erc721Contract, err error) {
	e = &Erc721Contract{
		Address: ethcommon.HexToAddress(contractAddress),
	}

	ethClient, err := InitEthClient(rpcUrl)
	if err != nil {
		log.Err(err).Msg("InitEthClient")
		return nil, err
	}

	c, err := erc721.NewErc721(e.Address, ethClient)
	if err != nil {
		log.Err(err).Msg("NewErc721")
		return nil, err
	}
	e.contract = c
	return
}

func InitEthClient(dialUrl string) (ethClient *ethclient.Client, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	ethClient, err = ethclient.DialContext(ctx, dialUrl)
	return
}

func (e *Erc721Contract) TotalSupply() (totalSupply int64, err error) {
	res, err := e.contract.TotalSupply(nil)
	if err != nil {
		return 0, err
	}
	totalSupply = cast.ToInt64(res.String())
	return
}

func (e *Erc721Contract) NftMeta(id int64) (meta *Erc721Meta, err error) {
	tokenURI, err := e.contract.TokenURI(nil, big.NewInt(id))
	if err != nil {
		return nil, err
	}

	rawMeta, err := tool3storage.GetIPFSContent(tokenURI)
	if err != nil {
		return nil, err
	}

	meta = new(Erc721Meta)
	err = json.Unmarshal(rawMeta, meta)
	if err != nil {
		return nil, err
	}
	meta.RawMeta = rawMeta
	return
}

func (e *Erc721Contract) NftImage(id int64) (body []byte, meta *Erc721Meta, err error) {
	meta, err = e.NftMeta(id)
	if err != nil {
		return nil, nil, err
	}

	body, err = tool3storage.GetIPFSContent(meta.Image)
	if err != nil {
		return nil, nil, err
	}
	return
}
