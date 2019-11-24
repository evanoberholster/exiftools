module github.com/evanoberholster/exiftools/example

go 1.13

require (
	github.com/TylerBrock/colorjson v0.0.0-20180527164720-95ec53f28296
	github.com/evanoberholster/exiftools/exif v0.0.0-00010101000000-000000000000
	github.com/evanoberholster/exiftools/mknote v0.0.0-00010101000000-000000000000
	github.com/evanoberholster/exiftools/models v0.0.0-00010101000000-000000000000
	github.com/evanoberholster/exiftools/xmp v0.0.0-00010101000000-000000000000
	github.com/evanoberholster/filetype v1.0.5+incompatible
	github.com/fatih/color v1.7.0 // indirect
	github.com/h2non/filetype v1.0.10 // indirect
	github.com/hokaccha/go-prettyjson v0.0.0-20190818114111-108c894c2c0e // indirect
	github.com/mattn/go-colorable v0.1.4 // indirect
	github.com/mattn/go-isatty v0.0.10 // indirect
	trimmer.io/go-xmp v0.0.0-20181216222714-4f6a2fb384a3 // indirect
)

//replace github.com/evanoberholster/exiftools => ../
replace github.com/evanoberholster/exiftools/exif => ../exif

replace github.com/evanoberholster/exiftools/mknote => ../mknote

replace github.com/evanoberholster/exiftools/models => ../models

replace github.com/evanoberholster/exiftools/xmp => ../xmp

replace github.com/evanoberholster/exiftools => ../

replace github.com/evanoberholster/filetype => ../../filetype/
