package tool3storage

import (
	"errors"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
)

var (
	ipfsClient  = resty.New().SetTimeout(10 * time.Second)
	ipfsGateway string
)

func CommonHeader() map[string]string {
	header := map[string]string{
		`Cache-Control`: `no-cache`,
		`Connection`:    `keep-alive`,
		`User-Agent`:    `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.61 Safari/537.36`,
	}
	return header
}

func SetIpfsGateway(gateway string) {
	ipfsGateway = gateway
}

func checkGateway() {
	if ipfsGateway == "" {
		panic("please SetIpfsGateway first")
	}
}

func GetIPFSContent(hash string) (body []byte, err error) {
	checkGateway()

	hash = strings.ReplaceAll(hash, "ipfs://", "")
	resp, err := ipfsClient.R().SetHeaders(CommonHeader()).Get(ipfsGateway + hash)
	if err != nil {
		return nil, err
	}

	if strings.Contains(resp.String(), "</html>") {
		return nil, errors.New("ipfs too many request")
	}

	return resp.Body(), nil
}
