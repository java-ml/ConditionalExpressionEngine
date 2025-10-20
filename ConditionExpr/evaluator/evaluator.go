// Package evaluator evaluator/evaluator.go
package evaluator

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"ConditionExpr/parser"

	"github.com/antlr4-go/antlr/v4"
)

// VariableResolver 变量解析器接口
type VariableResolver interface {
	ResolveVariable(name string) interface{}
}

// FunctionCaller 函数调用接口
type FunctionCaller interface {
	CallFunction(name string, args []interface{}) interface{}
}

// DefaultFunctionCaller 默认函数实现
type DefaultFunctionCaller struct{}

func (d *DefaultFunctionCaller) CallFunction(name string, args []interface{}) interface{} {
	switch name {
	case "contains":
		if len(args) != 2 {
			return false
		}
		str := fmt.Sprintf("%v", args[0])
		substr := fmt.Sprintf("%v", args[1])
		return strings.Contains(str, substr)
	case "len":
		if len(args) != 1 {
			return 0
		}
		str := fmt.Sprintf("%v", args[0])
		return len(str)
	case "sqrt":
		if len(args) != 1 {
			return 0.0
		}
		if num, ok := args[0].(float64); ok {
			return math.Sqrt(num)
		}
		return 0.0
	case "upper":
		if len(args) != 1 {
			return ""
		}
		return strings.ToUpper(fmt.Sprintf("%v", args[0]))
	case "lower":
		if len(args) != 1 {
			return ""
		}
		return strings.ToLower(fmt.Sprintf("%v", args[0]))
	}
	return nil
}

// ExpressionEvaluator 表达式求值器
type ExpressionEvaluator struct {
	parser.BaseConditionExprListener
	stack     []interface{}
	variables VariableResolver
	functions FunctionCaller
}

func NewExpressionEvaluator(variables VariableResolver, functions FunctionCaller) *ExpressionEvaluator {
	if functions == nil {
		functions = &DefaultFunctionCaller{}
	}
	return &ExpressionEvaluator{
		stack:     make([]interface{}, 0),
		variables: variables,
		functions: functions,
	}
}

func (e *ExpressionEvaluator) push(value interface{}) {
	e.stack = append(e.stack, value)
}

func (e *ExpressionEvaluator) pop() interface{} {
	if len(e.stack) == 0 {
		return nil
	}
	value := e.stack[len(e.stack)-1]
	e.stack = e.stack[:len(e.stack)-1]
	return value
}

// ExitNumberLiteral 访问各种语法节点
func (e *ExpressionEvaluator) ExitNumberLiteral(ctx *parser.NumberLiteralContext) {
	text := ctx.GetText()
	if strings.Contains(text, ".") {
		if value, err := strconv.ParseFloat(text, 64); err == nil {
			e.push(value)
		}
	} else {
		if value, err := strconv.ParseInt(text, 10, 64); err == nil {
			e.push(float64(value))
		}
	}
}

