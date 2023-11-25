package basictypes

/*
* 变量声明
* var 变量名 变量类型
* 函数外的每个语句都必须以关键字开始（var、const、func等）
* :=不能使用在函数外。
* _多用于占位，表示忽略值。
*/

import (
	"fmt"
)

// 常量声明
const pi = 3.1415
const e = 2.7182

// 全局变量m
var m = 100

func Variable() {
	// 变量声明
	var name string
  var age int
  var isOk bool

	// 变量批量声明
	var (
			a string
			b int
			c bool
			d float32
	)

	// 变量初始化
	var name1 string = "pprof.cn"
  var sex int = 1

	// 匿名变量
	x, _ := foo();

	fmt.Println(pi, e, m, name, age, isOk, a, b, c , d, name1, sex, x);
}

func foo() (int, string) {
	return 10, "Q1mi"
}