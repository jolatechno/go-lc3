module github.com/jolatechno/go-lc3/src/lc3/instructions

replace (
	github.com/jolatechno/go-lc3/src/interfaces => ../../interfaces
	github.com/jolatechno/go-lc3/src/lc3/registers => ../registers
)

require (
	github.com/jolatechno/go-lc3/src/interfaces v0.0.0-20200820120128-611df7bf9bcb
	github.com/jolatechno/go-lc3/src/lc3/registers v0.0.0-20200820120128-611df7bf9bcb
)

go 1.15
