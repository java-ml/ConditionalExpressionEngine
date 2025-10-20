# 基于Antlr4的条件表达式引擎
## 功能
- 支持自定义函数
- 支持算数表达式
- 支持=、!=、>、>=、<、<=、!、||、&&条件运算符

### Java中使用, App.registerDefaultFunctions()默认函数
```
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

```

### Go中使用，func (d *DefaultFunctionCaller) CallFunction(name string, args []interface{})实现默认方法

```
// main.go
package main

import (
	"fmt"
	"log"

	"ConditionExpr/evaluator"
)

// MyVariableResolver 自定义变量解析器
type MyVariableResolver struct {
	vars map[string]interface{}
}

func (m *MyVariableResolver) ResolveVariable(name string) interface{} {
	if value, exists := m.vars[name]; exists {
		return value
	}
	return nil
}

func main() {
	// 设置变量
	resolver := &MyVariableResolver{
		vars: map[string]interface{}{
			"age":       25.0,
			"name":      "John",
			"score":     85.5,
			"isStudent": true,
			"items":     []string{"apple", "banana", "orange"},
		},
	}

	// 测试表达式
	expressions := []string{
		`age > 18 && age < 30`,
		`name == "John" && !isStudent`,
		`score >= 60 && score <= 100`,
		`(age > 20 || score > 90) && isStudent`,
		`contains("hello ", "world")`,
		`len(name) > 3`,
		`age + 5 * 2`,
		`score / 10 + 5`,
		`!("hello" == "")`,
		`"test" + "string"`,
		`9 * -30`,
		`contains(name, "oh") && len(name) == 4`,
		`"line1\nline2" == "line1\nline2"`, // 测试转义字符
	}

	for _, expr := range expressions {
		result, err := evaluator.EvaluateExpression(expr, resolver, nil)
		if err != nil {
			log.Printf("Error evaluating '%s': %v", expr, err)
			continue
		}
		fmt.Printf("'%s' => %v\n", expr, result)
	}
}



```
