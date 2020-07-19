package response

const (
	Failed = 20000
	Success = 10000
)

type Result struct{
	Code int `json:"code"`
	Msg string `json:"message"`
	Result interface{} `json:"data"`
}