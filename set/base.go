package set

type Set struct {
	m map[interface{}]struct{}
}

var Exists = struct{}{}

func New(items ...interface{}) *Set {
	// 获取地址
	s := &Set{}
	// 声明map类型的数据结构
	s.m = make(map[interface{}]struct{})
	s.Add(items)
	return s
}
func (s *Set) Add(items ...interface{}) *Set {
	for _, item := range items {
		s.m[item] = Exists
	}
	return s
}

func (s *Set) Contains(item interface{}) bool {
	_, ok := s.m[item]
	return ok
}

func (s *Set) Size() int {
	return len(s.m)
}
func (s *Set) Clear() {
	s.m = make(map[interface{}]struct{})
}

// 2个集合是否相等
func (s *Set) Equal(other *Set) bool {
	if s.Size() != other.Size() {
		return false
	}
	for key := range s.m {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

//子集
func (s *Set) IsSubset(other *Set) bool {
	// s的size长于other，不用说了
	if s.Size() > other.Size() {
		return false
	}
	// 迭代遍历
	for key := range s.m {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}
