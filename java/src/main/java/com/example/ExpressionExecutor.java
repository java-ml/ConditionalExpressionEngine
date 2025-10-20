package com.example;

import java.util.*;

public class ExpressionExecutor extends ConditionExprBaseVisitor<Object> {
    
    private Map<String, Object> variables = new HashMap<>();
    private Map<String, Function> functions = new HashMap<>();
    
    public interface Function {
        Object execute(List<Object> args);
    }
    
    public ExpressionExecutor() {
        registerDefaultFunctions();
    }
    
    public void setVariable(String name, Object value) {
        variables.put(name, value);
    }
    
    public void registerFunction(String name, Function function) {
        functions.put(name, function);
    }
    
    private void registerDefaultFunctions() {
        registerFunction("abs", args -> {
            validateArgs("abs", args, 1);
            Number num = toNumber(args.get(0));
            return num instanceof Integer ? 
                Math.abs(num.intValue()) : Math.abs(num.doubleValue());
        });
        
        registerFunction("max", args -> {
            validateArgs("max", args, 2);
            Number n1 = toNumber(args.get(0));
            Number n2 = toNumber(args.get(1));
            return Math.max(n1.doubleValue(), n2.doubleValue());
        });
        
        registerFunction("min", args -> {
            validateArgs("min", args, 2);
            Number n1 = toNumber(args.get(0));
            Number n2 = toNumber(args.get(1));
            return Math.min(n1.doubleValue(), n2.doubleValue());
        });
        
        registerFunction("length", args -> {
            System.out.println("length--- args: " + args);
            validateArgs("length", args, 1);
            return args.get(0).toString().length();
        });
        
        registerFunction("upper", args -> {
            validateArgs("upper", args, 1);
            return args.get(0).toString().toUpperCase();
        });
    }
    
    private void validateArgs(String funcName, List<Object> args, int expected) {
        if (args.size() != expected) {
            throw new RuntimeException("Function " + funcName + " expects " + 
                expected + " arguments, got " + args.size());
        }
    }
    
    // 访问器方法实现
    @Override
    public Object visitExpression(ConditionExprParser.ExpressionContext ctx) {
        return visit(ctx.logicalOrExpression());
    }


    
    @Override
    public Object visitLogicalOrExpression(ConditionExprParser.LogicalOrExpressionContext ctx) {
        Object result = visit(ctx.logicalAndExpression(0));
        
        for (int i = 1; i < ctx.logicalAndExpression().size(); i++) {
            boolean left = toBoolean(result);
            if (left) return true;
            
            Object right = visit(ctx.logicalAndExpression(i));
            result = left || toBoolean(right);
        }
        
        return result;
    }
    
    @Override
    public Object visitLogicalAndExpression(ConditionExprParser.LogicalAndExpressionContext ctx) {
        Object result = visit(ctx.equalityExpression(0));
        
        for (int i = 1; i < ctx.equalityExpression().size(); i++) {
            boolean left = toBoolean(result);
            if (!left) return false;
            
            Object right = visit(ctx.equalityExpression(i));
            result = left && toBoolean(right);
        }
        
        return result;
    }
    
    @Override
    public Object visitEqualityExpression(ConditionExprParser.EqualityExpressionContext ctx) {
        Object result = visit(ctx.relationalExpression(0));
        
        for (int i = 1; i < ctx.relationalExpression().size(); i++) {
            Object right = visit(ctx.relationalExpression(i));
            
            if (ctx.getChild(2*i-1).getText().equals("==")) {
                result = equals(result, right);
            } else {
                result = !equals(result, right);
            }
        }
        
        return result;
    }
    
    @Override
    public Object visitRelationalExpression(ConditionExprParser.RelationalExpressionContext ctx) {
        Object left = visit(ctx.additiveExpression(0));
        
        for (int i = 1; i < ctx.additiveExpression().size(); i++) {
            Object right = visit(ctx.additiveExpression(i));
            String operator = ctx.getChild(2*i-1).getText();
            
            switch (operator) {
                case ">": left = compare(left, right) > 0; break;
                case ">=": left = compare(left, right) >= 0; break;
                case "<": left = compare(left, right) < 0; break;
                case "<=": left = compare(left, right) <= 0; break;
                default: throw new RuntimeException("Unknown operator: " + operator);
            }
        }
        
        return left;
    }
    
    @Override
    public Object visitAdditiveExpression(ConditionExprParser.AdditiveExpressionContext ctx) {
        Object result = visit(ctx.multiplicativeExpression(0));
        
        for (int i = 1; i < ctx.multiplicativeExpression().size(); i++) {
            Object right = visit(ctx.multiplicativeExpression(i));
            String operator = ctx.getChild(2*i-1).getText();
            
            if (operator.equals("+")) {
                result = add(result, right);
            } else {
                result = subtract(result, right);
            }
        }
        
        return result;
    }
    
    @Override
    public Object visitMultiplicativeExpression(ConditionExprParser.MultiplicativeExpressionContext ctx) {
        Object result = visit(ctx.unaryExpression(0));
        
        for (int i = 1; i < ctx.unaryExpression().size(); i++) {
            Object right = visit(ctx.unaryExpression(i));
            String operator = ctx.getChild(2*i-1).getText();
            
            switch (operator) {
                case "*": result = multiply(result, right); break;
                case "/": result = divide(result, right); break;
                case "%": result = modulo(result, right); break;
                default: throw new RuntimeException("Unknown operator: " + operator);
            }
        }
        
        return result;
    }
    
