package ports

type ArithmeticPorts interface {
	Addition(int32, int32) (int32, error)
	Subtraction(int32, int32) (int32, error)
	Multiplication(int32, int32) (int32, error)
	Division(int32, int32) (int32, error)
}
