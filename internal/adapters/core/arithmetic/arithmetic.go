package arithmetic

type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (aa *Adapter) Addition(a, b int32) (int32, error) {
	return a + b, nil
}

func (aa *Adapter) Subtraction(a, b int32) (int32, error) {
	return a - b, nil
}

func (aa *Adapter) Multiplication(a, b int32) (int32, error) {
	return a * b, nil
}

func (aa *Adapter) Division(a, b int32) (int32, error) {
	return a / b, nil
}
