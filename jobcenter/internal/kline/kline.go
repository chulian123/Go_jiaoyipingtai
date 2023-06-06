package kline

import (
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
	"mscoin-common/tools"
	"sync"
	"time"
)

type OkxResult struct {
	Code string     `json:"code"`
	Msg  string     `json:"msg"`
	Data [][]string `json:"data"`
}

type OkxConfig struct {
	ApiKey    string
	SecretKey string
	Pass      string
	Host      string
	Proxy     string
}

type Kline struct {
	wg sync.WaitGroup
	c  OkxConfig
}

func (k *Kline) Do(period string) {
	k.wg.Add(2)
	//获取某个数据 币 BTC-USDT  ETH-USDT
	go k.getKlineData("BTC-USDT", period)
	go k.getKlineData("ETH-USDT", period)
	k.wg.Wait()
}

func (k *Kline) getKlineData(instId string, period string) {
	//发起http请求
	api := k.c.Host + "/api/v5/market/candles?instId=" + instId + "&bar=" + period
	header := make(map[string]string)
	timestamp := tools.ISO(time.Now())
	sign := tools.ComputeHmacSha256(timestamp+"GET"+"/api/v5/market/candles?instId="+instId+"&bar="+period, k.c.SecretKey)
	//	sign := base64.StdEncoding.EncodeToString([]byte(sha256))
	header["OK-ACCESS-KEY"] = k.c.ApiKey
	header["OK-ACCESS-SIGN"] = sign
	header["OK-ACCESS-TIMESTAMP"] = timestamp
	header["OK-ACCESS-PASSPHRASE"] = k.c.Pass
	resp, err := tools.GetWithHeader(api, header, k.c.Proxy) //这里地址是指你的代理软件地址 vpn clash软件的http端口
	if err != nil {
		logx.Error(err)
		k.wg.Done()
		return
	}
	var result = &OkxResult{}
	err = json.Unmarshal(resp, result)
	if err != nil {
		logx.Error(err)
		k.wg.Done()
		return
	}
	log.Println("====开始获取数据======")
	log.Println("instId:", instId, "period", period)
	log.Println("result  kline data: ", string(resp))
	k.wg.Done()
	log.Println("==================End====================")

}

func NewKline(c OkxConfig) *Kline {
	return &Kline{
		c: c,
	}
}
