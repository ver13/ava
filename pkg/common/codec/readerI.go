package codec

type ReaderI interface {
	ReadHeader(*Message, MessageType) error
	ReadBody(interface{}) error
}
