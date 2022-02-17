package parser

import "fmt"

type Expr interface {
	PrettyPrint() string
}

type Binary struct {
	left     Expr
	operator Token
	right    Expr
}

func (b *Binary) PrettyPrint() string {
	return PrettyPrint(b.operator.lexeme, b.left, b.right)
}

type Grouping struct {
	expr Expr
}

func (g *Grouping) PrettyPrint() string {
	return PrettyPrint("group", g.expr)
}

type Literal struct {
	value interface{}
}

func (l *Literal) PrettyPrint() string {
	return fmt.Sprint(l.value)
}

type Unary struct {
	operator Token
	right    Expr
}

func (u *Unary) PrettyPrint() string {
	return PrettyPrint(u.operator.lexeme, u.right)
}

func PrettyPrint(name string, exprs ...Expr) string {
	res := "("
	res += name
	for _, expr := range exprs {
		res += " "
		res += expr.PrettyPrint()
	}
	res += ")"

	return res
}

func Test() {
	expr := Binary{
		&Unary{
			Token{TOK_MINUS, "-", nil, 1},
			&Literal{123},
		},
		Token{TOK_STAR, "*", nil, 1},
		&Grouping{
			&Literal{45.67},
		},
	}

	fmt.Println(PrettyPrint("", &expr))
}
