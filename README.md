
# Exif

[![License][License-Image]][License-Url]
[![Godoc][Godoc-Image]][Godoc-Url]
[![ReportCard][ReportCard-Image]][ReportCard-Url]
[![Coverage Status](https://coveralls.io/repos/github/evanoberholster/exif/badge.svg?branch=master)](https://coveralls.io/github/evanoberholster/exif?branch=master)
[![Build][Build-Status-Image]][Build-Status-Url]

Provides decoding of basic exif and tiff encoded data.

Suggestions and pull requests are welcome.

Example usage:

```go
package main

import (
   "fmt"
   "log"
   "os"
   "github.com/evanoberholster/exif/exif"
   "github.com/evanoberholster/exif/mknote"
)

func ExampleDecode() {
    fname := "sample1.jpg"

    f, err := os.Open(fname)
    if err != nil {
        log.Fatal(err)
    }

    // Optionally register camera makenote data parsing - currently Nikon and
    // Canon are supported.
    exif.RegisterParsers(mknote.All...)

    x, err := exif.Decode(f)
    if err != nil {
        log.Fatal(err)
    }

    camModel, _ := x.Get(exif.Model) // normally, don't ignore errors!
    fmt.Println(camModel.StringVal())

    focal, _ := x.Get(exif.FocalLength)
    numer, denom, _ := focal.Rat2(0) // retrieve first (only) rat. value
    fmt.Printf("%v/%v", numer, denom)

    // Two convenience functions exist for date/time taken and GPS coords:
    tm, _ := x.DateTime()
    fmt.Println("Taken: ", tm)

    lat, long, _ := x.LatLong()
    fmt.Println("lat, long: ", lat, ", ", long)
}
```

## Based On

Based on [https://github.com/rwcarlsen/goexif](https://github.com/rwcarlsen/goexif)

Inspired by [https://github.com/dsoprea/go-exif](https://github.com/dsoprea/go-exif)

## LICENSE

Copyright (c) 2019, Evan Oberholster & Contributors

Copyright (c) 2016, Jerry Jacobs & Contributors

Copyright (c) 2012, Robert Carlsen & Contributors

[License-Url]: https://opensource.org/licenses/BSD-2-Clause
[License-Image]: https://img.shields.io/badge/license-2%20Clause%20BSD-blue.svg?maxAge=2592000
[Godoc-Url]: https://godoc.org/github.com/evanoberholster/exif
[Godoc-Image]: https://godoc.org/github.com/evanoberholster/exif?status.svg
[ReportCard-Url]: https://goreportcard.com/report/github.com/evanoberholster/exif
[ReportCard-Image]: https://goreportcard.com/badge/github.com/evanoberholster/exif
[Build-Status-Url]: https://travis-ci.com/evanoberholster/exif?branch=master
[Build-Status-Image]: https://travis-ci.com/evanoberholster/exif.svg?branch=master
