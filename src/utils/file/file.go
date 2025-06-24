package file

import (
	"github.com/ihangsen/common/src/catch"
	"os"
)

func Load(filename string) string {
	return string(catch.Try1(os.ReadFile(filename)))
}
