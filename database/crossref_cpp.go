package database

import (
	"errors"
	"path"
	"strings"
)

func (db *Database) getCppImportedFiles(ref ImportRef) ([]File, error) {
	files := []File{}
	if ref.type_ == SYSTEM {
		return files, nil
	}
	for filepath, file := range (*db).files {
		switch {
		case filepath == path.Join(ref.dir, ref.symbol):
			return []File{file}, nil
		case strings.HasSuffix(filepath, ref.symbol):
			files = append(files, file)
		}
	}
	if len(files) <= 1 {
		return files, nil
	}
	return files, errors.New("Ambiguous import " + ref.symbol + " in " +
		ref.filename)
}
