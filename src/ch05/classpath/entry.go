package classpath

import (
	"os"
	"strings"
)

const pathListSeperator = string(os.PathListSeparator)

type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeperator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildwardEntry(path)
	}
	if strings.HasSuffix(strings.ToLower(path), ".jar") ||
		strings.HasSuffix(strings.ToLower(path), ".zip") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}
