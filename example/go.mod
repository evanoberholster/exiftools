module github.com/evanoberholster/exif/example

go 1.13

require (
	github.com/TylerBrock/colorjson v0.0.0-20180527164720-95ec53f28296 // indirect
	github.com/evanoberholster/exif v0.0.0+incompatible
	github.com/evanoberholster/exif/exif v0.0.0+incompatible
	github.com/evanoberholster/filetype v1.0.5+incompatible
	github.com/fatih/color v1.7.0 // indirect
	github.com/h2non/filetype v1.0.10 // indirect
	github.com/mattn/go-colorable v0.1.4 // indirect
	github.com/mattn/go-isatty v0.0.10 // indirect
	trimmer.io/go-xmp v0.0.0-20181216222714-4f6a2fb384a3 // indirect
)

replace github.com/evanoberholster/exif => ../

replace github.com/evanoberholster/exif/exif => ../exif

replace github.com/evanoberholster/filetype => ../../filetype/
