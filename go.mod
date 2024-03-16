module example.com/v2

go 1.18

require (
	example.com/library v0.0.0-00010101000000-000000000000
	github.com/mattn/go-sqlite3 v1.14.22
)

replace example.com/library => ./library
