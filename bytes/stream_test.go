package bytes

import (
	"testing"
	"fmt"
)

func TestLittleEndianStreamImpl_CleanBuff(t *testing.T) {
	buffer := NewLittleEndianStream(make([]byte, 20, 20))
	tmp := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9}
	buffer.WriteBuff(tmp)

	tmpbuffer, _ := buffer.ReadBuff(2)
	fmt.Println(tmpbuffer)

	tmpbuffer, _ = buffer.PeekBuff(2)
	tmpbuffer, _ = buffer.ReadBuff(3)

	buffer.CleanBuff()

	fmt.Println(buffer.Data())
}