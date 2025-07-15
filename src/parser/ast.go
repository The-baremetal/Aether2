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
	Names []*Identifier `json:"names"`
	Value Expression    `json:"value"`
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
	Name   *Identifier `json:"name"`
	Fields []*Field    `json:"fields"`
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

type Package struct {
	Name *Identifier `json:"name"`
}

func (p *Package) node()      {}
func (p *Package) statement() {}

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

type ArrayIndex struct {
	Array Expression `json:"array"`
	Index Expression `json:"index"`
}

func (a *ArrayIndex) node()       {}
func (a *ArrayIndex) expression() {}
func (a *ArrayIndex) statement()  {}

type StructInstantiation struct {
	TypeName *Identifier            `json:"type_name"`
	Fields   map[string]Expression `json:"fields"`
}

func (s *StructInstantiation) node()       {}
func (s *StructInstantiation) expression() {}
func (s *StructInstantiation) statement()  {}

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
	PackageKind            NodeKind = "Package"
	CCommentKind           NodeKind = "CComment"
	ArrayKind              NodeKind = "Array"
	CallKind               NodeKind = "Call"
	PropertyAccessKind     NodeKind = "PropertyAccess"
	ArrayIndexKind         NodeKind = "ArrayIndex"
	StructInstantiationKind NodeKind = "StructInstantiation"
	ForKind                NodeKind = "For"
	PartialApplicationKind NodeKind = "PartialApplication"
	MatchKind              NodeKind = "Match"
	CaseKind               NodeKind = "Case"
	BreakKind              NodeKind = "Break"
	ContinueKind           NodeKind = "Continue"
	SpreadKind             NodeKind = "Spread"
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
	if e == nil {
		// Nil expression, return nil to avoid panic
		return nil
	}
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
		var left *ASTNode
		if expr.Object != nil {
			left = expressionToASTNode(expr.Object)
		}
		var property *ASTNode
		if expr.Property != nil {
			property = expressionToASTNode(expr.Property)
		}
		return &ASTNode{
			NodeKind: PropertyAccessKind,
			Left:     left,
			Inner:    []*ASTNode{property},
		}
	case *ArrayIndex:
		return &ASTNode{
			NodeKind: ArrayIndexKind,
			Left:     expressionToASTNode(expr.Array),
			Right:    expressionToASTNode(expr.Index),
		}
	case *StructInstantiation:
		fields := make([]*ASTNode, 0, len(expr.Fields))
		for name, value := range expr.Fields {
			var right *ASTNode
			if value != nil {
				right = expressionToASTNode(value)
			}
			fieldNode := &ASTNode{
				NodeKind: "Field",
				Name:     name,
				Right:    right,
			}
			fields = append(fields, fieldNode)
		}
		var typeName string
		if expr.TypeName != nil {
			typeName = expr.TypeName.Value
		}
		return &ASTNode{
			NodeKind: StructInstantiationKind,
			Name:     typeName,
			Inner:    fields,
		}
	case *PartialApplication:
		return &ASTNode{
			NodeKind: PartialApplicationKind,
			Left:     expressionToASTNode(expr.Function),
			Inner:    mapArgsToASTNodes(expr.Args),
		}
	case *Spread:
		return &ASTNode{
			NodeKind: SpreadKind,
			Value:    expr.Name,
		}
	}
	return nil
}

func mapArgsToASTNodes(args []Expression) []*ASTNode {
	result := make([]*ASTNode, 0, len(args))
	for _, a := range args {
		if a == nil {
			result = append(result, nil)
		} else {
			result = append(result, expressionToASTNode(a))
		}
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

type Break struct{}

func (b *Break) node()      {}
func (b *Break) statement() {}

type Continue struct{}

func (c *Continue) node()      {}
func (c *Continue) statement() {}

type ExpressionStatement struct {
	Expr Expression
}

func (e *ExpressionStatement) node()      {}
func (e *ExpressionStatement) statement() {}

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
	if b == nil {
		return nil
	}
	statements := make([]*ASTNode, len(b.Statements))
	for i, stmt := range b.Statements {
		if stmt == nil {
			statements[i] = nil
		} else {
			statements[i] = statementToASTNode(stmt)
		}
	}
	return &ASTNode{
		NodeKind: BlockKind,
		Inner:    statements,
	}
}

func statementToASTNode(s Statement) *ASTNode {
	if s == nil {
		// Nil statement, return nil to avoid panic
		return nil
	}
	switch stmt := s.(type) {
	case *Match:
		return matchToASTNode(stmt)
	case *Assignment:
		var names []*ASTNode
		for _, n := range stmt.Names {
			if n == nil {
				names = append(names, nil)
			} else {
				names = append(names, expressionToASTNode(n))
			}
		}
		return &ASTNode{
			NodeKind: AssignmentKind,
			Params:   names, // or Inner: names,
			Right:    expressionToASTNode(stmt.Value),
		}
	case *Function:
		params := make([]*ASTNode, len(stmt.Params))
		for i, param := range stmt.Params {
			if param == nil {
				params[i] = nil
			} else {
				params[i] = &ASTNode{
					NodeKind: ParamKind,
					Value:    param.Value,
				}
			}
		}
		var name string
		if stmt.Name != nil {
			name = stmt.Name.Value
		}
		return &ASTNode{
			NodeKind: FunctionDeclKind,
			Name:     name,
			Params:   params,
			Body:     blockToASTNode(stmt.Body),
		}
	case *StructDef:
		fields := make([]*ASTNode, len(stmt.Fields))
		for i, field := range stmt.Fields {
			if field == nil {
				fields[i] = nil
			} else {
				fields[i] = &ASTNode{
					NodeKind: ParamKind,
					Name:     field.Name.Value,
					Value:    field.Type,
				}
			}
		}
		var name string
		if stmt.Name != nil {
			name = stmt.Name.Value
		}
		return &ASTNode{
			NodeKind: StructDefKind,
			Name:     name,
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
		var index *ASTNode
		if stmt.Index != nil {
			index = &ASTNode{
				NodeKind: ParamKind,
				Value:    stmt.Index.Value,
			}
		}
		var value *ASTNode
		if stmt.Value != nil {
			value = &ASTNode{
				NodeKind: ParamKind,
				Value:    stmt.Value.Value,
			}
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
		var left, right *ASTNode
		if stmt.Name != nil {
			left = expressionToASTNode(stmt.Name)
		}
		if stmt.As != nil {
			right = expressionToASTNode(stmt.As)
		}
		return &ASTNode{
			NodeKind: ImportKind,
			Left:     left,
			Right:    right,
		}
	case *Package:
		var value string
		if stmt.Name != nil {
			value = stmt.Name.Value
		}
		return &ASTNode{
			NodeKind: PackageKind,
			Value:    value,
		}
	case *CComment:
		return &ASTNode{
			NodeKind: CCommentKind,
			Value:    stmt.Content,
		}
	case *Break:
		return &ASTNode{
			NodeKind: BreakKind,
		}
	case *Continue:
		return &ASTNode{
			NodeKind: ContinueKind,
		}
	case *ExpressionStatement:
		if stmt.Expr == nil {
			return nil
		}
		return expressionToASTNode(stmt.Expr)
	}
	return nil
}

type Field struct {
	Name *Identifier `json:"name"`
	Type string      `json:"type"`
}
