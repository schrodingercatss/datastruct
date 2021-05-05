package set

import "testing"

func TestSet_Insert(t *testing.T) {
	s := NewSet()
	s.Insert("abc")
	if !s.Find("abc") || s.Find("bcd") {
		t.Fatal("function Insert() is wrong!")
	}
}

func TestSet_Find(t *testing.T) {
	s := NewSet()
	s.Insert("abc")
	if !s.Find("abc") || s.Find("bcd") {
		t.Fatal("function Find() is wrong!")
	}
}

func TestSet_Len(t *testing.T) {
	s := NewSet()
	s.Insert("abc")
	if s.Len() != 1 {
		t.Fatal("function Len() is wrong!")
	}
}

func TestSet_Erase(t *testing.T) {
	s := NewSet()
	s.Insert(123)
	s.Insert("abc")
	s.Erase(123)
	if s.Len() != 1 {
		t.Fatal("function Erase() is wrong!")
	}
}

func TestSet_Clear(t *testing.T) {
	s := NewSet()
	s.Insert(123)
	s.Insert("abc")
	s.Clear()
	if s.Len() != 0 {
		t.Fatal("function Clear() is wrong!")
	}
}