package aliU

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

// FastFail is a GomegaFailHandler that, when run with `go test -failfast` , fails immidiately and halts other long running tests prints the line number of the test that failed
// from conversation with @onsi https://github.com/onsi/gomega/issues/660
// this makes it possible to fail fast and not wait for all the tests to finish if they are going to fail
// usage:
// in your test, as the first line of the test, add this:
//
//	RegisterFailHandler(aliU.FastFail)
func FastFail(failure string, args ...int) {
	// find the line number of the test that failed
	var culprit *runtime.Func
	var fileName string
	var line int
	// TODO: probably this needs to start from 2, not 1
	for i := 1; i < 10; i++ {
		pc, _, _, ok := runtime.Caller(i)
		details := runtime.FuncForPC(pc)
		if ok && details != nil && !strings.Contains(details.Name(), "gomega") {
			culprit = details
			fileName, line = culprit.FileLine(pc)
			break
		}
	}

	if culprit != nil {
		fmt.Printf("called from %s\n%s:%d:", culprit.Name(), fileName, line)
	} else {
		fmt.Println("couldn't pinpoint the culprit")
	}

	//debug.PrintStack()
	fmt.Fprintf(os.Stderr, "\n%s %v", failure, args)
	os.Exit(1)
}
