package note

import (
	"os"
	"path/filepath"
)

func (n *Note) Delete() error {
	return os.Remove(filepath.Join(".", n.ID))
}
