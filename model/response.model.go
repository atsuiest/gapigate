package model

type Response struct {
	Code    string      `json:"code"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
	Message string      `json:"message"`
}

type PublicToken struct {
	Token     string `json:"token"`
	PublicKey string `json:"publicKey"`
}
