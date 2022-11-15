package xcalc

import (
	"encoding/binary"
	"fmt"
)

func Int16ToBytes(i int16) []byte {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, uint16(i))
	return buf
}

func BytesToInt16(b []byte) (int16, error) {
	if len(b) != 2 {
		return 0, fmt.Errorf("invalid bytes")
	}
	u := binary.BigEndian.Uint16(b)
	return int16(u), nil
}

func Int32ToBytes(i int32) []byte {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(i))
	return buf
}

func BytesToInt32(b []byte) (int32, error) {
	if len(b) != 4 {
		return 0, fmt.Errorf("invalid bytes")
	}
	u := binary.BigEndian.Uint32(b)
	return int32(u), nil
}

func Int64ToBytes(i int64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToInt64(b []byte) (int64, error) {
	if len(b) != 8 {
		return 0, fmt.Errorf("invalid bytes")
	}
	u := binary.BigEndian.Uint64(b)
	return int64(u), nil
}
