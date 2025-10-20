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
