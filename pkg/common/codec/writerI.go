package codec

type WriterI interface {
	Write(*Message, interface{}) error
}
