package main

import (
	"errors"
	"fmt"

	"github.com/ysuzuki19/collections-go/traceback"
)

func main() {
	err := level1()
	if err != nil {
		if te, ok := err.(*traceback.Error); ok {
			fmt.Println("=== StackTrace ===")
			fmt.Println(te.String())
		} else {
			fmt.Println(err)
		}
		if frames := traceback.FramesOf(err); frames.Len() > 0 {
			fmt.Println(frames.Format(func(f traceback.FormatterArgs) string {
				return fmt.Sprintf("at %s (%s:%d)", f.Function, f.File, f.Line)
			}))
		}
	}
}

func level1() error {
	return level2()
}

func level2() error {
	return level3()
}

func level3() error {
	err := doSomething()
	if err != nil {
		return traceback.Wrap(err, "level3 failed")
	}
	return nil
}

func doSomething() error {
	return errors.New("original error")
}
