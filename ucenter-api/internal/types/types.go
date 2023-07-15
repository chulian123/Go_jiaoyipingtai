// Code generated by goctl. DO NOT EDIT.
package types

type Request struct {
	Username     string      `json:"username,optional"`
	Password     string      `json:"password,optional"`
	Captcha      *CaptchaReq `json:"captcha,optional"`
	Phone        string      `json:"phone,optional"`
	Promotion    string      `json:"promotion,optional"`
	Code         string      `json:"code,optional"`
	Country      string      `json:"country,optional"`
	SuperPartner string      `json:"superPartner,optional"`
	Ip           string      `json:"ip,optional"`
}

type CaptchaReq struct {
	Server string `json:"server"`
	Token  string `json:"token"`
}

type Response struct {
	Message string `json:"message"`
}

type CodeRequest struct {
	Phone   string `json:"phone,optional"`
	Country string `json:"country,optional"`
}
type CodeResponse struct {
}

type LoginReq struct {
	Username string      `json:"username"`
	Password string      `json:"password"`
	Captcha  *CaptchaReq `json:"captcha,optional"`
	Ip       string      `json:"ip,optional"`
}

type LoginRes struct {
	Username      string `json:"username"`
	Token         string `json:"token"`
	MemberLevel   string `json:"memberLevel"`
	RealName      string `json:"realName"`
	Country       string `json:"country"`
	Avatar        string `json:"avatar"`
	PromotionCode string `json:"promotionCode"`
	Id            int64  `json:"id"`
	LoginCount    int    `json:"loginCount"`
	SuperPartner  string `json:"superPartner"`
	MemberRate    int    `json:"memberRate"`
}

type AssetReq struct {
	CoinName string `json:"coinName,optional" path:"coinName,optional"`
	Ip       string `json:"ip,optional"`
}

type Coin struct {
	Id                int     `json:"id" from:"id"`
	Name              string  `json:"name" from:"name"`
	CanAutoWithdraw   int     `json:"canAutoWithdraw" from:"canAutoWithdraw"`
	CanRecharge       int     `json:"canRecharge" from:"canRecharge"`
	CanTransfer       int     `json:"canTransfer" from:"canTransfer"`
	CanWithdraw       int     `json:"canWithdraw" from:"canWithdraw"`
	CnyRate           float64 `json:"cnyRate" from:"cnyRate"`
	EnableRpc         int     `json:"enableRpc" from:"enableRpc"`
	IsPlatformCoin    int     `json:"isPlatformCoin" from:"isPlatformCoin"`
	MaxTxFee          float64 `json:"maxTxFee" from:"maxTxFee"`
	MaxWithdrawAmount float64 `json:"maxWithdrawAmount" from:"maxWithdrawAmount"`
	MinTxFee          float64 `json:"minTxFee" from:"minTxFee"`
	MinWithdrawAmount float64 `json:"minWithdrawAmount" from:"minWithdrawAmount"`
	NameCn            string  `json:"nameCn" from:"nameCn"`
	Sort              int     `json:"sort" from:"sort"`
	Status            int     `json:"status" from:"status"`
	Unit              string  `json:"unit" from:"unit"`
	UsdRate           float64 `json:"usdRate" from:"usdRate"`
	WithdrawThreshold float64 `json:"withdrawThreshold" from:"withdrawThreshold"`
	HasLegal          int     `json:"hasLegal" from:"hasLegal"`
	ColdWalletAddress string  `json:"coldWalletAddress" from:"coldWalletAddress"`
	MinerFee          float64 `json:"minerFee" from:"minerFee"`
	WithdrawScale     int     `json:"withdrawScale" from:"withdrawScale"`
	AccountType       int     `json:"accountType" from:"accountType"`
	DepositAddress    string  `json:"depositAddress" from:"depositAddress"`
	Infolink          string  `json:"infolink" from:"infolink"`
	Information       string  `json:"information" from:"information"`
	MinRechargeAmount float64 `json:"minRechargeAmount" from:"minRechargeAmount"`
}

type MemberWallet struct {
	Id             int64   `json:"id" from:"id"`
	Address        string  `json:"address" from:"address"`
	Balance        float64 `json:"balance" from:"balance"`
	FrozenBalance  float64 `json:"frozenBalance" from:"frozenBalance"`
	ReleaseBalance float64 `json:"releaseBalance" from:"releaseBalance"`
	IsLock         int     `json:"isLock" from:"isLock"`
	MemberId       int64   `json:"memberId" from:"memberId"`
	Version        int     `json:"version" from:"version"`
	Coin           Coin    `json:"coin" from:"coinId"`
	ToReleased     float64 `json:"toReleased" from:"toReleased"`
}
