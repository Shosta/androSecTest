package folder

import (
	"os"
)

// Delete : Delete a folder and all of its children.
func Delete(path string) {
	os.RemoveAll(path)
}
