package basictypes

/*
* “_”是特殊标识符，用来忽略结果。
*  import 下划线（如：import _ hello/imp）的作用：当导入一个包时，该包下的文件里所有init()函数都会被执行，然而，有些时候我们并不需要把整个包都导入进来，仅仅是是希望它执行init()函数而已。这个时候就可以使用 import _ 引用该包。即使用【import _ 包路径】只是引用该包，仅仅是为了调用init()函数，所以无法通过包名来调用包中的其他函数。
*/

import (
	"fmt"
	"os"
	_ "github.com/go-sql-driver/mysql"  // 不直接使用mysql包，只是执行一下这个包的init函数，把mysql的驱动注册到sql包里，然后程序里就可以使用sql包来访问mysql数据库了。
)

func Underline1() {
	buf := make([]byte, 1024);
	// 下划线意思是忽略这个变量, 普通写法是f,err := os.Open("xxxxxxx"), 如果此时不需要知道返回的错误值, 就可以用f, _ := os.Open("xxxxxx")
	f, _ := os.Open("../main.go");  
	defer f.Close();
	for {
		n, _ := f.Read(buf);
		if n == 0 {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}

func Init() {
	fmt.Println("imp-init() come here.")
}

