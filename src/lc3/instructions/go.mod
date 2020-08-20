module github.com/jolatechno/go-lc3/src/lc3/instructions

replace (
	github.com/jolatechno/go-lc3/src/interfaces => ../../interfaces
	github.com/jolatechno/go-lc3/src/lc3/registers => ../registers
)

require (
	github.com/jolatechno/go-lc3/src/interfaces v0.0.0-20200820123214-15648bf02e06
	github.com/jolatechno/go-lc3/src/lc3/registers v0.0.0-20200820123214-15648bf02e06
)

go 1.15
