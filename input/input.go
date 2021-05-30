package input

// 数据生产
type Input interface {
	retrieveData() ([]byte, error)
}

func NewInput(in string) Input {
	if in == "file" {
		return NewFileInput()
	}
	return nil
}
