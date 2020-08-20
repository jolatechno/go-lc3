module github.com/jolatechno/go-lc3/src/lc3/instructions

replace (
	github.com/jolatechno/go-lc3/src/interfaces => ../../interfaces
	github.com/jolatechno/go-lc3/src/lc3/opcode => ../opcode
	github.com/jolatechno/go-lc3/src/lc3/registers => ../registers
)

require (
	github.com/jolatechno/go-lc3/src/interfaces v0.0.0-20200817210030-60019fa04c0a
	github.com/jolatechno/go-lc3/src/lc3/opcode v0.0.0-20200817210030-60019fa04c0a
	github.com/jolatechno/go-lc3/src/lc3/registers v0.0.0-20200817210030-60019fa04c0a
)

go 1.15
