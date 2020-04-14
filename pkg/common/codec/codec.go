package codec

import (
	"io"
)

// Takes in a connection/buffer and returns a new Codec
type NewCodec func(io.ReadWriteCloser) CodecI
