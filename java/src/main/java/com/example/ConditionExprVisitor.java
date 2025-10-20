// Generated from ConditionExpr.g4 by ANTLR 4.13.0
package com.example;
import org.antlr.v4.runtime.tree.ParseTreeVisitor;

/**
 * This interface defines a complete generic visitor for a parse tree produced
 * by {@link ConditionExprParser}.
 *
 * @param <T> The return type of the visit operation. Use {@link Void} for
 * operations with no return type.
 */
public interface ConditionExprVisitor<T> extends ParseTreeVisitor<T> {
	/**
	 * Visit a parse tree produced by {@link ConditionExprParser#expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitExpression(ConditionExprParser.ExpressionContext ctx);
	/**
	 * Visit a parse tree produced by {@link ConditionExprParser#logicalOrExpression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitLogicalOrExpression(ConditionExprParser.LogicalOrExpressionContext ctx);
	/**
	 * Visit a parse tree produced by {@link ConditionExprParser#logicalAndExpression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitLogicalAndExpression(ConditionExprParser.LogicalAndExpressionContext ctx);
	/**
	 * Visit a parse tree produced by {@link ConditionExprParser#equalityExpression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitEqualityExpression(ConditionExprParser.EqualityExpressionContext ctx);
	/**
	 * Visit a parse tree produced by {@link ConditionExprParser#relationalExpression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitRelationalExpression(ConditionExprParser.RelationalExpressionContext ctx);
	/**
	 * Visit a parse tree produced by {@link ConditionExprParser#additiveExpression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitAdditiveExpression(ConditionExprParser.AdditiveExpressionContext ctx);
	/**
	 * Visit a parse tree produced by {@link ConditionExprParser#multiplicativeExpression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitMultiplicativeExpression(ConditionExprParser.MultiplicativeExpressionContext ctx);
	/**
	 * Visit a parse tree produced by {@link ConditionExprParser#unaryExpression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitUnaryExpression(ConditionExprParser.UnaryExpressionContext ctx);
	/**
	 * Visit a parse tree produced by {@link ConditionExprParser#primaryExpression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPrimaryExpression(ConditionExprParser.PrimaryExpressionContext ctx);
	/**
	 * Visit a parse tree produced by {@link ConditionExprParser#functionCall}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFunctionCall(ConditionExprParser.FunctionCallContext ctx);
	/**
	 * Visit a parse tree produced by {@link ConditionExprParser#argumentList}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitArgumentList(ConditionExprParser.ArgumentListContext ctx);
	/**
	 * Visit a parse tree produced by {@link ConditionExprParser#variable}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitVariable(ConditionExprParser.VariableContext ctx);
	/**
	 * Visit a parse tree produced by {@link ConditionExprParser#number}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitNumber(ConditionExprParser.NumberContext ctx);
}