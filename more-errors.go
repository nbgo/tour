package main
import (
	"github.com/phayes/errors"
	stderrors "errors"
	"fmt"
)

var (
	Err1 = stderrors.New("error1")
	Err2 = stderrors.New("error2")
)

func main() {
	err1 := errors.Wrap(Err1, Err2)
	err2 := errors.Append(Err2, Err1)
	fmt.Printf("%T: %v\n", err1, err1)
	fmt.Printf("%T: %v\n", err2, err2)
}