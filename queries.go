package main 

import (
	"strings"
)

type SymbolKind string

const (
	DeclFunction SymbolKind = "decl.function"
	DeclClass    SymbolKind = "decl.class"
	DeclVariable SymbolKind = "decl.variable"
	DeclMethod   SymbolKind = "decl.method"
	DeclImport	 SymbolKind = "decl.import"
	DeclExport	 SymbolKind = "decl.export"
	RefCall           SymbolKind = "ref.call"
	RefNew            SymbolKind = "ref.new"
	RefProperty       SymbolKind = "ref.property"
	RefIdentifier     SymbolKind = "ref.identifier"
	RefDynamicImport  SymbolKind = "ref.dynamic_import"
	RefImportedSymbol SymbolKind = "ref.imported_symbol"
)

type Symbol struct {
	Name string
	Kind SymbolKind
	StartByte uint32
	EndByte uint32
	StartRow uint32
	StartCol uint32
	EndRow uint32
	EndCol uint32
}

type nodeRange struct {
	Start uint32
	End uint32
}

func symbolKey(s Symbol) string {
	return string(s.Kind) + ":" + s.Name + ":" + uint32ToString(s.StartByte) + ":" + uint32ToString(s.EndByte)
}

func rangeKey(start, end uint32) string {
	return uint32ToString(start) + ":" + uint32ToString(end)
}

func uint32ToString(n uint32) string {
	return string(n)
}

func declarationPriority(kind SymbolKind) int {
	switch kind {
	case DeclClass:
		return 100
	case DeclFunction:
		return 90
	case DeclMethod:
		return 80
	case DeclImport:
		return 70
	case DeclExport:
		return 60
	case DeclVariable:
		return 50
	default:
		return 0
	}
}

func isDeclaration(kind SymbolKind) bool {
	return strings.HasPrefix(string(kind), "decl.")
}

func isReference(kind SymbolKind) bool {
	return strings.HasPrefix(string(kind), "ref.")
}
