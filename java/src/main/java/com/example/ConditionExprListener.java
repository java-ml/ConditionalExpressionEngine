// Generated from ConditionExpr.g4 by ANTLR 4.13.0
package com.example;
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link ConditionExprParser}.
 */
public interface ConditionExprListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link ConditionExprParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpression(ConditionExprParser.ExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link ConditionExprParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpression(ConditionExprParser.ExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link ConditionExprParser#logicalOrExpression}.
	 * @param ctx the parse tree
	 */
	void enterLogicalOrExpression(ConditionExprParser.LogicalOrExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link ConditionExprParser#logicalOrExpression}.
	 * @param ctx the parse tree
	 */
	void exitLogicalOrExpression(ConditionExprParser.LogicalOrExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link ConditionExprParser#logicalAndExpression}.
	 * @param ctx the parse tree
	 */
	void enterLogicalAndExpression(ConditionExprParser.LogicalAndExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link ConditionExprParser#logicalAndExpression}.
	 * @param ctx the parse tree
	 */
	void exitLogicalAndExpression(ConditionExprParser.LogicalAndExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link ConditionExprParser#equalityExpression}.
	 * @param ctx the parse tree
	 */
	void enterEqualityExpression(ConditionExprParser.EqualityExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link ConditionExprParser#equalityExpression}.
	 * @param ctx the parse tree
	 */
	void exitEqualityExpression(ConditionExprParser.EqualityExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link ConditionExprParser#relationalExpression}.
	 * @param ctx the parse tree
	 */
	void enterRelationalExpression(ConditionExprParser.RelationalExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link ConditionExprParser#relationalExpression}.
	 * @param ctx the parse tree
	 */
	void exitRelationalExpression(ConditionExprParser.RelationalExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link ConditionExprParser#additiveExpression}.
	 * @param ctx the parse tree
	 */
	void enterAdditiveExpression(ConditionExprParser.AdditiveExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link ConditionExprParser#additiveExpression}.
	 * @param ctx the parse tree
	 */
	void exitAdditiveExpression(ConditionExprParser.AdditiveExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link ConditionExprParser#multiplicativeExpression}.
	 * @param ctx the parse tree
	 */
	void enterMultiplicativeExpression(ConditionExprParser.MultiplicativeExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link ConditionExprParser#multiplicativeExpression}.
	 * @param ctx the parse tree
	 */
	void exitMultiplicativeExpression(ConditionExprParser.MultiplicativeExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link ConditionExprParser#unaryExpression}.
	 * @param ctx the parse tree
	 */
	void enterUnaryExpression(ConditionExprParser.UnaryExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link ConditionExprParser#unaryExpression}.
	 * @param ctx the parse tree
	 */
	void exitUnaryExpression(ConditionExprParser.UnaryExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link ConditionExprParser#primaryExpression}.
	 * @param ctx the parse tree
	 */
	void enterPrimaryExpression(ConditionExprParser.PrimaryExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link ConditionExprParser#primaryExpression}.
	 * @param ctx the parse tree
	 */
	void exitPrimaryExpression(ConditionExprParser.PrimaryExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link ConditionExprParser#functionCall}.
	 * @param ctx the parse tree
	 */
	void enterFunctionCall(ConditionExprParser.FunctionCallContext ctx);
	/**
	 * Exit a parse tree produced by {@link ConditionExprParser#functionCall}.
	 * @param ctx the parse tree
	 */
	void exitFunctionCall(ConditionExprParser.FunctionCallContext ctx);
	/**
	 * Enter a parse tree produced by {@link ConditionExprParser#argumentList}.
	 * @param ctx the parse tree
	 */
	void enterArgumentList(ConditionExprParser.ArgumentListContext ctx);
	/**
	 * Exit a parse tree produced by {@link ConditionExprParser#argumentList}.
	 * @param ctx the parse tree
	 */
	void exitArgumentList(ConditionExprParser.ArgumentListContext ctx);
	/**
	 * Enter a parse tree produced by {@link ConditionExprParser#variable}.
	 * @param ctx the parse tree
	 */
	void enterVariable(ConditionExprParser.VariableContext ctx);
	/**
	 * Exit a parse tree produced by {@link ConditionExprParser#variable}.
	 * @param ctx the parse tree
	 */
	void exitVariable(ConditionExprParser.VariableContext ctx);
	/**
	 * Enter a parse tree produced by {@link ConditionExprParser#number}.
	 * @param ctx the parse tree
	 */
	void enterNumber(ConditionExprParser.NumberContext ctx);
	/**
	 * Exit a parse tree produced by {@link ConditionExprParser#number}.
	 * @param ctx the parse tree
	 */
	void exitNumber(ConditionExprParser.NumberContext ctx);
}