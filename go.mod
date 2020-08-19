module main

go 1.15

replace (
	github.com/jolatechno/go-lc3/src/cpu => ./src/cpu
	github.com/jolatechno/go-lc3/src/interfaces => ./src/interfaces
	github.com/jolatechno/go-lc3/src/lc3 => ./src/lc3
)

require (
	github.com/jolatechno/go-lc3/src/cpu v1.0.0
	github.com/jolatechno/go-lc3/src/interfaces v1.0.0
	github.com/jolatechno/go-lc3/src/lc3 v1.0.0
)
