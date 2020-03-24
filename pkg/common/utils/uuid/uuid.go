package uuid

import (
	"bytes"
	"encoding/base32"

	"github.com/pborman/uuid"
)

var encoding = base32.NewEncoding("ybndrfg8ejkmcpqxot1uwisza345h769")

type UUID struct {
}

// NewUUID is a globally unique identifier.  It is a [A-Z0-9] string 26 characters long.  It is a UUID version 4 Guid that is zbased32 encoded with the padding stripped off.
func (f *UUID) NewUUID() string {
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
