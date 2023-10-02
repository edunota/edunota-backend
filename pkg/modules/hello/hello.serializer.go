package hello

type SayHelloSerializer struct {
	Status uint8  `json:"status"`
	Data   string `json:"data"`
}
