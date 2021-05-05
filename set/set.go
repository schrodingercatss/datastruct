package set

type Set struct {
	m map[interface{}]bool
}

func NewSet() *Set {
	return &Set {
		m: map[interface{}]bool{},
	}
}

// 插入元素，并返回是否插入成功
func(set *Set) Insert(key interface{}) bool {
	defer func() {
		if r:= recover(); r != nil {}
	}()
	if _, ok := set.m[key]; ok {
		return false
	} else {
		set.m[key] = true
		return true
	}
}

// 删除元素
func(set *Set) Erase(key interface{}) {
	defer func() {
		if r:= recover(); r != nil {}
	}()
	delete(set.m, key)
}

// 清除元素
func(set *Set) Clear() {
	set.m = make(map[interface{}]bool)
}

// 查找元素
func(set *Set) Find(key interface{}) bool {
	defer func() {
		if r:= recover(); r != nil {}
	}()
	return set.m[key]
}


// 获取长度
func(set *Set) Len() int {
	return len(set.m)
}


