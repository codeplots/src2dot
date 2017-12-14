package database

func (db *Database) GetImportedFiles(ref ImportRef) ([]File, error) {
        switch ref.language {
        case CPP:
            return (*db).getCppImportedFiles(ref)
        case GO:
            return (*db).getGoImportedFiles(ref)
        case PYTHON:
            return (*db).getPythonImportedFiles(ref)
        }
	return []File{}, nil
}

func (db *Database) GetImports() []ImportRef {
    imports := []ImportRef{}
    for _, importRef := range (*db).imports {
        imports = append(imports, importRef)
    }
    return imports
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
