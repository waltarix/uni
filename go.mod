module zgo.at/uni/v2

go 1.25

require (
	github.com/mattn/go-runewidth v0.0.0-00010101000000-000000000000
	zgo.at/termtext v1.5.0
	zgo.at/zli v0.0.0-20250704045222-08cb210424f2
	zgo.at/zstd v0.0.0-20251128053228-ec259dea6715
)

require (
	github.com/clipperhouse/uax29/v2 v2.2.0 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	zgo.at/runewidth v0.1.0 // indirect
)

replace github.com/rivo/uniseg => github.com/waltarix/uniseg v0.4.7-custom-r2

replace github.com/mattn/go-runewidth => github.com/waltarix/go-runewidth v0.0.19-custom
