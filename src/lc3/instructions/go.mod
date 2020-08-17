module github.com/jolatechno/go-lc3/src/lc3/instructions

replace (
	github.com/jolatechno/go-lc3/src/interfaces => ../../interfaces
	github.com/jolatechno/go-lc3/src/lc3/opcode => ../opcode
	github.com/jolatechno/go-lc3/src/lc3/registers => ../registers
)

require (
	github.com/jolatechno/go-lc3/src/interfaces v0.0.0-20200817162726-ae7719404ec6
	github.com/jolatechno/go-lc3/src/lc3/opcode v0.0.0-20200817162726-ae7719404ec6
	github.com/jolatechno/go-lc3/src/lc3/registers v0.0.0-20200817162726-ae7719404ec6
)

go 1.13
