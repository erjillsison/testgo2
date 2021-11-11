package dir

import (
	"github.com/erjillsison/testgo"
	"github.com/erjillsison/testgo2/dironly/packageindir"
)

func SomeFunc() {
	testgo.SomeFunc2()
	packageindir.Func()
}
