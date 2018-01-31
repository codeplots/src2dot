package database

import (
	"regexp"
	"strings"
)

const (
	SCOPE_CLASS_PREFIX  string = "class:"
	SCOPE_STRUCT_PREFIX string = "struct:"
)

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

func (db *Database) GetClasses() []ClassRef {
	return (*db).classes
}

func (db *Database) GetCaller(ref FuncCallRef) FuncRef {
	return FuncRef{}
}

func (db *Database) GetCallee(ref FuncCallRef) FuncRef {
	return FuncRef{}
}

func (db *Database) GetMethods(ref ClassRef) []FuncRef {
	methods := []FuncRef{}
	for _, f := range (*db).functions {
		switch {
		case f.dir == ref.dir && f.scope == SCOPE_CLASS_PREFIX+ref.symbol:
			methods = append(methods, f)
		case ref.language == GO:
			r := regexp.MustCompile(`func\s(\s*\S+\s*\*?\s*` + regexp.QuoteMeta(ref.symbol) + `\s*)`)
			if r.MatchString(f.lineSrc) {
				methods = append(methods, f)
			}
		}
	}
	return methods
}

func (db *Database) GetMembers(ref ClassRef) []MemberRef {
	members := []MemberRef{}
	for _, m := range (*db).members {
		switch {
		case m.dir == ref.dir && m.scope == SCOPE_CLASS_PREFIX+ref.symbol:
			members = append(members, m)
		case ref.language == GO:
			if m.dir == ref.dir &&
				(m.scope == SCOPE_STRUCT_PREFIX+ref.symbol ||
					(strings.HasPrefix(m.scope, SCOPE_STRUCT_PREFIX) && strings.HasSuffix(m.scope, ("."+ref.symbol)))) {
				members = append(members, m)
			}
		}
	}
	return members
}

func (db *Database) GetParents(ref ClassRef) []ClassRef {
	parents := []ClassRef{}
	for _, c := range (*db).classes {
		if c.symbol == ref.inherits {
			parents = append(parents, c)
		}
	}
	return parents
}

func (db *Database) GetChildren(ref ClassRef) []ClassRef {
	return []ClassRef{}
}

func (db *Database) GetAssociates(ref ClassRef) []ClassRef {
	assoc := []ClassRef{}
	signConcat := ""
	for _, m := range (*db).GetMethods(ref) {
		signConcat += m.signature
	}
	for _, c := range (*db).classes {
		if strings.Contains(signConcat, c.symbol) {
			assoc = append(assoc, c)
		}
	}
	return assoc
}

func (db *Database) GetAggregations(ref ClassRef) []ClassRef {
	aggr := []ClassRef{}
	memberSrcConcat := ""
	memberTypeConcat := ""
	for _, m := range (*db).GetMembers(ref) {
		memberSrcConcat += m.lineSrc
		memberTypeConcat += m.typeref
	}
	for _, c := range (*db).classes {
		if strings.Contains(memberSrcConcat, c.symbol) ||
			strings.Contains(memberTypeConcat, c.symbol) {
			aggr = append(aggr, c)
			continue
		}
		r := regexp.MustCompile(`\W` + regexp.QuoteMeta(c.symbol) + `\W`)
		if ref.language == GO && r.MatchString(ref.lineSrc) && ref.symbol != c.symbol {
			aggr = append(aggr, c)
		}
	}
	return aggr
}
