package database

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"reflect"
	"strconv"
	"strings"
)

const (
	LANGUAGE_TAG    string = "language"
	END_TAG         string = "end"
	LINE_TAG        string = "line"
	SIGNATURE_TAG   string = "signature"
	SCOPE_TAG       string = "scope"
	ACCESS_TAG      string = "access"
	TYPEREF_TAG     string = "typeref"
	KIND_TAG        string = "kind"
	REFERENCE_TAG   string = "reference"
	EXTRAS_TAG      string = "extras"
	ROLE_TAG        string = "role"
	INHERITS_TAG    string = "inherits"
	TYPENAME_PREFIX string = "typename:"
)

// pseudo constant mapping between ctags language literals and LanguageType
var LANGUAGE_TAG_MAP = map[string]LanguageType{
	"C":      C,
	"C++":    CPP,
	"Python": PYTHON,
	"Go":     GO,
	"Java":   JAVA,
}

// pseudo constant mapping between ctag kind literals and a dummy of the
// corresponding type
var KIND_TAG_MAP = map[string]interface{}{
	"class":                     ClassRef{},
	string(GO) + "struct":       ClassRef{},
	"func":                      FuncRef{},
	"function":                  FuncRef{},
	string(PYTHON) + "member":   FuncRef{},
	"member":                    MemberRef{},
	string(PYTHON) + "variable": MemberRef{},
	"file": File{},
}

var ACCESS_TAG_MAP = map[string]AccessType{
	"private":   PRIVATE,
	"protected": PROTECTED,
	"public":    PUBLIC,
}

var REFERENCE_TAG_MAP = map[string]ImportType{
	"local":  LOCAL,
	"system": SYSTEM,
}

func getAccessType(tag string) AccessType {
	access, found := ACCESS_TAG_MAP[tag]
	if found {
		return access
	}
	return UNKNOWN_ACCESS
}

func getLanguageType(tag string) LanguageType {
	lang, found := LANGUAGE_TAG_MAP[tag]
	if found {
		return lang
	}
	return UNKNOWN_LANGUAGE
}

func FromCtags(ctagsFile string) Database {
	tags, err := ioutil.ReadFile(ctagsFile)
	if err != nil {
		log.Fatal(err)
	}
	tagLines := strings.Split(string(tags), "\n")
	db := Database{}
	db.files = make(map[string]File)
	db.imports = make(map[string]ImportRef)
	for _, line := range tagLines {
		// Split line into [symbol filename rest]
		cols := strings.SplitN(line, "\t", 3)
		if len(cols) < 3 {
			continue
		}

		ref := Ref{}
		ref.symbol = cols[0]
		ref.dir, ref.filename = path.Split(cols[1])
		ref.module = strings.Split(ref.filename, ".")[0]

		// Split rest of line into [lineSrc rest]
		cols = strings.SplitN(cols[2], ";\"\t", 2)
		if len(cols) < 2 {
			continue
		}

		ref.lineSrc = strings.TrimPrefix(
			strings.TrimSuffix(cols[0], "$/"), "/^")

		cols = strings.Split(cols[1], "\t")

		fields := map[string]string{}
		for _, col := range cols {
			parts := strings.SplitN(col, ":", 2)
			if len(parts) < 2 {
				continue
			}
			fields[parts[0]] = parts[1]
		}

		ref.lineNo, _ = strconv.Atoi(fields[LINE_TAG])
		ref.language = getLanguageType(fields[LANGUAGE_TAG])

		extras, found := fields[EXTRAS_TAG]
		if found && extras == REFERENCE_TAG {
			importType, found := REFERENCE_TAG_MAP[fields[ROLE_TAG]]
			if !found {
				importType = UNKNOWN_IMPORT
			}
			importRef := ImportRef{
				Ref:   ref,
				type_: importType,
			}
			db.imports[fmt.Sprintf(
				"%s:%i:%s",
				path.Join(ref.dir, ref.filename),
				ref.lineNo,
				ref.symbol)] = importRef
			continue
		}

		// Prioritiy for special 'language' + fields['kind_tag'] rule
		// e.g. PYTHONmember, GOfunc
		typeDummy, found := KIND_TAG_MAP[string(ref.language)+fields[KIND_TAG]]
		if !found {
			// fallback to generic 'kind_tag' rule
			typeDummy, found = KIND_TAG_MAP[fields[KIND_TAG]]
		}
		if !found {
			continue
		}

		switch reflect.TypeOf(typeDummy) {
		case reflect.TypeOf(FuncRef{}):
			funcRef := FuncRef{
				Ref:       ref,
				signature: fields[SIGNATURE_TAG],
				scope:     fields[SCOPE_TAG],
				typeref: strings.TrimPrefix(fields[TYPEREF_TAG],
					TYPENAME_PREFIX),
			}
			funcRef.end, _ = strconv.Atoi(fields[END_TAG])
			db.functions = append(db.functions, funcRef)
		case reflect.TypeOf(MemberRef{}):
			memberRef := MemberRef{
				Ref:   ref,
				scope: fields[SCOPE_TAG],
				typeref: strings.TrimPrefix(fields[TYPEREF_TAG],
					TYPENAME_PREFIX),
				access: getAccessType(fields[ACCESS_TAG]),
			}
			db.members = append(db.members, memberRef)
		case reflect.TypeOf(ClassRef{}):
			parent, _ := fields[INHERITS_TAG]
			if ref.language == GO {
				ref.lineSrc = getSrcBlock(ref.lineNo,
					path.Join(ref.dir, ref.filename), ctagsFile)
			}
			classRef := ClassRef{
				Ref:      ref,
				inherits: parent,
			}
			db.classes = append(db.classes, classRef)
		case reflect.TypeOf(File{}):
			file := File{
				name:     ref.filename,
				dir:      ref.dir,
				language: ref.language,
			}
			file.lineCount, _ = strconv.Atoi(fields[END_TAG])
			db.files[path.Join(ref.dir, ref.filename)] = file
		}
	}
	return db
}
func getSrcBlock(startLine int, filepath string, ctagsFile string) string {
	if !path.IsAbs(filepath) {
		base, _ := path.Split(ctagsFile)
		filepath = path.Join(base, filepath)
	}
	src, _ := ioutil.ReadFile(filepath)
	srcLines := strings.Split(string(src), "\n")
	nb_open := 0
	for i, l := range srcLines[startLine-1:] {
		for _, c := range l {
			switch c {
			case '{':
				nb_open += 1
			case '}':
				nb_open -= 1
			}
		}
		if nb_open <= 0 {
			return strings.Join(srcLines[startLine-1:startLine+i+1], "\n")
		}
	}
	return srcLines[startLine]
}
