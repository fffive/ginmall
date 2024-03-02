package serializer

type Response struct {
	Status int `json:"status"`
	Data interface{} `json:"data"`
	Msg string `json:"message"`
	Error string `json:"error"`
}