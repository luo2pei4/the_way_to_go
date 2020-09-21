package main

import "fmt"

type funcType func()

func main() {

	test5()
}

func dftest() {

	i := 0
	defer fmt.Println(i)
	i++
}

func test() (ret int) {

	defer func() {
		ret++
		fmt.Println("ret is: ", ret)
	}()

	return 1
}

func test2() (ret int) {

	for i := 0; i < 10; i++ {
		defer fmt.Println(i + ret)
	}

	return 1
}

func test3() (ret int) {

	tmp := 1

	for i := 0; i < 10; i++ {

		defer func() {

			fmt.Println(ret + tmp)
			tmp++
		}()
	}

	return 0
}

func test4() (ret int) {

	tmp := 0
	fmt.Printf("The address of tmp is %v\n", &tmp)
	fmt.Printf("The address of ret is %v\n", &ret)

	for i := 0; i < 5; i++ {

		ptr := func() {

			fmt.Printf("tmp is %d, tmp address is %p, ret is %d, ret address is %p, ret + tmp = %d\n", tmp, &tmp, ret, &ret, ret+tmp)
			tmp++
		}

		fmt.Printf("The address of ptr is %v\n", &ptr)
		fmt.Printf("The address of non-declaration func is %p\n", ptr)
		defer ptr()
	}

	return 1
}

func test5() (ret int) {

	tmp := 0
	fmt.Printf("The address of tmp is %v\n", &tmp)
	fmt.Printf("The address of ret is %v\n", &ret)
	var ptr funcType

	for i := 0; i < 5; i++ {

		ptr = func() {

			fmt.Printf("tmp is %d, tmp address is %p, ret is %d, ret address is %p, ret + tmp = %d\n", tmp, &tmp, ret, &ret, ret+tmp)
			tmp++
		}

		fmt.Printf("The address of ptr is %v\n", &ptr)
		fmt.Printf("The address of non-declaration func is %p\n", ptr)
		defer ptr()
	}

	return 1
}
