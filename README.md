# undergo
Just another pet project.

### Manage dependencies

        docker run \
          --rm -it -v $PWD:/go/src/github.com/stsh89/undergo \
          -w /go/src/github.com/stsh89/undergo \
          golang bash -c "go get -u github.com/tools/godep && godep save ./..."

### Run tests

        docker run \
          --rm -it -v $PWD:/go/src/github.com/stsh89/undergo \
          -w /go/src/github.com/stsh89/undergo \
          golang go test -v ./...

### Check app functionality

        docker run \
          --rm -it -p 3000:3000 -v $PWD:/go/src/github.com/stsh89/undergo \
          -w /go/src/github.com/stsh89/undergo \
          -e PORT=3000 \
          golang go run main.go

        curl localhost:3000
