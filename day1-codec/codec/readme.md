# codec.go

这段代码是一个包含了一些类型和函数定义的 Go 代码。它定义了一个 codec 包，用于处理编解码的操作。

首先，定义了一个结构体 Header，包含以下字段：

ServiceMethod 字段，表示服务和方法，格式为 "Service.Method"。
Seq 字段，表示由客户端选择的序列号。
Error 字段，表示错误信息。如果客户端调用没有发生错误，该字段为空；如果服务端发生错误，错误信息将被放置在 Error 中。
接下来，定义了一个接口 Codec，包含以下方法：

ReadHeader(*Header) error：读取头部信息（Header），返回一个错误（error）。
ReadBody(interface{}) error：读取消息体（Body），返回一个错误（error）。
Write(*Header, interface{}) error：写入头部信息（Header）和消息体（Body），返回一个错误（error）。
然后，定义了一个类型别名 NewCodecFunc，表示一个具有 io.ReadWriteCloser 参数并返回 Codec 接口的函数。

接着，定义了一个 Type 类型别名，表示一个字符串类型。

最后，定义了 const 常量，分别为 GobType 和 JsonType，分别表示 "application/gob" 和 "pplication/json"。

在 init 函数中，初始化了一个 NewCodecFuncMap 的映射表，将 Type 作为键，将 NewCodecFunc 函数作为值。这个映射表为 GobType 设置了 NewGobCodec 函数。

总的来说，这段代码定义了一些与编解码相关的结构、接口和函数，提供了一种方式来处理消息的编码和解码。

# gob.go
这段代码是一个 Go 语言程序包，名为 codec。它定义了一个名为 GobCodec 的结构体，该结构体实现了一个编解码器接口 Codec。

该编解码器使用 Go 的 encoding/gob 包来进行对象的编码和解码操作。GobCodec 结构体包含了连接对象、缓冲区、解码器和编码器的实例。通过使用 NewGobCodec 函数，可以创建一个新的 GobCodec 实例。

该编解码器提供了一些方法：

ReadHeader 方法用于解码数据包的头部。
ReadBody 方法用于解码数据包的主体部分。
Write 方法用于编码数据包的头部和主体部分。
Close 方法用于关闭连接。
除了以上的方法，代码还包含了一些错误处理和日志输出。

总的来说，这段代码实现了一个基于 encoding/gob 包的编解码器，用于在 Go 程序中进行数据的序列化和反序列化操作。

# server.go
这段代码是一个简单的基于 Go 语言实现的 RPC（Remote Procedure Call）库。RPC 是一种用于实现远程方法调用的协议，它允许程序通过网络调用另一个程序的方法，就像调用本地方法一样。

这个库包含了服务器端和客户端的实现，代码中的主要结构和函数如下：

- `Option`: 一个可选的参数结构，用于配置服务器端和客户端的一些选项，包括魔术数字（MagicNumber）和编解码器类型（CodecType）等。

- `Server`: 服务器端结构，包含了处理请求的方法和函数。

- `NewServer()`: 创建一个新的服务器端实例。

- `DefaultServer`: 默认的服务器实例。

- `Accept(lis net.Listener)`: 接受一个连接并为每个连接提供服务的函数。

- `ServeConn(conn io.ReadWriteCloser)`: 处理单个连接的函数。

- `serveCodec(cc codec.Codec)`: 使用指定的编解码器来处理请求。

- `request`: 请求结构，包含了请求的头部信息、参数和响应的值。

- `readRequestHeader(cc codec.Codec)`: 从连接中读取请求的头部信息。

- `readRequest(cc codec.Codec)`: 从连接中读取完整的请求，包括头部和参数。

- `sendResponse(cc codec.Codec, h *codec.Header, body interface{}, sending *sync.Mutex)`: 发送响应给客户端。

- `handleRequest(cc codec.Codec, req *request, sending *sync.Mutex, wg *sync.WaitGroup)`: 处理单个请求的函数。

这个库的主要功能是在服务器端接受连接，并使用指定的编解码器处理请求。服务器接受请求后会调用注册的 RPC 方法来获取正确的响应，并将响应返回给客户端。

