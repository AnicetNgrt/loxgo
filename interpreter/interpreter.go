package interpreter

type Val interface{}

func Interpret(expr Expr) (Val, Error) {
	return expr.Interpret()
}

func (u *Unary) Interpret() (Val, Error) {
	right, _ := u.right.Interpret()

	switch u.operator.tType {
	case TOK_MINUS:
		right, _ := right.(float64)
		return -right, nil
	case TOK_BANG:
		return !isTruthy(right), nil
	}

	return nil, nil
}

func (l *Literal) Interpret() (Val, Error) {
	return l.value, nil
}

func (g *Grouping) Interpret() (Val, Error) {
	return g.expr.Interpret()
}

func (b *Binary) Interpret() (Val, Error) {
	leftVal, _ := b.left.Interpret()
	rightVal, _ := b.right.Interpret()
	opType := b.operator.tType

	{
		left, leftIsFloat := leftVal.(float64)
		right, rightIsFloat := rightVal.(float64)

		if leftIsFloat && rightIsFloat {
			switch opType {
			case TOK_MINUS:
				return left - right, nil
			case TOK_PLUS:
				return left + right, nil
			case TOK_SLASH:
				return left / right, nil
			case TOK_STAR:
				return left * right, nil
			case TOK_LESS:
				return left < right, nil
			case TOK_GREATER:
				return left > right, nil
			case TOK_LESS_EQUAL:
				return left <= right, nil
			case TOK_GREATER_EQUAL:
				return left >= right, nil
			case TOK_EQUAL_EQUAL:
				return left == right, nil
			case TOK_BANG_EQUAL:
				return left != right, nil
			}
		}
	}

	{
		left, leftIsString := leftVal.(string)
		right, rightIsString := rightVal.(string)

		if leftIsString && rightIsString {
			switch opType {
			case TOK_PLUS:
				return left + right, nil
			case TOK_EQUAL_EQUAL:
				return left == right, nil
			case TOK_BANG_EQUAL:
				return left != right, nil
			}
		}
	}

	switch b.operator.tType {
	case TOK_BANG_EQUAL:
		return !isEqual(leftVal, rightVal), nil
	case TOK_EQUAL_EQUAL:
		return isEqual(leftVal, rightVal), nil
	}

	return nil, nil
}

// Everything truthy except false and nil. We're no JS here ;)
func isTruthy(val Val) bool {
	if val == nil {
		return false
	}
	valAsBool, ok := val.(bool)
	if ok {
		return valAsBool
	}
	return true
}

func isEqual(left Val, right Val) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil {
		return false
	}

	return left == right
}
