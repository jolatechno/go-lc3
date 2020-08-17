module github.com/jolatechno/go-lc3/src/lc3/memory

go 1.13

replace (
	github.com/jolatechno/go-lc3/src/interfaces => ../../interfaces
	github.com/jolatechno/go-lc3/src/lc3/opcode => ../opcode
)

require (
	github.com/jolatechno/go-lc3/src/interfaces v0.0.0-20200817162726-ae7719404ec6
	github.com/jolatechno/go-lc3/src/lc3/opcode v0.0.0-20200817162726-ae7719404ec6
)
