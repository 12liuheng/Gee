# 同步调用和异步调用
同步调用和异步调用是两种不同的编程调用方式，其区别在于程序在等待调用结果时的行为不同。

同步调用指的是程序发出一个调用请求，然后等待调用结果返回后再继续执行后面的代码。在等待调用结果的过程中，程序会一直阻塞等待，直到结果返回为止。如果调用过程出现问题，程序也会一直阻塞等待直到出现超时或者异常。

异步调用指的是程序发出一个调用请求，然后立即开始执行后面的代码，不会等待调用结果返回。当调用结果返回后，程序会根据预设的回调函数来处理结果。因此，异步调用不会阻塞程序的执行，提高了程序的性能和并发能力。同时，由于异步调用不会等待结果返回，所以当调用过程出现问题时，程序可以继续执行后面的代码。

总的来说，同步调用适用于需要立即获得结果并且结果需要立即处理的场景，而异步调用适用于处理时间比较长或者需要处理多个调用请求的场景。同时，异步调用的实现相对比较复杂，需要考虑回调函数的实现和多线程的并发问题。

# go管道
http://c.biancheng.net/uploads/allimg/180817/1-1PQG035203K.jpg

# 变量
函数体内声明局部变量时，应使用简短声明语法 := 使用变量的首选形式，但是它只能被用在函数体内，而不可以用于全局变量的声明与赋值。使用操作符 := 可以高效地创建一个新的变量，称之为初始化声明

# init()函数
变量除了可以在全局声明中初始化，也可以在 init 函数中初始化。这是一类非常特殊的函数，它不能够被人为调用，而是在每个包完成初始化后自动执行，并且执行优先级比 main 函数高.

# 位运算
按位与 &：
按位或 |：
按位异或 ^：
位清除 &^：将指定位置上的值设置为 0。

# 类型别名
当你在使用某个类型时，你可以给它起另一个名字，然后你就可以在你的代码中使用新的名字（用于简化名称或解决名称冲突）。

在 type TZ int 中，TZ 就是 int 类型的新名称（用于表示程序中的时区），然后就可以使用 TZ 来操作 int 类型的数据。