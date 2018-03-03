# undergo
Just another pet project.

### Manage dependencies

        docker run --rm -it -v $PWD:/go/src/github.com/stsh89/undergo -w /go/src/github.com/stsh89/undergo golang bash -c "go get -u github.com/tools/godep && godep save ./..."

### Run test

        docker run --rm -it -v $PWD:/go/src/github.com/stsh89/undergo -w /go/src/github.com/stsh89/undergo golang go test -v ./...
