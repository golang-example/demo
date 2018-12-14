package main //声明本文件的package名

import "fmt" //import语言的fmt库——---用于输出

func main() {
	fmt.Println("Hello, World!")
}

// go run hello_world.go 实际是编译成二进制文件 再执行

// go build hello_world.go
// ./hello_world

/*
第一行代码 package main 定义了包名。你必须在源文件中非注释的第一行指明这个文件属于哪个包，
如：package main。package main表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包。

下一行 import "fmt" 告诉 Go 编译器这个程序需要使用 fmt 包（的函数，或其他元素），fmt 包实现了格式化 IO（输入/输出）的函数。

下一行 func main() 是程序开始执行的函数。main 函数是每一个可执行程序所必须包含的，
一般来说都是在启动后第一个执行的函数（如果有 init() 函数则会先执行该函数）。
 */
