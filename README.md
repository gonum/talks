This is a repository for talks given by Gonum developers about Go and Gonum. 

## 2015

- [Introduction to Go and Concurrency](http://talks.godoc.org/github.com/gonum/talks/2015/intro_concurrency.slide)

## Running presentations locally

Some slides in stacks may include playable code examples that depend on github.com/gonum tree imports, or other
packages outside the standard library. To play these examples the presentations must be served locally using the
present tool.

The present tool is easily installed using go get:

```
go get golang.org/x/tools/cmd/present
```

Individual talks and their dependencies can then be installed and run, e.g.:

```
go get github.com/gonum/talks/2015/intro_concurrency
cd $GOPATH/src/github.com/gonum/talks/2015
present
```

The entire talks repo and dependencies may by installed:

```
go get github.com/gonum/talks/...
```


## License

Please see github.com/gonum/license for general license information, contributors, authors, etc on the Gonum suite of packages.
