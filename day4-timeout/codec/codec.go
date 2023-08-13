package codec

import "io"

type Header struct {
	ServiceMethod string // format "Service.Method"
	Seq           uint64 // sequence number chosen by client
	Error         string //错误信息，客户端置为空，服务端如果如果发生错误，将错误信息置于 Error 中
}

type Codec interface {
	io.Closer
	ReadHeader(*Header) error // error为返回值
	ReadBody(interface{}) error
	Write(*Header, interface{}) error
}

type NewCodecFunc func(io.ReadWriteCloser) Codec

type Type string //将string用type表示

const (
	GobType  Type = "application/gob"
	JsonType Type = "pplication/json"
)

var NewCodecFuncMap map[Type]NewCodecFunc

func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}
