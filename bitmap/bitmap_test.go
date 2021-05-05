package bitmap

import (
	"testing"
)


func TestBitSet_Set(t *testing.T) {
	bitBitMap := NewBitSet(7)
	bitBitMap.Set(0)
	bitBitMap.Set(3)
	bitBitMap.Set(6)
	for i := 0; i < 7; i += 3 {
		if (*bitBitMap)[0] & (uint64(0x01) << i) == 0 {
			t.Fatal("function Get() is wrong!")
		}
	}
}


func TestBitSet_String(t *testing.T) {
	TrueResult := "1001001000000000000000000000000000000000000000000000000000000000"
	bitBitMap := NewBitSet(7)
	bitBitMap.Set(0)
	bitBitMap.Set(3)
	bitBitMap.Set(6)
	if bitBitMap.String() != TrueResult {
		t.Fatal("function String() is wrong!")
	}
}

func TestBitSet_Get(t *testing.T) {
	bitBitMap := NewBitSet(7)
	bitBitMap.Set(0)
	bitBitMap.Set(3)
	bitBitMap.Set(6)
	if !(bitBitMap.Get(0) && bitBitMap.Get(3) && bitBitMap.Get(6)) {
		t.Fatal("function Get() is wrong!")
	}
}

func TestBitSet_Clear(t *testing.T) {
	bitBitMap := NewBitSet(7)
	bitBitMap.Set(0)
	bitBitMap.Clear(0)
	bitBitMap.Set(3)
	bitBitMap.Clear(3)
	bitBitMap.Set(6)
	bitBitMap.Clear(6)
	if bitBitMap.Get(0) || bitBitMap.Get(3) || bitBitMap.Get(6) {
		t.Fatal("function Clear() is wrong!")
	}
}

