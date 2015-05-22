package store

type store interface {
	Init() error
	Get(pasteID string) ([]byte, error)
	Create(pasteID string, data []byte) error
	Update(pasteID string, data []byte) error
}