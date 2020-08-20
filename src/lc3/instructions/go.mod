module github.com/jolatechno/go-lc3/src/lc3/instructions

replace (
	github.com/jolatechno/go-lc3/src/interfaces => ../../interfaces
	github.com/jolatechno/go-lc3/src/lc3/registers => ../registers
)

require (
	github.com/jolatechno/go-lc3/src/interfaces v0.0.0-20200820123546-021fa1f13b13
	github.com/jolatechno/go-lc3/src/lc3/registers v0.0.0-20200820123546-021fa1f13b13
)

go 1.15
