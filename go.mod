module main

go 1.15

replace (
	github.com/jolatechno/go-lc3/src/cpu => ./src/cpu
	github.com/jolatechno/go-lc3/src/interfaces => ./src/interfaces
	github.com/jolatechno/go-lc3/src/lc3 => ./src/lc3
)

require (
	github.com/jolatechno/go-lc3/src/cpu v0.0.0-20200820123546-021fa1f13b13
	github.com/jolatechno/go-lc3/src/interfaces v0.0.0-20200820123546-021fa1f13b13
	github.com/jolatechno/go-lc3/src/lc3 v0.0.0-20200820123546-021fa1f13b13
)
