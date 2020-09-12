package package_test

import "fmt"

// PrintVarInThisPackage is a function
func PrintVarInThisPackage() {

	// 由于处于同一个package_test包内，虽然是在不同的文件中定义的，但此处可以直接使用其他文件中定义的变量。
	// 但是如果被应用到其他包中，只有V1对外可见，x2不可见。因为V1的采用大写之母开头来命名。
	fmt.Println("V1: ", V1)
	fmt.Println("v2: ", x2)
}
