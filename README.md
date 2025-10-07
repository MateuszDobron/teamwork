initial state benchmark results:
import
3915590 ns/op       1464166 B/op      20035 allocs/op
export
84287 ns/op	        4264 B/op	      4 allocs/op

after refinement benchmark results:
import
3155526 ns/op	  652122 B/op	    10036 allocs/op
export
130630 ns/op	   12456 B/op	    5 allocs/op


sum comparison
inital:
3999877 ns/op   1468430 B/op    20039 allocs/op
after:
3286156 ns/op   664578 B/op     10041 allocs/op
initial/after:
1.217           2.209            1.995

So the cli app is now ~20% faster and uses ~50% less memory