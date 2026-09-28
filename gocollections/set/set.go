package set


type Set struct {
	items []int
}

func NewSet() *Set {
	return &Set{}
}

func(s *Set) Add(value int) {
	if len(s.items)<1 {
		s.items = append(s.items, value)
		return
	}
	for _,ex:=range s.items {
		if(value==ex){
			return
		}
	}
	s.items = append(s.items, value)
}

