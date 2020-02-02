package basic

import (
	"fmt"
	"io/ioutil"
	"math"
	"math/cmplx"
	"reflect"
	"runtime"
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableTypeDeduc() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variavleShorter() {
	a, b, c, s := 3, 4, false, "test"
	b = 5
	fmt.Println(a, b, c, s)
}

func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))

	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
}

func consts() {
	const filename = "abc"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)

}
func enums() {
	//iota means it is a self-add condition
	const (
		cpp = iota
		_
		java
		python
		golang
	)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, java, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func variablesFunc() {
	fmt.Println("Hello World!\n")
	variableZeroValue()
	variableTypeDeduc()
	variavleShorter()

	euler()
	consts()
	enums()
}

func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("unsupported operator:" + op)
	}
}

func readtext() {
	const filename = "abc.txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

func div(a, b int) (int, int) {
	return a / b, a % b
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args"+
		"(%d,%d)", opName, a, b)
	return op(a, b)
}

func main() {
	//readtext()
	fmt.Println(eval(777, 4396, "-"))
	fmt.Println(eval(3, 4, "+"))
	q, r := div(10, 3)
	fmt.Println(q, r)
	fmt.Println(apply(func(a, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))
}
