// you can find the tutorial here: https://learnxinyminutes.com/zh-cn/go/
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Hello World!")

	beyondHello()
}

func beyondHello() {
	var x int // variable declaration, you have to declare the variable before use it.
	x = 3     // variable assignment
	y := 4    // short variable declaration, combined the type declaration and value assignment
	sum, prod := learnMultiple(x, y)
	fmt.Println("sum: ", sum, "prod: ", prod)
	learnTypes()
}

func learnMultiple(x, y int) (sum, prod int) {
	return x + y, x * y
}

func learnTypes() {
	// short variable declaration will automatically determine the type
	str := "Speak less and learn more!"

	// raw string literal, use `` to declare it.
	s2 := `this is a string
    that can be wrapped`
	fmt.Println(str, "\n", s2)

	// rune type, alias for int32, use UTF-8 encoding method
	g := 'a'

	f := 3.14 // float64, using IEEE-754

	var u uint = 7 // you can directly assign a value in the declaration
	var pi float32 = 22. / 7

	// turn into byte
	n := byte('\n')

	// array
	var a4 [4]int
	a3 := [...]int{3, 1, 5}

	s3 := []int{3, 1, 5}
	s4 := make([]int, 4)
	s5 := make([]int, 0, 4)
	var d2 [][]float64 // we only declare d2 and no memory space has been allocated
	bs := []byte("a slice")

	s := []int{1, 2, 3}
	s = append(s, 4, 5, 6)
	fmt.Println(s)

	// Compared that didn't return the result to s
	// s := []int{1, 2, 3}
	// append(s, 4, 5, 6)
	// fmt.Println(s)

	// 后面的省略号表示引用一个slice，解包其中的元素
	s = append(s, []int{7, 8, 9}...)
	fmt.Println(s)

	p, q := learnMemory()

	m := map[string]int{"three": 3, "four": 4}
	m["one"] = 1

	// use _ to shut up the compiler complaining the variable is daclared but not used
	_, _, _, _, _, _, _, _, _, _, _, _, _, _ = a4, a3, s3, s4, s5, p, q, bs, d2, g, f, n, u, pi
	// 通常的用法是，在调用拥有多个返回值的函数时，
	// 用下划线抛弃其中的一个参数。下面的例子就是一个脏套路，
	// 调用os.Create并用下划线变量扔掉它的错误代码。
	// 因为我们觉得这个文件一定会成功创建。
	file, _ := os.Create("test.txt")
	fmt.Fprint(file, "文件是这样写入的！")
	file.Close()

	learnFlowControl()
}

func learnMemory() (p, q *int) {
	// the new function will allocate memory for p and the value will be set to zero
	p = new(int)
	s := make([]int, 20)
	s[3] = 7
	r := -2
	// memory escape
	return &s[3], &r
}

func learnFlowControl() {
	// if statement don't need ()
	if true {
		fmt.Println("我踏马来啦")
	}
	x := 1
	switch x {
	case 0:
	case 1:
		// implicitly call break
	case 2:
		// unreachable
	}

	for x := 0; x < 3; x++ {
		fmt.Println("x: ", x)
	}
	// note that the x is still 1
	// you can use range to enumerate the array, slice, map, string, channel
	// for channel, the range return one value
	// the other will return two
	for key, value := range map[string]int{"three": 3, "two": 2, "one": 1} {
		fmt.Printf("key:%s value:%d\n", key, value)
	}

	// in case you only want the value, drop the index using _
	for _, name := range []string{"asan", "xiaoli", "honghong"} {
		fmt.Printf("you are %s\n", name)
	}

	// closure funtion
	expensiveComputation := func() int {
		return 1e6
	}

	// same as for, use := to assign y firstly and then compare it to x
	if y := expensiveComputation(); y > x {
		x = y
	}

	// closure function can capture the out sider variable x
	xBig := func() bool {
		return x > 100
	}
	fmt.Println("xBig:", xBig())
	x /= 1e5
	fmt.Println("xBig:", xBig())

	// 除此之外，函数体可以在其他函数中定义并调用，
	// 满足下列条件时，也可以作为参数传递给其他函数：
	//   a) 定义的函数被立即调用
	//   b) 函数返回值符合调用者对类型的要求
	fmt.Println("两数相加乘二", func(a, b int) int {
		return (a + b) * 2
	})


    // Go also have goto and you will surely love it!
    goto GPoint
GPoint:
    learnFunctionFactory() // we will create some functions that return a function
    learnDefer()
    learnInterface()
}

func learnFunctionFactory() {
    f := sentenceFactory("原谅")
    fmt.Println(f("当然选择", "她"))
}

func sentenceFactory(mystring string) func(before, after string) string {
    return func(before, after string) string {
        return fmt.Sprintf("%s %s %s", before, mystring, after)
    }
}

func learnDefer() bool {
    // defer expression will execute before the moment when this function returns
    defer fmt.Println("defer执行的顺序是后进先出，类比成栈")
    defer fmt.Println("这条语句会比上面的先执行，因为")
    return true
}

type Stringer interface {
    String() string
}

type pair struct {
    x, y int
}

func (p pair) String() string {
    return fmt.Sprintf("(%d.%d)", p.x, p.y)
}

func learnInterface() {
    p := pair{1, 2}
    var i Stringer
    i = p
    // the print function will automatically call the String() for us
    // read the source code
    fmt.Println(p, i)

    learnVariadicParams("How", "are", "you", "these", "days")
}

func learnVariadicParams(myStrings ...any) {
    for _, param := range myStrings {
        fmt.Println("param: ", param)
    }

    fmt.Println("params: ", fmt.Sprintln(myStrings...))
    // fmt.Println("params: ", fmt.Sprintln(myStrings))
    // Without the ... to decode the slice, we can still work well.
    // Why?
    // cause []any is also any!

    learnErrorHandling()
}

func learnErrorHandling() {
    m := map[int]string{3: "three", 4: "four"}
    if x, ok := m[1]; !ok {
        fmt.Println("the key you search don't exist!")
    } else {
        fmt.Println("never want to get here and know what x is", x)
    }

    if _, err := strconv.Atoi("non-int"); err != nil {
        fmt.Println(err)
    }
    learnConcurrency()
}

func inc(i int, c chan<- int) {
    c <- i + 1
}

func learnConcurrency() {
    c := make(chan int)
    go inc(1, c)
    go inc(2, c)
    go inc(3, c)

    fmt.Println(<-c, <-c, <-c)
    cs := make(chan string)
    cc := make(chan chan string)
    go func() {
        c <- 4
    }()
    go func() {
        cs <- "Can you hear me"
    }()
    select {
    case i := <- c:
        fmt.Println("receive:", i)
    case <- cs:
        fmt.Println("我草，我好像被注入了字符串")
    case <- cc:
        fmt.Println("根本不会过来呢")
    }
}
