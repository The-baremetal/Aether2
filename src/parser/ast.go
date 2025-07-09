package parser

type Node interface {
	node()
}

type Statement interface {
	Node
	statement()
}

type Expression interface {
	Node
	expression()
}

type Program struct {
	Statements []Statement `json:"statements"`
}

func (p *Program) node() {}

type Assignment struct {
	Name  *Identifier `json:"name"`
	Value Expression  `json:"value"`
}

func (a *Assignment) node()      {}
func (a *Assignment) statement() {}

type Function struct {
	Name   *Identifier   `json:"name"`
	Params []*Identifier `json:"params"`
	Body   *Block        `json:"body"`
}

func (f *Function) node()      {}
func (f *Function) statement() {}

type StructDef struct {
	Name   *Identifier   `json:"name"`
	Fields []*Identifier `json:"fields"`
}

func (s *StructDef) node()      {}
func (s *StructDef) statement() {}

type If struct {
	Condition   Expression `json:"condition"`
	Consequence *Block     `json:"consequence"`
	Alternative *Block     `json:"alternative"`
}

func (i *If) node()      {}
func (i *If) statement() {}

type While struct {
	Condition Expression `json:"condition"`
	Body      *Block     `json:"body"`
}

func (w *While) node()      {}
func (w *While) statement() {}

type Repeat struct {
	Count Expression `json:"count"`
	Body  *Block     `json:"body"`
}

func (r *Repeat) node()      {}
func (r *Repeat) statement() {}

type For struct {
	Index    *Identifier `json:"index,omitempty"`
	Value    *Identifier `json:"value"`
	Iterable Expression  `json:"iterable"`
	Body     *Block      `json:"body"`
}

func (f *For) node()      {}
func (f *For) statement() {}

type Block struct {
	Statements []Statement `json:"statements"`
}

func (b *Block) node()       {}
func (b *Block) statement()  {}
func (b *Block) expression() {}

type Return struct {
	Value Expression `json:"value"`
}

func (r *Return) node()      {}
func (r *Return) statement() {}

type Import struct {
	Name *Identifier `json:"name"`
	As   *Identifier `json:"as"`
}

func (i *Import) node()      {}
func (i *Import) statement() {}

type Identifier struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}

func (i *Identifier) node()       {}
func (i *Identifier) expression() {}
func (i *Identifier) statement()  {}

type Literal struct {
	Value interface{} `json:"value"`
}

func (l *Literal) node()       {}
func (l *Literal) expression() {}
func (l *Literal) statement()  {}

type Array struct {
	Elements []Expression `json:"elements"`
}

func (a *Array) node()       {}
func (a *Array) expression() {}

type Call struct {
	Function Expression   `json:"function"`
	Args     []Expression `json:"args"`
	TailCall bool         `json:"tail_call,omitempty"`
}

func (c *Call) node()       {}
func (c *Call) expression() {}
func (c *Call) statement()  {}

type PropertyAccess struct {
	Object   Expression  `json:"object"`
	Property *Identifier `json:"property"`
}

func (p *PropertyAccess) node()       {}
func (p *PropertyAccess) expression() {}
func (p *PropertyAccess) statement()  {}

type PartialApplication struct {
	Function Expression
	Args     []Expression
}

func (p *PartialApplication) node()       {}
func (p *PartialApplication) expression() {}
func (p *PartialApplication) statement()  {}

type NodeKind string

const (
	TranslationUnitKind    NodeKind = "TranslationUnit"
	FunctionDeclKind       NodeKind = "FunctionDecl"
	ParamKind              NodeKind = "Param"
	BlockKind              NodeKind = "Block"
	ReturnKind             NodeKind = "Return"
	BinaryExprKind         NodeKind = "BinaryExpr"
	IdentifierKind         NodeKind = "Identifier"
	LiteralKind            NodeKind = "Literal"
	AssignmentKind         NodeKind = "Assignment"
	StructDefKind          NodeKind = "StructDef"
	IfKind                 NodeKind = "If"
	WhileKind              NodeKind = "While"
	RepeatKind             NodeKind = "Repeat"
	ImportKind             NodeKind = "Import"
	ArrayKind              NodeKind = "Array"
	CallKind               NodeKind = "Call"
	PropertyAccessKind     NodeKind = "PropertyAccess"
	ForKind                NodeKind = "For"
	PartialApplicationKind NodeKind = "PartialApplication"
)

type ASTNode struct {
	NodeKind NodeKind    `json:"kind"`
	ID       string      `json:"id,omitempty"`
	Name     string      `json:"name,omitempty"`
	Params   []*ASTNode  `json:"params,omitempty"`
	Body     *ASTNode    `json:"body,omitempty"`
	Operator string      `json:"operator,omitempty"`
	Left     *ASTNode    `json:"left,omitempty"`
	Right    *ASTNode    `json:"right,omitempty"`
	Value    interface{} `json:"value,omitempty"`
	Inner    []*ASTNode  `json:"inner,omitempty"`
}

func (n *ASTNode) Kind() NodeKind {
	return n.NodeKind
}

func expressionToASTNode(e Expression) *ASTNode {
	switch expr := e.(type) {
	case *Identifier:
		return &ASTNode{
			NodeKind: IdentifierKind,
			Value:    expr.Value,
		}
	case *Literal:
		return &ASTNode{
			NodeKind: LiteralKind,
			Value:    expr.Value,
		}
	case *Array:
		return &ASTNode{
			NodeKind: ArrayKind,
			Inner:    mapArgsToASTNodes(expr.Elements),
		}
	case *Call:
		return &ASTNode{
			NodeKind: CallKind,
			Left:     expressionToASTNode(expr.Function),
			Inner:    mapArgsToASTNodes(expr.Args),
		}
	case *PropertyAccess:
		return &ASTNode{
			NodeKind: PropertyAccessKind,
			Left:     expressionToASTNode(expr.Object),
			Inner:    []*ASTNode{expressionToASTNode(expr.Property)},
		}
	case *PartialApplication:
		return &ASTNode{
			NodeKind: PartialApplicationKind,
			Left:     expressionToASTNode(expr.Function),
			Inner:    mapArgsToASTNodes(expr.Args),
		}
	}
	return nil
}

func mapArgsToASTNodes(args []Expression) []*ASTNode {
	result := make([]*ASTNode, 0, len(args))
	for _, a := range args {
		result = append(result, expressionToASTNode(a))
	}
	return result
}
