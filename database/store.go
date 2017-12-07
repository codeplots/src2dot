package database

type LanguageType string

const (
	UNKNOWN_LANGUAGE LanguageType = "UNKNOWN_LANGUAGE_TYPE"
	C                LanguageType = "C"
	CPP              LanguageType = "CPP"
	PYTHON           LanguageType = "PYTHON"
	GO               LanguageType = "GO"
	JAVA             LanguageType = "JAVA"
)

type ImportType string

const (
	UNKNOWN_IMPORT ImportType = "UNKNOWN_IMPORT_TYPE"
	SYSTEM         ImportType = "SYSTEM_IMPORT_TYPE"
	LOCAL          ImportType = "LOCAL_IMPORT_TYPE"
)

type AccessType string

const (
	UNKNOWN_ACCESS AccessType = "UNKNOWN_ACCESS_TYPE"
	PRIVATE        AccessType = "PRIVATE"
	PROTECTED      AccessType = "PROTECTED"
	PUBLIC         AccessType = "PUBLIC"
)

type Ref struct {
	symbol   string
	filename string
	module   string
	dir      string
	lineNo   int
	lineSrc  string
	language LanguageType
}

type FuncRef struct {
	Ref
	signature string
	end       int
	scope     string
	typeref   string
}

type FuncCallRef struct {
	Ref
}

type ImportRef struct {
	Ref
	type_ ImportType
}

type ClassRef struct {
	Ref
}

type MemberRef struct {
	Ref
	scope   string
	typeref string
	access  AccessType
}

type PackageRef struct {
	Ref
}

type VariableRef struct {
	Ref
	scope string
}

type File struct {
	name      string
	dir       string
	language  LanguageType
	lineCount int
}

type Dir struct {
	name      string
	parentDir string
}

type Database struct {
	functions []FuncRef
	funcCalls []FuncCallRef
	imports   []ImportRef
	files     map[string]File
	dirs      []Dir
	classes   []ClassRef
	packages  []PackageRef
	members   []MemberRef
}

func (db *Database) GetLangForFile(path string) LanguageType {
	file, found := (*db).files[path]
	if found {
		return file.language
	}
	return UNKNOWN_LANGUAGE
}