func (e *ExpressionEvaluator) ExitStringLiteral(ctx *parser.StringLiteralContext) {
	text := ctx.GetText()
	// 去除引号并处理转义字符
	if len(text) >= 2 {
		text = text[1 : len(text)-1]
		text = strings.ReplaceAll(text, `\\`, `\`)
		text = strings.ReplaceAll(text, `\"`, `"`)
		text = strings.ReplaceAll(text, `\/`, `/`)
		text = strings.ReplaceAll(text, `\b`, "\b")
		text = strings.ReplaceAll(text, `\f`, "\f")
		text = strings.ReplaceAll(text, `\n`, "\n")
		text = strings.ReplaceAll(text, `\r`, "\r")
		text = strings.ReplaceAll(text, `\t`, "\t")
	}
	e.push(text)
}

func (e *ExpressionEvaluator) ExitBooleanLiteral(ctx *parser.BooleanLiteralContext) {
	text := ctx.GetText()
	e.push(text == "true")
}

func (e *ExpressionEvaluator) ExitVariable(ctx *parser.VariableContext) {
	name := ctx.GetText()
	if e.variables != nil {
		value := e.variables.ResolveVariable(name)
		e.push(value)
	} else {
		e.push(nil)
	}
}

func (e *ExpressionEvaluator) ExitFunctionCall(ctx *parser.FunctionCallContext) {
	funcName := ctx.IDENTIFIER().GetText()

	// 收集参数
	argCount := 0
	if ctx.ArgumentList() != nil {
		argCount = len(ctx.ArgumentList().(*parser.ArgumentListContext).AllExpression())
	}

	args := make([]interface{}, argCount)
	for i := argCount - 1; i >= 0; i-- {
		args[i] = e.pop()
	}

	result := e.functions.CallFunction(funcName, args)
	e.push(result)
}

func (e *ExpressionEvaluator) ExitAdditiveExpression(ctx *parser.AdditiveExpressionContext) {
	if len(ctx.AllMultiplicativeExpression()) > 1 {
		right := e.pop()
		left := e.pop()
		log.Println(right, left)
		for i, tree := range ctx.GetChildren() {
			log.Println(i, tree.(antlr.ParseTree).GetText())
		}
		switch ctx.GetChild(1).(antlr.ParseTree).GetText() {
		case "+":
			e.push(e.add(left, right))
		case "-":
			e.push(e.subtract(left, right))
		}
	}
}

func (e *ExpressionEvaluator) ExitMultiplicativeExpression(ctx *parser.MultiplicativeExpressionContext) {
	if len(ctx.AllUnaryExpression()) > 1 {
		right := e.pop()
		left := e.pop()

		switch ctx.GetChild(1).(antlr.ParseTree).GetText() {
		case "*":
			e.push(e.multiply(left, right))
		case "/":
			e.push(e.divide(left, right))
		case "%":
			e.push(e.modulo(left, right))
		}
	}
}

func (e *ExpressionEvaluator) ExitComparisonExpression(ctx *parser.ComparisonExpressionContext) {
	if len(ctx.AllAdditiveExpression()) > 1 {
		right := e.pop()
		left := e.pop()

		switch ctx.GetChild(1).(antlr.ParseTree).GetText() {
		case ">":
			e.push(e.greaterThan(left, right))
		case ">=":
			e.push(e.greaterThanOrEqual(left, right))
		case "<":
			e.push(e.lessThan(left, right))
		case "<=":
			e.push(e.lessThanOrEqual(left, right))
		}
	}
}

func (e *ExpressionEvaluator) ExitEqualityExpression(ctx *parser.EqualityExpressionContext) {
	if len(ctx.AllComparisonExpression()) > 1 {
		right := e.pop()
		left := e.pop()

		switch ctx.GetChild(1).(antlr.ParseTree).GetText() {
		case "==":
			e.push(e.equal(left, right))
		case "!=":
			e.push(e.notEqual(left, right))
		}
	}
}

func (e *ExpressionEvaluator) ExitLogicalAndExpression(ctx *parser.LogicalAndExpressionContext) {
	if len(ctx.AllEqualityExpression()) > 1 {
		right := e.pop()
		left := e.pop()
		e.push(e.logicalAnd(left, right))
	}
}

func (e *ExpressionEvaluator) ExitLogicalOrExpression(ctx *parser.LogicalOrExpressionContext) {
	if len(ctx.AllLogicalAndExpression()) > 1 {
		right := e.pop()
		left := e.pop()
		e.push(e.logicalOr(left, right))
	}
}

func (e *ExpressionEvaluator) ExitUnaryExpression(ctx *parser.UnaryExpressionContext) {
	if ctx.NOT() != nil {
		value := e.pop()
		e.push(e.logicalNot(value))
	} else if ctx.MINUS() != nil {
		value := e.pop()
		e.push(e.negate(value))
	}
}

// 运算方法
func (e *ExpressionEvaluator) add(left, right interface{}) interface{} {
	switch l := left.(type) {
	case float64:
		switch r := right.(type) {
		case float64:
			return l + r
		case string:
			return fmt.Sprintf("%g%s", l, r)
		}
	case string:
		switch r := right.(type) {
		case float64:
			return fmt.Sprintf("%s%g", l, r)
		case string:
			return l + r
		}
	}
	return nil
}

func (e *ExpressionEvaluator) subtract(left, right interface{}) interface{} {
	if l, ok := left.(float64); ok {
		if r, ok := right.(float64); ok {
			return l - r
		}
	}
	return nil
}

func (e *ExpressionEvaluator) multiply(left, right interface{}) interface{} {
	if l, ok := left.(float64); ok {
		if r, ok := right.(float64); ok {
			return l * r
		}
	}
	return nil
}

func (e *ExpressionEvaluator) divide(left, right interface{}) interface{} {
	if l, ok := left.(float64); ok {
		if r, ok := right.(float64); ok {
			if r == 0 {
				return math.NaN()
			}
			return l / r
		}
	}
	return nil
}

func (e *ExpressionEvaluator) modulo(left, right interface{}) interface{} {
	if l, ok := left.(float64); ok {
		if r, ok := right.(float64); ok {
			if r == 0 {
				return math.NaN()
			}
			return float64(int64(l) % int64(r))
		}
	}
	return nil
}

func (e *ExpressionEvaluator) greaterThan(left, right interface{}) bool {
	return e.compare(left, right) > 0
}

func (e *ExpressionEvaluator) greaterThanOrEqual(left, right interface{}) bool {
	return e.compare(left, right) >= 0
}

func (e *ExpressionEvaluator) lessThan(left, right interface{}) bool {
	return e.compare(left, right) < 0
}

func (e *ExpressionEvaluator) lessThanOrEqual(left, right interface{}) bool {
	return e.compare(left, right) <= 0
}

func (e *ExpressionEvaluator) equal(left, right interface{}) bool {
	return e.compare(left, right) == 0
}

func (e *ExpressionEvaluator) notEqual(left, right interface{}) bool {
	return e.compare(left, right) != 0
}

func (e *ExpressionEvaluator) compare(left, right interface{}) int {
	switch l := left.(type) {
	case float64:
		switch r := right.(type) {
		case float64:
			if l < r {
				return -1
			} else if l > r {
				return 1
			}
			return 0
		}
	case string:
		if r, ok := right.(string); ok {
			return strings.Compare(l, r)
		}
	case bool:
		if r, ok := right.(bool); ok {
			if l == r {
				return 0
			} else if !l && r {
				return -1
			}
			return 1
		}
	}
	return -2 // 不可比较
}

func (e *ExpressionEvaluator) logicalAnd(left, right interface{}) bool {
	return e.toBool(left) && e.toBool(right)
}

func (e *ExpressionEvaluator) logicalOr(left, right interface{}) bool {
	return e.toBool(left) || e.toBool(right)
}

func (e *ExpressionEvaluator) logicalNot(value interface{}) bool {
	return !e.toBool(value)
}

func (e *ExpressionEvaluator) negate(value interface{}) interface{} {
	if v, ok := value.(float64); ok {
		return -v
	}
	return nil
}

func (e *ExpressionEvaluator) toBool(value interface{}) bool {
	switch v := value.(type) {
	case bool:
		return v
	case float64:
		return v != 0
	case string:
		return v != ""
	default:
		return value != nil
	}
}

// 获取求值结果
func (e *ExpressionEvaluator) Result() interface{} {
	if len(e.stack) == 0 {
		return nil
	}
	return e.stack[0]
}

// EvaluateExpression 解析并求值表达式
func EvaluateExpression(expr string, variables VariableResolver, functions FunctionCaller) (interface{}, error) {
	input := antlr.NewInputStream(expr)
	lexer := parser.NewConditionExprLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewConditionExprParser(stream)

	// 添加编译错误监听器
	errorListener := &ErrorListener{}
	p.RemoveErrorListeners()
	p.AddErrorListener(errorListener)
	// tree 可缓存，避免重复解析
	tree := p.Expression()

	if errorListener.HasErrors() {
		return nil, fmt.Errorf("parse error: %s", errorListener.ErrorString())
	}

	evaluator := NewExpressionEvaluator(variables, functions)
	antlr.ParseTreeWalkerDefault.Walk(evaluator, tree)

	return evaluator.Result(), nil
}

// ErrorListener 错误监听器
type ErrorListener struct {
	*antlr.DefaultErrorListener
	errors []string
}

func (l *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{},
	line, column int, msg string, e antlr.RecognitionException) {
	l.errors = append(l.errors, fmt.Sprintf("line %d:%d %s", line, column, msg))
}

func (l *ErrorListener) HasErrors() bool {
	return len(l.errors) > 0
}

func (l *ErrorListener) ErrorString() string {
	return strings.Join(l.errors, "; ")
}
