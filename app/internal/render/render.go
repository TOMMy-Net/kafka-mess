package render

var (
	GoodMsg = "the message was successfully delivered"
)

type Answer struct {
	Status  string `json:"status"`
	Message string `json:"msg"`
}
