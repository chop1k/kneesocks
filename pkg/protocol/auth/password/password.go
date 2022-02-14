package password

type RequestChunk struct {
	Version  byte
	Name     string
	Password string
}

type ResponseChunk struct {
	Version byte
	Status  byte
}
