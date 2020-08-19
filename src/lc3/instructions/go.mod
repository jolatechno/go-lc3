module github.com/jolatechno/go-lc3/src/lc3/instructions

replace (
	github.com/jolatechno/go-lc3/src/interfaces => ../../interfaces
	github.com/jolatechno/go-lc3/src/lc3/opcode => ../opcode
	github.com/jolatechno/go-lc3/src/lc3/registers => ../registers
)

require (
	github.com/jolatechno/go-lc3/src/interfaces v1.0.0
	github.com/jolatechno/go-lc3/src/lc3/opcode v1.0.0
	github.com/jolatechno/go-lc3/src/lc3/registers v1.0.0
)

go 1.15
