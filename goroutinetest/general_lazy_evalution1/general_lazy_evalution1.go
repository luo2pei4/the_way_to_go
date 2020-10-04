package main

import "fmt"

// Any is an interface
type Any interface{}

// EvalFunc is a type of func
type EvalFunc func(Any) (Any, Any)

func main() {

	// 实际执行的lambda函数
	evenFunc := func(state Any) (Any, Any) {

		fmt.Printf("Call evenFunc, state address is %p\n", &state)
		os := state.(int)
		ns := os + 2
		return os, ns
	}

	even := BuildLazyIntEvaluator(evenFunc, 0)

	for i := 0; i < 10; i++ {

		fmt.Printf("%vth even: %v\n", i, even())
	}
}

// BuildLazyEvaluator 共通的惰性生成器
// evalFunc是传入的函数，initState则是传入生成器的初始值（可以为任意值）
// 在惰性生成器中创建了一个通道retValChan
// 两个匿名函数loopFunc和retFunc，loopFunc函数在协程中调用，retFunc函数则用于返回
// loopFunc函数被创建为协程。在一个无限循环中获取传入函数的计算结果，并将其中一个计算结果放入retValChan通道内
// retFunc是一个从retValChan通道获取数据的函数。如果该函数不被调用，retValChan通道将被阻塞。
//
func BuildLazyEvaluator(evalFunc EvalFunc, initState Any) func() Any {

	retValChan := make(chan Any)

	loopFunc := func() {

		fmt.Println("Call loopFunc")
		var actState Any = initState
		var retVal Any

		for {

			fmt.Printf("In loopFunc, actState address is %p\n", &actState)
			retVal, actState = evalFunc(actState)
			retValChan <- retVal
		}
	}

	retFunc := func() Any {
		return <-retValChan
	}

	go loopFunc()

	return retFunc
}

// BuildLazyIntEvaluator is a function
func BuildLazyIntEvaluator(evalFunc EvalFunc, initState Any) func() int {

	ef := BuildLazyEvaluator(evalFunc, initState)

	return func() int {

		return ef().(int)
	}
}
