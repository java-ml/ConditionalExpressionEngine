package com.example;

import java.util.Map;
import java.util.HashMap;


import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.CharStreams;
import org.antlr.v4.runtime.CommonTokenStream;
import org.antlr.v4.runtime.tree.ParseTree;

public class App {

    public static Object execute(String expression, Map<String, Object> variables) {
        try {
            CharStream input = CharStreams.fromString(expression);
            ConditionExprLexer lexer = new ConditionExprLexer(input);
            CommonTokenStream tokens = new CommonTokenStream(lexer);
            ConditionExprParser parser = new ConditionExprParser(tokens);
            //parser.removeErrorListeners(); 编译错误处理
            //parser.addErrorListener(new ConditionExprErrorListener());
            // ParseTree 可缓存复用， 多线程安全
            ParseTree tree = parser.expression();

            ExpressionExecutor executor = new ExpressionExecutor();
            if (variables != null) {
                for (Map.Entry<String, Object> entry : variables.entrySet()) {
                    executor.setVariable(entry.getKey(), entry.getValue());
                }
            }
            
            return executor.visit(tree);

        } catch (Exception e) {
            throw new RuntimeException("Expression evaluation error: " + e.getMessage(), e);
        }
    }

    public static void main(String[] args) {
        Map<String, Object> variables = new HashMap<>();

        // 设置一些示例变量
        variables.put("age", 25);
        variables.put("salary", 50000.0);
        variables.put("name", "Jack");
        variables.put("isEmployed", true);

        try {
            Object result = execute("length(name) > 3 && (salary  - 500 > 40000) && isEmployed && age / 2 != 25", variables);
            System.out.println("结果: " + result + " (类型: " + result.getClass().getSimpleName() + ")");
        } catch (Exception e) {
            System.out.println("错误: " + e.getMessage());
            e.printStackTrace();
        }
    }
}
