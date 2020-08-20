module github.com/jolatechno/go-lc3/src/lc3

go 1.15

replace (
	github.com/jolatechno/go-lc3/src/interfaces => ../interfaces
	github.com/jolatechno/go-lc3/src/cpu => ../cpu
	github.com/jolatechno/go-lc3/src/lc3/instructions => ./instructions
	github.com/jolatechno/go-lc3/src/lc3/memory => ./memory
	github.com/jolatechno/go-lc3/src/lc3/opcode => ./opcode
	github.com/jolatechno/go-lc3/src/lc3/registers => ./registers
)

require (
	github.com/jolatechno/go-lc3/src/cpu v0.0.0-20200817210030-60019fa04c0a
	github.com/jolatechno/go-lc3/src/interfaces v0.0.0-20200817210030-60019fa04c0a
	github.com/jolatechno/go-lc3/src/lc3/instructions v0.0.0-20200817210030-60019fa04c0a
	github.com/jolatechno/go-lc3/src/lc3/memory v0.0.0-20200817210030-60019fa04c0a
	github.com/jolatechno/go-lc3/src/lc3/opcode v0.0.0-20200817210030-60019fa04c0a
	github.com/jolatechno/go-lc3/src/lc3/registers v0.0.0-20200817210030-60019fa04c0a
)
