package ports

type DBPort interface {
	CloseDBConntection()
	AddToHistory(answer int32, operation string) error
}
