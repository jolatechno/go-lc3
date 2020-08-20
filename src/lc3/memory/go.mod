module github.com/jolatechno/go-lc3/src/lc3/memory

go 1.15

replace (
	github.com/jolatechno/go-lc3/src/interfaces => ../../interfaces
	github.com/jolatechno/go-lc3/src/lc3/opcode => ../opcode
)

require (
	github.com/jolatechno/go-lc3/src/interfaces v0.0.0-20200817210030-60019fa04c0a
	github.com/jolatechno/go-lc3/src/lc3/opcode v0.0.0-20200817210030-60019fa04c0a
)
