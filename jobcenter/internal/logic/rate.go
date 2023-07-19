package logic

import (
	"encoding/json"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"log"
	"mscoin-common/tools"
	"sync"
	"time"
)

//汇率

type Rate struct {
	wg    sync.WaitGroup
	c     OkxConfig
	cache cache.Cache
}

type OkxExchangeRateResult struct {
	Code string         `json:"code"`
	Msg  string         `json:"msg"`
	Data []ExchangeRate `json:"data"`
}

type ExchangeRate struct {
	UsdCny string `json:"usdCny"`
}

func (r *Rate) Do() {
	//获取人民币兑换美金汇率
	r.wg.Add(1)
	go r.CynUsdRate()
	r.wg.Wait()
}

func (r *Rate) CynUsdRate() {
	//请求对应接口 获取最新的汇率 存入redis
	//发起http请求 获取数据
	api := r.c.Host + "/api/v5/market/exchange-rate"
	timestamp := tools.ISO(time.Now())
	sign := tools.ComputeHmacSha256(timestamp+"GET"+"/api/v5/market/exchange-rate", r.c.SecretKey)
	header := make(map[string]string)
	header["OK-ACCESS-KEY"] = r.c.ApiKey
	header["OK-ACCESS-SIGN"] = sign
	header["OK-ACCESS-TIMESTAMP"] = timestamp
	header["OK-ACCESS-PASSPHRASE"] = r.c.Pass
	resp, err := tools.GetWithHeader(api, header, r.c.Proxy)
	if err != nil {
		log.Println(err)
		r.wg.Done()
		return
	}
	var result = &OkxExchangeRateResult{}
	err = json.Unmarshal(resp, result)
	if err != nil {
		log.Println(err)
		r.wg.Done()
		return
	}
	cny := result.Data[0].UsdCny
	//存入redis
	r.cache.Set("USDT::CNY::RATE", cny)
	r.wg.Done()
}
func NewRate(c OkxConfig, cache2 cache.Cache) *Rate {
	return &Rate{
		c:     c,
		cache: cache2,
	}
}
