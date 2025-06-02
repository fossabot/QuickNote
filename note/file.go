package note

import (
	"os"
	"path/filepath"
)

func (n *Note) Write() error {
	data, err := n.Encode()
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(".", n.ID), data, 0644)
}

func (n *Note) Read() error {
	data, err := os.ReadFile(filepath.Join(".", n.ID))
	if err != nil {
		return err
	}
	return n.Decode(data)
}
