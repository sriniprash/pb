package store

import (
	"io/ioutil"
	"os"
)
type FileStore struct{
	RootDir string
}

func (self FileStore) Init() (error) {
	// Ensure if the drectory exists.
	return os.MkdirAll(self.RootDir, 0755)
}

func (self FileStore) Get(pasteID string) ([]byte, error) {
	data, err := ioutil.ReadFile(self.RootDir + "/" + pasteID)
	return data, err
}

func (self FileStore) Create(pasteID string, data []byte) error {
	err := ioutil.WriteFile(self.RootDir + "/" + pasteID, data, 0644)
	return err
}

func (self FileStore) Update(pasteID string, data []byte) error {
	err := ioutil.WriteFile(self.RootDir + "/" + pasteID, data, 0644)
	return err
}