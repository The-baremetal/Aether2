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

func (f *Function) node()       {}
func (f *Function) statement()  {}
func (f *Function) expression() {}

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

type CComment struct {
	Content string `json:"content"`
}

func (c *CComment) node()      {}
func (c *CComment) statement() {}

type Identifier struct {
	Value    string `json:"value"`
	Type     string `json:"type"`
	IsVararg bool   `json:"is_vararg,omitempty"`
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

type Spread struct {
	Name string
}

func (s *Spread) expression()    {}
func (s *Spread) node()          {}
func (s *Spread) Statement()     {}
func (s *Spread) String() string { return "..." + s.Name }

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
	CCommentKind           NodeKind = "CComment"
	ArrayKind              NodeKind = "Array"
	CallKind               NodeKind = "Call"
	PropertyAccessKind     NodeKind = "PropertyAccess"
	ForKind                NodeKind = "For"
	PartialApplicationKind NodeKind = "PartialApplication"
	MatchKind              NodeKind = "Match"
	CaseKind               NodeKind = "Case"
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

type Match struct {
	Expr  Expression
	Cases []*Case
}

func (m *Match) node()      {}
func (m *Match) statement() {}

type Case struct {
	Pattern Expression
	Body    *Block
}

func (c *Case) node() {}

func matchToASTNode(m *Match) *ASTNode {
	cases := make([]*ASTNode, len(m.Cases))
	for i, c := range m.Cases {
		cases[i] = caseToASTNode(c)
	}
	return &ASTNode{
		NodeKind: MatchKind,
		Left:     expressionToASTNode(m.Expr),
		Inner:    cases,
	}
}

func caseToASTNode(c *Case) *ASTNode {
	return &ASTNode{
		NodeKind: CaseKind,
		Left:     expressionToASTNode(c.Pattern),
		Body:     blockToASTNode(c.Body),
	}
}

func blockToASTNode(b *Block) *ASTNode {
	statements := make([]*ASTNode, len(b.Statements))
	for i, stmt := range b.Statements {
		statements[i] = statementToASTNode(stmt)
	}
	return &ASTNode{
		NodeKind: BlockKind,
		Inner:    statements,
	}
}

func statementToASTNode(s Statement) *ASTNode {
	switch stmt := s.(type) {
	case *Match:
		return matchToASTNode(stmt)
	case *Assignment:
		return &ASTNode{
			NodeKind: AssignmentKind,
			Left:     expressionToASTNode(stmt.Name),
			Right:    expressionToASTNode(stmt.Value),
		}
	case *Function:
		params := make([]*ASTNode, len(stmt.Params))
		for i, param := range stmt.Params {
			params[i] = &ASTNode{
				NodeKind: ParamKind,
				Value:    param.Value,
			}
		}
		return &ASTNode{
			NodeKind: FunctionDeclKind,
			Name:     stmt.Name.Value,
			Params:   params,
			Body:     blockToASTNode(stmt.Body),
		}
	case *StructDef:
		fields := make([]*ASTNode, len(stmt.Fields))
		for i, field := range stmt.Fields {
			fields[i] = &ASTNode{
				NodeKind: ParamKind, // Assuming ParamKind is used for fields
				Value:    field.Value,
			}
		}
		return &ASTNode{
			NodeKind: StructDefKind,
			Name:     stmt.Name.Value,
			Params:   fields,
		}
	case *If:
		return &ASTNode{
			NodeKind: IfKind,
			Left:     expressionToASTNode(stmt.Condition),
			Body:     blockToASTNode(stmt.Consequence),
			Inner:    []*ASTNode{blockToASTNode(stmt.Alternative)},
		}
	case *While:
		return &ASTNode{
			NodeKind: WhileKind,
			Left:     expressionToASTNode(stmt.Condition),
			Body:     blockToASTNode(stmt.Body),
		}
	case *Repeat:
		return &ASTNode{
			NodeKind: RepeatKind,
			Left:     expressionToASTNode(stmt.Count),
			Body:     blockToASTNode(stmt.Body),
		}
	case *For:
		index := &ASTNode{
			NodeKind: ParamKind,
			Value:    stmt.Index.Value,
		}
		value := &ASTNode{
			NodeKind: ParamKind,
			Value:    stmt.Value.Value,
		}
		return &ASTNode{
			NodeKind: ForKind,
			Left:     index,
			Right:    value,
			Body:     blockToASTNode(stmt.Body),
		}
	case *Return:
		return &ASTNode{
			NodeKind: ReturnKind,
			Left:     expressionToASTNode(stmt.Value),
		}
	case *Import:
		return &ASTNode{
			NodeKind: ImportKind,
			Left:     expressionToASTNode(stmt.Name),
			Right:    expressionToASTNode(stmt.As),
		}
	case *CComment:
		return &ASTNode{
			NodeKind: CCommentKind,
			Value:    stmt.Content,
		}
	}
	return nil
}
