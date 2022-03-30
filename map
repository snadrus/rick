package g

func Mapper[T, U any](feed []T, process func(T) U) []U {
   res := make([]U, len(feed))
   for i, v := range feed {
     res[i] = process(v)
   }
   return res
}

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
