package main

import "fmt"

// Any is an interface
type Any interface{}

// EvalFunc is a type of func
type EvalFunc func(Any) (Any, Any)

func main() {

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

// BuildLazyEvaluator is a function
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
