package domain

import (
	"time"

	"github.com/tidwall/gjson"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
)

var (
	domainClient = resty.New().SetTimeout(5 * time.Second)
)

func CheckBnbDomain(proxyList []string, domain string) (noRegistered bool, succeed bool, err error) {
	u := "https://backend.prd.space.id/nameof"

	payload := map[string]interface{}{
		"ChainID": 56,
		"name":    domain,
	}

	if len(proxyList) > 0 {
		i := time.Now().UnixNano() / 1e4 % int64(len(proxyList))
		proxy := proxyList[i]
		domainClient = domainClient.SetProxy(proxy)
	}

	resp, err := domainClient.R().SetBody(payload).Post(u)
	if err != nil {
		log.Err(err).Msg("CheckBnbDomain request")
		return
	}

	if resp.StatusCode() != 200 {
		return
	}

	succeed = true
	owner := gjson.Get(resp.String(), "Owner").String()
	if owner == "" {
		noRegistered = true

	}
	return
}
