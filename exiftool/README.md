
// JPG file with ExifData
BenchmarkParseExifHeader200-8   	  408568	      2915 ns/op	    1088 B/op	       3 allocs/op
BenchmarkParseExif2Header200-8   	    1047	    956161 ns/op	 8386637 B/op	      15 allocs/op

// JPG file with ExifData #2
BenchmarkParseExifHeader200-8   	  412591	      2883 ns/op	    1088 B/op	       3 allocs/op
BenchmarkParseExif2Header200-8   	     996	   1051139 ns/op	 8386637 B/op	      15 allocs/op

// CR2 file
BenchmarkParseExifHeader200-8   	  409156	      2654 ns/op	    1088 B/op	       3 allocs/op
BenchmarkParseExif2Header200-8   	      68	  16370650 ns/op	67106902 B/op	      18 allocs/op

// JPG file with No ExifData
BenchmarkParseExifHeader200-8   	     261	   4603787 ns/op	    1056 B/op	       2 allocs/op
BenchmarkParseExif2Header200-8   	     254	   4749712 ns/op	 1046565 B/op	      11 allocs/op
