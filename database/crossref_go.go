package database

import (
    "regexp"
    "path"
)

func (db *Database) getGoImportedFiles(ref ImportRef) ([]File, error) {
    files := []File{}
    r := regexp.MustCompile(`"((?:\S+/)*` + ref.symbol + `)"`)
    matches := r.FindStringSubmatch(ref.lineSrc)
    if !(len(matches) == 2) {
        return files, nil
    }
    for _, file := range (*db).files {
        switch {
        case path.Base(file.dir) == ref.symbol:
            files = append(files, file)
        }
    }
    return files, nil
}