请注意，这只是一个简化的代码示例，实际的使用场景可能需要更复杂的实现，例如支持多线程并发处理请求、错误处理、注册和调用远程方法等。

# main.go
这段代码是一个简单的 RPC（远程过程调用）示例，使用了来自 Geerpc 库的功能。RPC 允许在网络上的不同计算机之间进行函数调用。

代码中定义了一个 `startServer` 函数，该函数启动了一个 TCP 服务器，并将其地址发送到 `addr` 通道。

在 `main` 函数中，通过调用 `startServer` 函数来启动服务器。然后使用 `net.Dial` 函数连接到服务器。连接成功后，通过 JSON 编码器将 `geerpc.DefaultOption` 发送到服务器。

接下来，使用 Gob 编解码器创建一个 `codec.Codec` 实例，用于在网络连接上进行读写操作。然后，通过 `cc.Write` 方法发送请求，并通过 `cc.ReadHeader` 和 `cc.ReadBody` 方法读取服务器返回的响应数据。

最后，循环发送了 5 个请求，并打印出每个请求的响应数据。

请注意，这只是一个简化的示例，用于展示 RPC 的基本概念和使用 Geerpc 库的方法。

# 相关知识点
make 和 new 的区别
Go语言中 new 和 make 是两个内置函数，主要用来创建并分配类型的内存。在我们定义变量的时候，可能会觉得有点迷惑，不知道应该使用哪个函数来声明变量，其实他们的规则很简单，new 只分配内存，而 make 只能用于 slice、map 和 channel 的初始化，下面我们就来具体介绍一下。

# go map
在声明的时候不需要知道 map 的长度，map 是可以动态增长的。

未初始化的 map 的值是 nil。

key 可以是任意可以用 == 或者 != 操作符比较的类型，比如 string、int、float。所以数组、切片和结构体不能作为 key (译者注：含有数组切片的结构体不能作为 key，只包含内建类型的 struct 是可以作为 key 的），但是指针和接口类型可以。如果要用结构体作为 key 可以提供 Key() 和 Hash() 方法，这样可以通过结构体的域计算出唯一的数字或者字符串的 key

# go接口
https://go.timpaik.top/08.1.html#_8-1-1-%E6%A6%82%E5%BF%B5

# go defer()的用法和使用场景
https://tiancaiamao.gitbooks.io/go-internals/content/zh/03.4.html
https://blog.csdn.net/Cassie_zkq/article/details/108567205

工厂模式
https://www.runoob.com/design-pattern/factory-pattern.html

# go并发

# go Buffer
在 Go 语言中，`buf.Flush()` 是一种用于刷新缓冲区的方法。它通常用于刷新写入缓冲区但尚未写入底层数据源（例如文件或网络连接）的数据。当你使用缓冲读写时，数据可能首先被写入缓冲区，然后在达到一定条件时再被刷新到底层数据源。

具体而言，`buf.Flush()` 会将缓冲区中的数据写入到底层数据源，并清空缓冲区。这样可以确保之前写入缓冲区的数据被立即发送或持久化。比如，在使用 `bufio.Writer` 这个包提供的缓冲写入功能时，你可以使用 `Flush()` 方法来确保所有的数据都被写入并刷新到底层的输出流。

以下是一个简单的示例代码，演示了如何使用 `bufio.Writer` 和 `Flush()` 方法来刷新缓冲区：

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("创建文件失败：", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString("Hello, World!\n")
	writer.Flush() // 将缓冲区中的数据写入文件

	fmt.Println("数据已写入文件")
}
```

在上面的示例中，我们首先创建了一个文件并创建了一个 `bufio.Writer` 对象来包装该文件。然后我们调用 `writer.WriteString()` 将数据写入缓冲区，然后使用 `writer.Flush()` 将缓冲区中的数据刷新到文件中。最后，我们关闭文件以确保所有数据都已写入。

请注意，`Flush()` 方法并不总是必需的。当你关闭文件或程序运行结束时，缓冲区通常会自动被刷新。但是，有时你可能希望手动刷新缓冲区，以确保数据实时地写入底层数据源，而不是等待缓冲区满或程序结束时才刷新。
