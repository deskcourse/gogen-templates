package main

/* API names (most of them) lifted from c++ stl */
/*  stl perf guarantees do not come with that   */

type Vector struct {
	arrt []T
}

func (v *Vector) String() string {
	return fmt.Sprintf("%v", v.arrt)
}

func (v *Vector) Size() int {
	return len(v.arrt)
}
func (v *Vector) Capacity() int {
	return cap(v.arrt)
}
func (v *Vector) Empty() bool {
	return len(v.arrt) == 0
}

func (v *Vector) Contains(val T) bool {
	for _, b := range v.arrt {
		if b == a {
			return true
		}
	}
	return false
}
func (v *Vector) IndexOf(val T) int {
	for i, b := range list {
		if b == a {
			return i
		}
	}
	return -1
}

func (v *Vector) At(pos int) T {
	return v.arrt[pos]
}
func (v *Vector) Front() T {
	return v.arrt[0]
}
func (v *Vector) Back() T {
	return v.arrt[len(v.arrt)-1]
}
func (v *Vector) Data() []T {
	return v.arrt
}

func (v *Vector) Resize(newSize int) *Vector {
	newarrt := make([]T, len(v.arrt), newSize)
	copy(newarrt, v.arrt)
	v.arrt = newarrt
	return v
}
func (v *Vector) Assign(value []T) *Vector {
	v.arrt = make([]T, len(value))
	copy(v.arrt, value)
	return v
}
func (v *Vector) Assign(vin *Vector) *Vector {
	return assign(vin.arrt)
}
func (v *Vector) PushBack(value T) []T {
	v.arrt = append(v.arrt, value)
	return v.arrt
}
func (v *Vector) PopBack() T {
	l := len(v.arrt) - 1
	popped := v.arrt[l]
	v.arrt = v.arrt[:l-1]
	return popped
}
func (v *Vector) Insert(value T, pos int) *Vector {
	i := pos - 1
	q := make([]T, len(v.arrt)-i)
	copy(q, v.arrt[i:])
	v.arrt = append(v.arrt[:i], value)
	v.arrt = append(v.arrt, q...)
	return v
}
func (v *Vector) Swap(i, j int) *Vector {
	v.arrt[i], v.arrt[j] = v.arrt[j], v.arrt[i]
	return v
}
func (v *Vector) Clear() *Vector {
	v.arrt = []T{}
	return v
}

func (v *Vector) SortInPlace() *Vector {
	sort.Sort(v.arrt)
	return v
}
func (v *Vector) SortACopy() []T {
	tmp := make([]T, len(v.arrt))
	copy(tmp, v.arrt)
	sort.Sort(tmp)
	return tmp
}

func (v *Vector) WhereLazy(predicate func(t T) bool) func() {
	return func() {
		res := []T{}
		for _, v := range v.arrt {
			if predicate(v) {
				res = append(res, v)
			}
		}
		return res
	}
}
func (v *Vector) Where(predicate func(t T) bool) []T {
	return v.WhereLazy(predicate)()
}