    @Override
    public Object visitUnaryExpression(ConditionExprParser.UnaryExpressionContext ctx) {
        if (ctx.NOT() != null) {
            return !toBoolean(visit(ctx.unaryExpression()));
        } else if (ctx.MINUS() != null) {
            Object value = visit(ctx.unaryExpression());
            if (value instanceof Integer) {
                return -((Integer) value);
            } else if (value instanceof Double) {
                return -((Double) value);
            } else {
                throw new RuntimeException("Unary minus not supported for type: " + value.getClass());
            }
        } else {
            return visit(ctx.primaryExpression());
        }
    }
    
    @Override
    public Object visitPrimaryExpression(ConditionExprParser.PrimaryExpressionContext ctx) {
        if (ctx.expression() != null) {
            return visit(ctx.expression());
        } else if (ctx.functionCall() != null) {
            return visit(ctx.functionCall());
        } else if (ctx.variable() != null) {
            return visit(ctx.variable());
        } else if (ctx.number() != null) {
            return visit(ctx.number());
        } else if (ctx.BOOLEAN() != null) {
            return Boolean.parseBoolean(ctx.BOOLEAN().getText());
        } else if (ctx.STRING() != null) {
            System.out.println("========" +  ctx.STRING().getText());
            return EscapeUtils.unescapeString(ctx.STRING().getText());
        }
        throw new RuntimeException("Unknown primary expression");
    }
    
    @Override
    public Object visitFunctionCall(ConditionExprParser.FunctionCallContext ctx) {
        String funcName = ctx.IDENTIFIER().getText();
        Function function = functions.get(funcName);
        
        if (function == null) {
            throw new RuntimeException("Unknown function: " + funcName);
        }
        
        List<Object> args = new ArrayList<>();
        if (ctx.argumentList() != null) {
            for (ConditionExprParser.ExpressionContext argCtx : ctx.argumentList().expression()) {
                args.add(visit(argCtx));
            }
        }
        
        return function.execute(args);
    }
    
    @Override
    public Object visitVariable(ConditionExprParser.VariableContext ctx) {
        String varName = ctx.IDENTIFIER().getText();
        Object value = variables.get(varName);
        
        if (value == null) {
            throw new RuntimeException("Undefined variable: " + varName);
        }
        
        return value;
    }
    
    @Override
    public Object visitNumber(ConditionExprParser.NumberContext ctx) {
        if (ctx.INTEGER() != null) {
            return Integer.parseInt(ctx.INTEGER().getText());
        } else if (ctx.FLOAT() != null) {
            return Double.parseDouble(ctx.FLOAT().getText());
        }
        throw new RuntimeException("Unknown number format");
    }
    
    // 工具方法保持不变
    private boolean toBoolean(Object obj) {
        if (obj instanceof Boolean) return (Boolean) obj;
        if (obj instanceof Number) return ((Number) obj).doubleValue() != 0;
        if (obj instanceof String) return !((String) obj).isEmpty();
        return obj != null;
    }
    
    private Number toNumber(Object obj) {
        if (obj instanceof Number) return (Number) obj;
        if (obj instanceof String) {
            try {
                return Double.parseDouble((String) obj);
            } catch (NumberFormatException e) {
                throw new RuntimeException("Cannot convert to number: " + obj);
            }
        }
        throw new RuntimeException("Unsupported type for number conversion: " + obj.getClass());
    }
    
    private boolean equals(Object left, Object right) {
        if (left == null && right == null) return true;
        if (left == null || right == null) return false;
        return left.equals(right);
    }
    
    private int compare(Object left, Object right) {
        if (left instanceof Number && right instanceof Number) {
            double d1 = ((Number) left).doubleValue();
            double d2 = ((Number) right).doubleValue();
            return Double.compare(d1, d2);
        }
        if (left instanceof Comparable && right instanceof Comparable) {
            return ((Comparable) left).compareTo(right);
        }
        throw new RuntimeException("Cannot compare: " + left + " and " + right);
    }
    
    private Object add(Object left, Object right) {
        if (left instanceof Number && right instanceof Number) {
            Number n1 = (Number) left;
            Number n2 = (Number) right;
            
            if (n1 instanceof Integer && n2 instanceof Integer) {
                return n1.intValue() + n2.intValue();
            } else {
                return n1.doubleValue() + n2.doubleValue();
            }
        }
        if (left instanceof String || right instanceof String) {
            return left.toString() + right.toString();
        }
        throw new RuntimeException("Cannot add: " + left + " and " + right);
    }
    
    private Object subtract(Object left, Object right) {
        Number n1 = toNumber(left);
        Number n2 = toNumber(right);
        
        if (n1 instanceof Integer && n2 instanceof Integer) {
            return n1.intValue() - n2.intValue();
        } else {
            return n1.doubleValue() - n2.doubleValue();
        }
    }
    
    private Object multiply(Object left, Object right) {
        Number n1 = toNumber(left);
        Number n2 = toNumber(right);
        
        if (n1 instanceof Integer && n2 instanceof Integer) {
            return n1.intValue() * n2.intValue();
        } else {
            return n1.doubleValue() * n2.doubleValue();
        }
    }
    
    private Object divide(Object left, Object right) {
        Number n1 = toNumber(left);
        Number n2 = toNumber(right);
        
        if (n2.doubleValue() == 0) {
            throw new RuntimeException("Division by zero");
        }
        
        if (n1 instanceof Integer && n2 instanceof Integer) {
            return n1.intValue() / n2.intValue();
        } else {
            return n1.doubleValue() / n2.doubleValue();
        }
    }
    
    private Object modulo(Object left, Object right) {
        Number n1 = toNumber(left);
        Number n2 = toNumber(right);
        
        if (n1 instanceof Integer && n2 instanceof Integer) {
            return n1.intValue() % n2.intValue();
        } else {
            return n1.doubleValue() % n2.doubleValue();
        }
    }
}
