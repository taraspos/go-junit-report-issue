# Code to reproduce [go-juint-report Issue #71](https://github.com/jstemmer/go-junit-report/issues/71)

Running with:

    go test -v ./... 2>&1 | go-junit-report -set-exit-code; echo Exit Code $?
