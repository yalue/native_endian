package native_endian

import (
	"encoding/binary"
	"testing"
)

func TestNativeEndianness(t *testing.T) {
	native := NativeEndian()
	if native == binary.LittleEndian {
		t.Logf("Your system is little-endian.")
	} else if native == binary.BigEndian {
		t.Logf("Your system is big-endian.")
	} else {
		t.Logf("NativeEndian() did not return one of the ByteOrder values " +
			"from encoding/binary.")
	}

}
