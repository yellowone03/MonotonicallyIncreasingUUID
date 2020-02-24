package serverbase

import (
	"bytes"
	"fmt"

	"github.com/rs/xid"
)

// UUIDGenerator uuid
type UUIDGenerator struct {
	cnt     uint32
	cntLoop uint32
	guid    string
}

const defaultCntLoop = 4294967295

// NewUUIDGenerator new
func NewUUIDGenerator() *UUIDGenerator {
	ug := new(UUIDGenerator)
	ug.cnt = 0
	ug.cntLoop = defaultCntLoop
	ug.guid = xid.New().String()

	return ug
}

// GetMonotonicallyIncreasingUUID getuuid
func (ug *UUIDGenerator) GetMonotonicallyIncreasingUUID() string {
	ug.cnt++
	if ug.cnt%ug.cntLoop == 0 {
		ug.guid = xid.New().String()
		ug.cnt = 0
	}
	hexCnt := fmt.Sprintf("%08x", ug.cnt)

	var buffer bytes.Buffer
	buffer.WriteString(hexCnt)
	buffer.WriteString(ug.guid)

	return buffer.String()
}
