package database

import ()

func (db *Database) GetImportedFiles(ref ImportRef) []File {
	return []File{}
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
