// Code generated from ConditionExpr.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // ConditionExpr
import "github.com/antlr4-go/antlr/v4"

// BaseConditionExprListener is a complete listener for a parse tree produced by ConditionExprParser.
type BaseConditionExprListener struct{}

var _ ConditionExprListener = &BaseConditionExprListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseConditionExprListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseConditionExprListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseConditionExprListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseConditionExprListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseConditionExprListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseConditionExprListener) ExitExpression(ctx *ExpressionContext) {}

// EnterLogicalOrExpression is called when production logicalOrExpression is entered.
func (s *BaseConditionExprListener) EnterLogicalOrExpression(ctx *LogicalOrExpressionContext) {}

// ExitLogicalOrExpression is called when production logicalOrExpression is exited.
func (s *BaseConditionExprListener) ExitLogicalOrExpression(ctx *LogicalOrExpressionContext) {}

// EnterLogicalAndExpression is called when production logicalAndExpression is entered.
func (s *BaseConditionExprListener) EnterLogicalAndExpression(ctx *LogicalAndExpressionContext) {}

// ExitLogicalAndExpression is called when production logicalAndExpression is exited.
func (s *BaseConditionExprListener) ExitLogicalAndExpression(ctx *LogicalAndExpressionContext) {}

// EnterEqualityExpression is called when production equalityExpression is entered.
func (s *BaseConditionExprListener) EnterEqualityExpression(ctx *EqualityExpressionContext) {}

// ExitEqualityExpression is called when production equalityExpression is exited.
func (s *BaseConditionExprListener) ExitEqualityExpression(ctx *EqualityExpressionContext) {}

// EnterComparisonExpression is called when production comparisonExpression is entered.
func (s *BaseConditionExprListener) EnterComparisonExpression(ctx *ComparisonExpressionContext) {}

// ExitComparisonExpression is called when production comparisonExpression is exited.
func (s *BaseConditionExprListener) ExitComparisonExpression(ctx *ComparisonExpressionContext) {}

// EnterAdditiveExpression is called when production additiveExpression is entered.
func (s *BaseConditionExprListener) EnterAdditiveExpression(ctx *AdditiveExpressionContext) {}

// ExitAdditiveExpression is called when production additiveExpression is exited.
func (s *BaseConditionExprListener) ExitAdditiveExpression(ctx *AdditiveExpressionContext) {}

// EnterMultiplicativeExpression is called when production multiplicativeExpression is entered.
func (s *BaseConditionExprListener) EnterMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {
}

// ExitMultiplicativeExpression is called when production multiplicativeExpression is exited.
func (s *BaseConditionExprListener) ExitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {
}

// EnterUnaryExpression is called when production unaryExpression is entered.
func (s *BaseConditionExprListener) EnterUnaryExpression(ctx *UnaryExpressionContext) {}

// ExitUnaryExpression is called when production unaryExpression is exited.
func (s *BaseConditionExprListener) ExitUnaryExpression(ctx *UnaryExpressionContext) {}

// EnterPrimaryExpression is called when production primaryExpression is entered.
func (s *BaseConditionExprListener) EnterPrimaryExpression(ctx *PrimaryExpressionContext) {}

// ExitPrimaryExpression is called when production primaryExpression is exited.
func (s *BaseConditionExprListener) ExitPrimaryExpression(ctx *PrimaryExpressionContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseConditionExprListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseConditionExprListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterArgumentList is called when production argumentList is entered.
func (s *BaseConditionExprListener) EnterArgumentList(ctx *ArgumentListContext) {}

// ExitArgumentList is called when production argumentList is exited.
func (s *BaseConditionExprListener) ExitArgumentList(ctx *ArgumentListContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseConditionExprListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseConditionExprListener) ExitVariable(ctx *VariableContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseConditionExprListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseConditionExprListener) ExitLiteral(ctx *LiteralContext) {}

// EnterNumberLiteral is called when production numberLiteral is entered.
func (s *BaseConditionExprListener) EnterNumberLiteral(ctx *NumberLiteralContext) {}

// ExitNumberLiteral is called when production numberLiteral is exited.
func (s *BaseConditionExprListener) ExitNumberLiteral(ctx *NumberLiteralContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseConditionExprListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseConditionExprListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *BaseConditionExprListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *BaseConditionExprListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}
