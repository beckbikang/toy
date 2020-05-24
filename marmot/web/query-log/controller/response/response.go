package response

const (
	Failed = 20000
	Success = 10000
)

type Result struct{
	Code int
	Msg string
	Result interface{}
}