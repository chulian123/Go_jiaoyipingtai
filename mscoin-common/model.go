package common

type BizCode int

const SuccessCode BizCode = 0

type Result struct {
	Code    BizCode `json:"code"`
	Message string  `json:"message"`
	Data    any     `json:"data"`
}

func NewResult() *Result {
	return &Result{}
}

func (r *Result) Fail(code BizCode, msg string) {
	r.Code = code
	r.Message = msg
}

func (r *Result) Success(data any) {
	r.Code = SuccessCode
	r.Message = "success"
	r.Data = data
}

// Deal err 进行自定义 code 和 msg 有两个属性的err
func (r *Result) Deal(data any, err error) *Result {
	if err != nil {
		r.Fail(500, err.Error())
		return r
	}
	r.Success(data)
	return r
}
