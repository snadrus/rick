package g

func Mapper[T, U any](feed []T, process func(T) U) []U {
   res := make([]U, len(feed))
   for i, v := range feed {
     res[i] = process(v)
   }
   return res
}

// Reduce is the classic array reducer.
// Ex: sum := g.Reduce([]int[1, 3, 5, 7}, 0, func(a, soFar int) int { return soFar + a })
// Also useful for map filtering:
//   type mt map[int]int
//   m := map[int]int{2:4,6:8,10:12}
//   filteredMap := g.Reduce(g.Keys(m), mt{}, func(k int, soFar mt) mt {
//      if m[k] > 4 {
//         soFar[k] = m[k]
//      }
//      return soFar
//   }
func Reduce[T, U any](feed []T, base U, process func(T, U) U) U {
   for _, v := range feed {
      base = process(v, base)
   }
   return base
}

func Filter[T any](feed []T, cond func(T)bool) []T {
   var res []T
   for _, v := range feed {
     if cond(v) {
        res = append(res, v)
     }
   }
   return res
}

type Set[T comparable] map[T] bool
func Union[T comparable](maps Set[T]...) Set[T] {
   res := map[T]bool{}
   for _, m := range maps {
      for k := range m {
         res[k] = true
      }
   }
   return res
}

// Intersection implements Set intersections
func Intersection[T comparable](maps Set[T]...) Set[T] {
   if len(maps) == 0 {
      return nil 
   }
   shortest := 0
   for i, m := range maps {
      if len(m) < len(maps[shortest]) {
         shortest = i
      }
   }
   res := make(Set[T], len maps[shortest])
   for k := range maps[shortest] {
      res[k] = true
   }
   for _, maps := range [][]map[T]bool{ maps[:shortest], maps[shortest+1:] } {
      for _, m := range maps {
         for k := range res {
            if !m[k] {
               delete(res, k)
            }
         }
      }
   }
   return res
}

// Keys helps get keys from a map for various processing. 
func Keys[T, U any](m map[T]U) []T {
   res := make([]T, len(m))
   for k := range m {
      res = append(res, k)
   }
   return res
}

// TODO maptoys:
func MapToys[K comparable, V any](m map[K]V, ...toys[K,V] MapToy) map[K]V {
   for _, t := range toys {
      m = t(m)
   }
}

type MapToy[K comparable, V any] func(map[K]V) map[K]V)
func MapToyClone[K comparable, V any](m map[K]V) map[K]V { x:=make(map[K]V, len(m)); for k, v := range m { x[k]=v }; return x }
// MapToyFilter on keys
func MapToyFilter[K comparable, V any](op Op[K]) func(m map[K]V, op Op[K]) map[K]V {
   return func(map[K]V) map[K]V) {
      x := make(map[K]V, len(m))
      for k, v := range m {
         if op[k] {
            x[k] = v
         }
      }
      return x
   }
}
type Op[K comparable] func(k K) bool 

// In Usage: g.MapToy(mapA, g.MapToyFilter(g.In(g.ToSet(mapB.Keys()))
func In[V comparable](haystack Set[V]) Op {
   return func(needle V) bool {
      return haystack[needle]
   }
}

func FlowOf[V any](v []V) chan V {
   ch := make(chan V)
   go func(){
      defer close(ch)
      for _, val  := range v {
         v <- ch
      }
   }()
   return ch
}
func Where[V comparable](ch chan V, f Op[V]) chan V {
   ch2 := make(chan V)
   go func() {
      defer close(ch2)
      for x  := range ch {
         if f(x) {
            ch2 <- x
         }
      }   
   }()
   return ch2
}
func Emit[V any](ch chan V) []V {
   var res []V
   for x := range ch {
      res = append(res, x)
   }
   return res
}
func EmitMap[K comparable, V any](ch chan K, m map[K]V) map[K]V {
   res := map[K]V{}
   for x := range ch {
      res[k] = m[k]
   }
   return res
}

//f := g.FlowOf(g.Keys(m))
//f := g.Where(f, In(wantedKeys))
// g.Emit(f) || g.EmitMap(f, srcMap)
// Short version: g.EmitMap(g.Where(g.FlowOf(g.Keys(m)), g.In(wantedKeys)), m)
          
