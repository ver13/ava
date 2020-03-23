package uuid

import (
	"bytes"
	"encoding/base32"
	"sync"

	"github.com/pborman/uuid"
)

var encoding = base32.NewEncoding("ybndrfg8ejkmcpqxot1uwisza345h769")

type uuidGmf struct {
}

var (
	f *uuidGmf

	once sync.Once
)

func init() {
	once.Do(func() {
		f = &uuidGmf{}
	})
}

func GetInstance() *uuidGmf {
	return f
}

// NewUUID is a globally unique identifier.  It is a [A-Z0-9] string 26 characters long.  It is a UUID version 4 Guid that is zbased32 encoded with the padding stripped off.
func (f *uuidGmf) NewUUID() string {
	var b bytes.Buffer
	encoder := base32.NewEncoder(encoding, &b)
	encoder.Write(uuid.NewRandom())
	encoder.Close()
	b.Truncate(26) // removes the '==' padding
	return b.String()
}
func NewUUID() string {
	return GetInstance().NewUUID()
}
