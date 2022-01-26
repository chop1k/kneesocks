package protocol

type Builder struct {
}

func NewBuilder() (Builder, error) {
	return Builder{}, nil
}

func (b Builder) BuildResponse(response ResponseChunk) ([]byte, error) {
	return []byte{response.Status, 0, 0, 0, 0, 0, 0}, nil
}
