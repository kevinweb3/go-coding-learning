package basictypes

/**
* 基本类型介绍
* 整型：int, uint，int8, uint8，int16, uint16，int32, uint32，int64, uint64
* 布尔型：bool
* 字节：byte
* 浮点型：float32，float64
* 复数：complex64，complex128
* 字符串：string，s1 := "hello"（字符串转义符，常见转义符包含回车、换行、单双引号、制表符等）
* rune
*/

import (
	"fmt"
	"strings"
	"math"
)

func Types() {

	//字符串
	s1 := "hello!125"
	s2 := "你好"
	a := len(s1) // 字符串长度
	b := s1 + "world" // 字符串拼接
	c := fmt.Sprintf(s2, "world") // 字符串拼接
	d := strings.Split(s1, "!")  // 分割
	f := strings.Contains(s1, "6");  // 判断是否包含
	g := strings.HasPrefix(s1, "h");  // 判断前缀
	h := strings.HasSuffix(s1, "5");  // 判断后缀
	j := strings.Join([]string{"tom", "mike", "lucy"}, "-");  // join操作
	fmt.Println(s1, s2, a, b, c, d, f, g, h, j)
}

// 遍历字符串
func ForString() {
	s := "web3 go"

	// 第一种遍历方式
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v(%c) ", s[i], s[i]);
	}
	fmt.Println()

	// 第二种遍历方式
	for _, v := range s {
		fmt.Printf("%v(%c) ", v, v);
	}
	fmt.Println()
}

// 修改字符串
func ChangeString() {
	s1 := "hello world"
	// 强制类型转换
	byteS1 := []byte(s1);
	byteS1[0] = 'H';
	fmt.Println(string(byteS1));

	s2 := "Go best";
	runeS2 := []rune(s2);
	runeS2[0] = 'W';
	fmt.Println(string(runeS2));
}

// Demo
func SqrtDemo() {
	var a, b = 3, 4;
	var c int;
	c = int(math.Sqrt(float64(a*a + b*b)));
	fmt.Println(c)
}