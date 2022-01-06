module zgo.at/uni/v2

go 1.22

require (
	github.com/mattn/go-runewidth v0.0.0-00010101000000-000000000000
	zgo.at/termtext v1.5.0
	zgo.at/zli v0.0.0-20240522161040-91dbcffb2960
	zgo.at/zstd v0.0.0-20240521013615-10baa641d7d0
)

require (
	github.com/rivo/uniseg v0.4.7 // indirect
	zgo.at/runewidth v0.1.0 // indirect
)

replace github.com/rivo/uniseg => github.com/waltarix/uniseg v0.4.7-custom

replace github.com/mattn/go-runewidth => github.com/waltarix/go-runewidth v0.0.15-custom
