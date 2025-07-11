package aetherc

type CFunction struct {
	Name       string
	ReturnType string
	Parameters []CParameter
	Header     string
	Library    string
	Comment    string
}

type CParameter struct {
	Name string
	Type string
}

type CHeader struct {
	Name      string
	Path      string
	Functions []CFunction
	Constants []CConstant
	Types     []CType
	Libraries []string
}

type CConstant struct {
	Name  string
	Value string
	Type  string
}

type CType struct {
	Name   string
	Type   string
	Fields []CField
}

type CField struct {
	Name string
	Type string
}

type CBinding struct {
	PackageName string
	Headers     []CHeader
	Functions   []CFunction
	Constants   []CConstant
	Types       []CType
	Libraries   []string
	Includes    []string
}
