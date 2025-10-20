// Code generated from ConditionExpr.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // ConditionExpr
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by ConditionExprParser.
type ConditionExprVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ConditionExprParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by ConditionExprParser#logicalOrExpression.
	VisitLogicalOrExpression(ctx *LogicalOrExpressionContext) interface{}

	// Visit a parse tree produced by ConditionExprParser#logicalAndExpression.
	VisitLogicalAndExpression(ctx *LogicalAndExpressionContext) interface{}

	// Visit a parse tree produced by ConditionExprParser#equalityExpression.
	VisitEqualityExpression(ctx *EqualityExpressionContext) interface{}

	// Visit a parse tree produced by ConditionExprParser#comparisonExpression.
	VisitComparisonExpression(ctx *ComparisonExpressionContext) interface{}

	// Visit a parse tree produced by ConditionExprParser#additiveExpression.
	VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{}

	// Visit a parse tree produced by ConditionExprParser#multiplicativeExpression.
	VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{}

	// Visit a parse tree produced by ConditionExprParser#unaryExpression.
	VisitUnaryExpression(ctx *UnaryExpressionContext) interface{}

	// Visit a parse tree produced by ConditionExprParser#primaryExpression.
	VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{}

	// Visit a parse tree produced by ConditionExprParser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by ConditionExprParser#argumentList.
	VisitArgumentList(ctx *ArgumentListContext) interface{}

	// Visit a parse tree produced by ConditionExprParser#variable.
	VisitVariable(ctx *VariableContext) interface{}

	// Visit a parse tree produced by ConditionExprParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by ConditionExprParser#numberLiteral.
	VisitNumberLiteral(ctx *NumberLiteralContext) interface{}

	// Visit a parse tree produced by ConditionExprParser#stringLiteral.
	VisitStringLiteral(ctx *StringLiteralContext) interface{}

	// Visit a parse tree produced by ConditionExprParser#booleanLiteral.
	VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{}
}
