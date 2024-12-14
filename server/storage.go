package server

type ServerStorage interface {
	CreateFile([]byte) (string, int64, error)
	ReadFile(string) ([]byte, error)
	DeleteFile(string) error
}
