// +build armbe arm64be ppc64 mips mips64 mips64p32 ppc sparc sparc64 s390 s390x

package native_endian

import (
	"encoding/binary"
)

func NativeEndian() binary.ByteOrder {
	return binary.BigEndian
}
