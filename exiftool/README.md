
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

BenchmarkExifDecode200-8      	      99	  13733006 ns/op	67179901 B/op	    1741 allocs/op
BenchmarkExifDecodeOld200-8   	     205	   5963736 ns/op	51937948 B/op	   16014 allocs/op
PASS

BenchmarkExifDecode200-8      	      79	  12737531 ns/op	67157449 B/op	    1613 allocs/op
BenchmarkExifDecodeOld200-8   	     204	   6041121 ns/op	51937965 B/op	   16014 allocs/op
PASS

// New
BenchmarkExifDecode200-8      	      72	  14218250 ns/op	67157460 B/op	    1613 allocs/op
BenchmarkExifDecodeOld200-8   	     178	   7048897 ns/op	51937968 B/op	   16014 allocs/op

BenchmarkExifDecode200-8      	      92	  12446991 ns/op	67157500 B/op	    1613 allocs/op
BenchmarkExifDecodeOld200-8   	     200	   6108000 ns/op	51937958 B/op	   16014 allocs/op
PASS

// OLD
BenchmarkExifDecode200-8      	      97	  12935131 ns/op	67183371 B/op	    1748 allocs/op
BenchmarkExifDecodeOld200-8   	     202	   6060535 ns/op	51937946 B/op	   16014 allocs/op

BenchmarkExifDecode200-8      	      86	  13381042 ns/op	67183382 B/op	    1748 allocs/op
BenchmarkExifDecodeOld200-8   	     202	   6053679 ns/op	51937956 B/op	   16014 allocs/op


// With Exif Reader (CR2)
BenchmarkExifDecode200-8      	    1669	    702338 ns/op	   71743 B/op	    1711 allocs/op
BenchmarkExifDecodeOld200-8   	     194	   6545560 ns/op	51937968 B/op	   16014 allocs/op

// With Exif Reader (2.CR2)
BenchmarkExifDecode200-8   	        4789	    243761 ns/op	  573898 B/op	    1184 allocs/op
BenchmarkExifDecode200-8   	        4356	    259563 ns/op	  593020 B/op	    1683 allocs/op

// With Exif Reader (13.JPG)
BenchmarkExifDecode200-8      	    3212	    358326 ns/op	   32621 B/op	     877 allocs/op
BenchmarkExifDecodeOld200-8   	     368	   3056049 ns/op	25363364 B/op	    2846 allocs/op
BenchmarkExifDecode200-8   	        3355	    347313 ns/op	   32478 B/op	     873 allocs/op
BenchmarkExifDecode200-8   	        8563	    134895 ns/op	  547342 B/op	     627 allocs/op
BenchmarkExifDecode200-8   	        8798	    134872 ns/op	  547329 B/op	     627 allocs/op
BenchmarkExifDecode200-8   	        7563	    151395 ns/op	  544724 B/op	     794 allocs/op
BenchmarkExifDecode200-8   	        7686	    160795 ns/op	  544730 B/op	     794 allocs/op
BenchmarkExifDecode200-8   	        8925	    124778 ns/op	  545589 B/op	     611 allocs/op
BenchmarkExifDecode200-8   	        9078	    127747 ns/op	  545589 B/op	     611 allocs/op
BenchmarkExifDecode200-8   	        9258	    128516 ns/op	  544951 B/op	     607 allocs/op
BenchmarkExifDecode200-8   	        9202	    131917 ns/op	  546264 B/op	     627 allocs/op
BenchmarkExifDecode200-8   	        8256	    135960 ns/op	  545912 B/op	     612 allocs/op
BenchmarkExifDecode200-8   	        9229	    127098 ns/op	  545912 B/op	     612 allocs/op

BenchmarkExifDecode200-8   	        7579	    159448 ns/op	  555621 B/op	     861 allocs/op
BenchmarkExifDecode200-8   	        8698	    129389 ns/op	  545908 B/op	     612 allocs/op