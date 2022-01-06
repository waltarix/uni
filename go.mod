module zgo.at/uni/v2

go 1.25

require (
	github.com/mattn/go-runewidth v0.0.0-00010101000000-000000000000
	zgo.at/termtext v1.5.0
	zgo.at/zli v0.0.0-20240614180544-47534b1ce136
	zgo.at/zstd v0.0.0-20240827020003-f7ed9341ec67
)

require (
	github.com/clipperhouse/uax29/v2 v2.2.0 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	zgo.at/runewidth v0.1.0 // indirect
)

replace github.com/rivo/uniseg => github.com/waltarix/uniseg v0.4.7-custom-r2

replace github.com/mattn/go-runewidth => github.com/waltarix/go-runewidth v0.0.19-custom
