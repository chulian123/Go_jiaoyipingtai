package common

type BizCode int

const SuccessCode BizCode = 0

type Result struct {
	Code BizCode `json:"code"`
	Msg  string  `json:"msg"`
	Data any     `json:"data"`
}

// NewResult 如果成功就返回的内容
func NewResult() *Result {
	return &Result{}
}

// Fail  如果失败就返回自定义的内容和code
func (r *Result) Fail(code BizCode, msg string) {
	r.Code = code
	r.Msg = msg
}

// Success  如果失败就返回自定义的内容和code
func (r *Result) Success(data any) {
	r.Code = SuccessCode
	r.Msg = "success"
	r.Data = data
}

// Deal  通过自定义err 定义code和msg 有两个属性err
func (r *Result) Deal(data any, err error) *Result {
	if err != nil {
		r.Fail(-999, err.Error())
		return r
	}
	r.Success(data)
	return r
}
