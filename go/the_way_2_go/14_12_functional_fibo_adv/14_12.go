/* Using proper closure to make Fibonacci series generator */
package main

import (
    "fmt"
)

type Any interface{}
type EvalFunc func(Any, Any) (Any)
type GeneratorFunc func() Any
type GeneratorUint64Func func() uint64

func main() {

    // lower order function signature must be consistent with generator
    // requirements but must return concrete values to be useful (uint64 in
    // this case to delay (silent)overflow)
    fiboFunc := func(prev, prevprev Any) (Any) {           // general signature
        return prevprev.(uint64) + prev.(uint64)     // specific implementation
    }

    fibo := BuildLazyUint64Evaluator(fiboFunc, uint64(0), uint64(1))

    for i := 1; i < 10; i++ {
        fmt.Printf("%vth fibo number: %v.\n", i, fibo())
    }
}

// general evaluator(not for explicit use in main() function)
func BuildLazyEvaluator(evalFunc EvalFunc, first Any, second Any) GeneratorFunc {
    chGenerator := make(chan Any)                            // general channel

    // channel stuffing loop in the form of dedicated goroutine
    go func() {
        var retVal Any

        for {
            retVal = evalFunc(first, second)
            first, second = second, retVal
            chGenerator <-retVal                                      // blocks
        }
    }()

    // instead of returning of an unbufferd channel handler, return (pointer to?)
    // dedicated function which hides said channel inside → no need for the caller
    // to declare channel outside of this context, just use returned value
    // as an generator
    funcGenerator := func() Any {
        return <-chGenerator
    }

    return funcGenerator
}

// specific evaluator(for use in main function)
func BuildLazyUint64Evaluator(evalFunc EvalFunc, first Any, second Any) GeneratorUint64Func {
    generalEv := BuildLazyEvaluator(evalFunc, first, second)
    return func() uint64 {
        return generalEv().(uint64)
    }
}
