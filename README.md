# rick
Functional programming funcs in Go without a fuss.

Should probably change to iterators because 
 - ch goroutines leak if ignored
 - slices everywhere are wasteful allocs
 - super-controllers can't handle multiple generics composed

start:
```go
// Mapper maps an iterator's value into a slice.
// Usage; g.Mapper(g.Keys(myMap), func(i int) j str { return myMap[i] * 4 })
// TODO this SHOULD return an iterator too. NO lists.  Would end in .ToList() for this result. 
func Mapper[T, U any](in Iterator[T], conv func(T) U) []U {
	var res []U
	for v, ok := in(); ok; v, ok = in() {
		res = append(res, conv(v))
	}
	return res
}

// An Iterator walks values over some list-like thing.
type Iterator[T any] func() (T, bool)

// Keys returns a map's keys as an Iterator.
func Keys[T comparable, U any](in map[T]U) Iterator[T] {
	r := reflect.ValueOf(in).MapRange()
	return func() (T, bool) {
		ok := r.Next()
    if !ok {
      return T{}, false
    }
		return r.Key().Interface().(T), ok
	}
}

func Iter[T any](t []t) Iterator[T] {
  l := len(t)
  this := 0
  return func()(T, bool) {
    if this >= l {
      return T{}, false
    }
    v, this := t[this], this+1
    return v, ok
  }
}

func (it Iterator[T])ToList() []T {
  var res []T
  for v, ok := it(); ok; v, ok = it() {
    res = append(res, T)
  }
  return res
}
```
