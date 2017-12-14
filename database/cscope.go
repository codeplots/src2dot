package database

import (
    "fmt"
	"io/ioutil"
	"log"
	"path"
	"strconv"
	"strings"
)

const (
	FILE_PREFIX              = "\t@"
	FUNC_DEF_PREFIX          = "\t$"
	FUNC_CALL_PREFIX         = "\t`"
	FUNC_END_PREFIX          = "\t}"
	DEFINE_PREFIX            = "\t#"
	DEFINE_END_PREFIX        = "\t)"
	IMPORT_PREFIX            = "\t~"
	LOCAL_IMPORT_PREFIX      = "\t~\""
	SYS_IMPORT_PREFIX        = "\t~<"
	ASSIGN_PREFIX            = "\t="
	ENUM_ST_UN_END_PREFIX    = "\t;"
	CLASS_DEF_PREFIX         = "\tc"
	ENUM_DEF_PREFIX          = "\te"
	GLOB_DEF_PREFIX          = "\tg"
	LOCAL_DEF_PREFIX         = "\tl"
	GLOB_EN_ST_UN_MEM_PREFIX = "\tm"
	FUNC_PARAM_PREFIX        = "\tp"
	STRUCT_PREFIX            = "\ts"
	TYPEDEF_PREFIX           = "\tt"
	UNION_PREFIX             = "\tu"
)

func getImportRef(line string) (string, ImportType) {
	switch {
	case strings.HasPrefix(line, LOCAL_IMPORT_PREFIX):
		return strings.TrimPrefix(line, LOCAL_IMPORT_PREFIX), LOCAL
	case strings.HasPrefix(line, SYS_IMPORT_PREFIX):
		return strings.TrimPrefix(line, SYS_IMPORT_PREFIX), SYSTEM
	}
	return strings.TrimPrefix(line, IMPORT_PREFIX), UNKNOWN_IMPORT
}

func (db *Database) AddCscope(cscopeFile string) {
	cscope, err := ioutil.ReadFile(cscopeFile)
	if err != nil {
		log.Fatal(err)
	}
	cscopeLines := strings.Split(string(cscope), "\n")

	var filename, dir, module, lineSrc string
        lineNo := 0
	for i, line := range cscopeLines {
		switch {
		case strings.HasPrefix(line, IMPORT_PREFIX):
			symbol, type_ := getImportRef(line)
			ref := Ref{
				symbol:   symbol,
				filename: filename,
				module:   module,
				dir:      dir,
				lineNo:   lineNo,
                                lineSrc:  lineSrc,
				language: (*db).GetLangForFile(
					path.Join(dir, filename)),
			}
			importRef := ImportRef{
				Ref:   ref,
				type_: type_,
			}

                        // add the import reference only, if there is none with
                        // the same (dir, filename, lineNo, symbol) in the
                        // database
                        _, found := (*db).imports[fmt.Sprintf( "%s:%i:%s",
                            path.Join(dir, filename),
                            lineNo,
                            symbol)]

                            if !found {
                                (*db).imports[fmt.Sprintf( "%s:%i:%s",
                                path.Join(dir, filename),
                                lineNo,
                                symbol)] = importRef
                            }

		case strings.HasPrefix(line, FUNC_CALL_PREFIX):
			ref := Ref{
				symbol: strings.TrimPrefix(
					line, FUNC_CALL_PREFIX),
				filename: filename,
				module:   module,
				dir:      dir,
				lineNo:   lineNo,
                                lineSrc:  lineSrc,
				language: (*db).GetLangForFile(
					path.Join(dir, filename)),
			}
			funcCallRef := FuncCallRef{
				Ref: ref,
			}
			(*db).funcCalls = append((*db).funcCalls, funcCallRef)
		case strings.HasPrefix(line, FILE_PREFIX):
			dir, filename = path.Split(
				strings.TrimPrefix(line, FILE_PREFIX))
			module = strings.Split(filename, ".")[0]
		case len(line) > 0 && line[0] >= '0' && line[0] <= '9':
			lineNo, _ = strconv.Atoi(strings.Split(line, " ")[0])
                        if lineNo >= 1 {
                            lineEnd := findLineEnd(cscopeLines[i+1:])
                            if lineEnd != 0 {
                                lineSrc = concatLineSrc(cscopeLines[i:lineEnd+i])
                            }
                        }
		}
	}
}

func findLineEnd(lines []string) int {
    for i, l := range lines {
        switch {
        case len(l) > 0 && l[0] >= '0' && l[0] <= '9':
            return i
        case strings.HasPrefix(l, FILE_PREFIX):
            return 0
        }
    }
    return 0
}

func concatLineSrc(lines []string) string {
    lines[0] = strings.Join(strings.Split(lines[0], " ")[1:], " ")
    concat := strings.Join(lines, "")
    concat = strings.Replace(concat, FUNC_DEF_PREFIX, "", -1)
    concat = strings.Replace(concat, FUNC_CALL_PREFIX, "", -1)
    concat = strings.Replace(concat, FUNC_END_PREFIX, "", -1)
    concat = strings.Replace(concat, DEFINE_PREFIX, "", -1)
    concat = strings.Replace(concat, DEFINE_END_PREFIX, "", -1)
    concat = strings.Replace(concat, LOCAL_IMPORT_PREFIX, "", -1)
    concat = strings.Replace(concat, SYS_IMPORT_PREFIX, "", -1)
    concat = strings.Replace(concat, IMPORT_PREFIX, "", -1)
    concat = strings.Replace(concat, ASSIGN_PREFIX, "", -1)
    concat = strings.Replace(concat, ENUM_ST_UN_END_PREFIX, "", -1)
    concat = strings.Replace(concat, CLASS_DEF_PREFIX, "", -1)
    concat = strings.Replace(concat, ENUM_DEF_PREFIX, "", -1)
    concat = strings.Replace(concat, GLOB_DEF_PREFIX, "", -1)
    concat = strings.Replace(concat, LOCAL_DEF_PREFIX, "", -1)
    concat = strings.Replace(concat, GLOB_EN_ST_UN_MEM_PREFIX, "", -1)
    concat = strings.Replace(concat, FUNC_PARAM_PREFIX, "", -1)
    concat = strings.Replace(concat, STRUCT_PREFIX, "", -1)
    concat = strings.Replace(concat, TYPEDEF_PREFIX, "", -1)
    concat = strings.Replace(concat, UNION_PREFIX, "", -1)
    concat = strings.TrimSpace(concat)
    return concat
}
