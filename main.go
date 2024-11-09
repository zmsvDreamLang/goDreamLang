package main

import (
	"fmt"
	"os"
	"time"

	"dreamlang/parser"

	"github.com/sanity-io/litter"
)

// main 函数是程序的入口点。它执行以下操作：
// 1. 读取名为 "test.lang" 的文件内容，并将其转换为字符串。
// 2. 记录解析操作的开始时间。
// 3. 使用 parser.Parse 函数解析源代码字符串，生成抽象语法树（AST）。
// 4. 计算解析操作所花费的时间。
// 5. 使用 litter.Dump 函数输出生成的 AST。
// 6. 打印解析操作所花费的时间。
func main() {
	sourceBytes, _ := os.ReadFile("test.lang")
	source := string(sourceBytes)
	start := time.Now()
	ast := parser.Parse(source)
	duration := time.Since(start)

	litter.Dump(ast)
	fmt.Printf("Duration: %v\n", duration)
}
