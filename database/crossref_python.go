package database

import (
    "regexp"
    "strings"
    "path"
)

func (db *Database) getPythonImportedFiles(ref ImportRef) ([]File, error) {
    r := regexp.MustCompile(`(?:from\s+(\S+)\s)?\s*import\s+((?:\S+\s*)(?:,\s*\S+)*)(?:\sas.*)?$`)
    matches := r.FindStringSubmatch(ref.lineSrc)
    if len(matches) != 3 {
        return []File{}, nil
    }

    from := matches[1]
    from = strings.TrimSpace(from)
    switch {
    case from == ".":
        from = path.Base(ref.dir)
    case from == "..":
        parent, _ := path.Split(path.Clean(ref.dir))
        from = path.Clean(parent)
    case from == "":
    default:
        from = strings.Replace(from, ".", "/", -1)
        from = path.Clean(from)
        files:= (*db).getFilesForModule(from)
        if len(files) != 0 {
            return files, nil
        }
    }

    files := []File{}
    import_ := matches[2]
    for _, imp := range strings.Split(import_, ",") {
        imp = strings.TrimSpace(imp)
        switch {
        case imp == "*":
            files, _ = (*db).getAllFilesInDir(from)
            return files, nil
        case strings.HasPrefix(imp, "..") && from == "":
            parent, _ := path.Split(path.Clean(ref.dir))
            fs := (*db).getFilesForModule(path.Join(parent,
                strings.TrimPrefix(imp, "..")))
            files = append(files, fs...)
        case strings.HasPrefix(imp, ".") && from == "":
            fs := (*db).getFilesForModule(path.Join(ref.dir,
                strings.TrimPrefix(imp, ".")))
            files = append(files, fs...)
        default:
            fs := []File{}
            imp = strings.Replace(imp, ".", "/", -1)
            initFile, isPkg := (*db).findFile(path.Join(ref.dir, imp, "__init__.py"))
            if isPkg {
                fs = (*db).getPythonFilesImportedBy(initFile)
            }
            if len(fs) == 0 {
                fs = (*db).getFilesForModule(path.Join(ref.dir, imp))
            }
            if len(fs) == 0 {
                initFile, isPkg = (*db).findFile(path.Join(imp, "__init__.py"))
                if isPkg {
                    fs = (*db).getPythonFilesImportedBy(initFile)
                }
            }
            if len(fs) == 0 {
                fs = (*db).getFilesForModule(imp)
            }
            files = append(files, fs...)
        }
    }
    return files, nil
}
func (db *Database) getPythonFilesImportedBy(f File) []File {
    files := []File{}
    for _, imp := range (*db).imports {
        if imp.filename == f.name && imp.dir == f.dir {
            fs, err := (*db).getPythonImportedFiles(imp)
            if err == nil {
                files = append(files, fs...)
            }
        }
    }
    return files
}

func (db *Database) findFile(filepath string) (File, bool){
    for fpath, file := range (*db).files {
        if strings.HasSuffix(fpath, filepath) {
            return file, true
        }
    }
    return File{}, false
}

func (db *Database) getAllFilesInDir(dir string) ([]File, bool) {
    files := []File{}
    dirExists := false
    for _, f := range (*db).files {
        if strings.HasSuffix(path.Clean(f.dir), path.Clean(dir)) {
            dirExists = true
            files = append(files, f)
        }
    }
    return files, dirExists
}

func (db *Database) getFilesForModule(modName string) []File {
    files := []File{}
    dir, mod := path.Split(modName)
    dir = path.Clean(dir)
    if dir == "." {
        dir = ""
    }
    for _, f := range (*db).files {
        fileMod := strings.Split(f.name, ".")[0]
        if strings.HasSuffix(path.Clean(f.dir), dir) && mod == fileMod {
            files = append(files, f)
        }
    }
    return files
}
