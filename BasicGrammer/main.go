package main

import (
	"BasicGrammer/array"
	"BasicGrammer/basictypes"
)

func main() {
	// 基本类型
	basictypes.Underline1();  // 下划线
	basictypes.Init();   // 
	basictypes.Variable();  // 变量（局部变量、全局变量）常量定义
	basictypes.Types();   // 
	basictypes.ForString();
	basictypes.ChangeString();
	basictypes.SqrtDemo();

	// 数组
	array.ArrayDefine();
	array.ArrayTraversal();
	array.ArrayCopy();
	array.ArraySum();
	array.FindArrayIndex();
}