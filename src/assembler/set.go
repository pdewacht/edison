package main

type Set struct{ hi, lo uint64 }

func Singleton(el int) Set {
	if el >= 0 && el < 64 {
		return Set{lo: uint64(1) << el}
	} else if el >= 64 && el < 128 {
		return Set{hi: uint64(1) << (el - 64)}
	} else {
		panic("set element out of range")
	}
}

func Construct(elements ...int) Set {
	set := Set{}
	for _, el := range elements {
		set = set.Add(el)
	}
	return set
}

func ConstructFromString(s string) Set {
	set := Set{}
	for _, el := range s {
		set = set.AddChar(el)
	}
	return set
}

func (set Set) Contains(el int) bool {
	return set.Intersect(Singleton(el)) != Set{}
}

func (set Set) ContainsChar(el rune) bool {
	return set.Contains(int(el))
}

func (set Set) Add(el int) Set {
	return set.Union(Singleton(el))
}

func (set Set) AddChar(el rune) Set {
	return set.Add(int(el))
}

func (set Set) Remove(el int) Set {
	return set.Except(Singleton(el))
}

func (set Set) RemoveChar(el rune) Set {
	return set.Remove(int(el))
}

func (set Set) Union(other Set) Set {
	return Set{hi: set.hi | other.hi, lo: set.lo | other.lo}
}

func (set Set) Intersect(other Set) (result Set) {
	return Set{hi: set.hi & other.hi, lo: set.lo & other.lo}
}

func (set Set) Except(other Set) (result Set) {
	return Set{hi: set.hi & ^other.hi, lo: set.lo & ^other.lo}
}
