module github.com/jolatechno/go-lc3/src/lc3/memory

go 1.13

replace (
	github.com/jolatechno/go-lc3/src/interfaces => ../../interfaces
	github.com/jolatechno/go-lc3/src/lc3/opcode => ../opcode
)

require (
	github.com/jolatechno/go-lc3/src/interfaces v1.0.0
	github.com/jolatechno/go-lc3/src/lc3/opcode v1.0.0
)
