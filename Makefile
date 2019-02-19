MAKEFLAGS = -s

mongo.go: mongo.y
	goyacc -o mongo.go mongo.y
	gofmt -w mongo.go

clean:
	rm -f y.output mongo.go
