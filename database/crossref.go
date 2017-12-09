package database

import (
    "path"
    "strings"
    "errors"
)

func (db *Database) getCppImportedFiles(ref ImportRef) ([]File, error) {
    files := []File{}
    for filepath, file := range (*db).files {
        switch {
        case filepath == path.Join(ref.dir, ref.symbol):
            return []File{ file }, nil
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

func (db *Database) GetImportedFiles(ref ImportRef) ([]File, error) {
        if ref.type_ == SYSTEM {
            return []File{}, nil
        }
        switch ref.language {
        case CPP:
            return (*db).getCppImportedFiles(ref)
        }
	return []File{}, nil
}

func (db *Database) GetImports() []ImportRef {
        return (*db).imports
}

func (db *Database) GetCaller(ref FuncCallRef) FuncRef {
	return FuncRef{}
}

func (db *Database) GetCallee(ref FuncCallRef) FuncRef {
	return FuncRef{}
}

func (db *Database) GetMembers(ref ClassRef) []MemberRef {
	return []MemberRef{}
}

func (db *Database) GetMethods(ref ClassRef) []FuncRef {
	return []FuncRef{}
}

func (db *Database) GetParents(ref ClassRef) []ClassRef {
	return []ClassRef{}
}

func (db *Database) GetChildren(ref ClassRef) []ClassRef {
	return []ClassRef{}
}

func (db *Database) GetAssociates(ref ClassRef) []ClassRef {
	return []ClassRef{}
}
