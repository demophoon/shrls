package server

type ServerStorage interface {
	CreateFile([]byte) (string, error)
	ReadFile(string) ([]byte, error)
	DeleteFile(string) error
}
