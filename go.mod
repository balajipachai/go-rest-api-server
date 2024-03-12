module example.com/v2

go 1.18

require example.com/library v0.0.0-00010101000000-000000000000

require github.com/gorilla/mux v1.8.1 // indirect

replace example.com/library => ./library
