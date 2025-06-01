package parentchild

import (
	"errors"
	"fmt"
)

//Using channels to send errors from child to parent.

func Child(errCh chan error) {
	// simulate error
	errCh <- errors.New("something went wrong in child")
}

func ERROR_PROPAGRATION() {
	errCh := make(chan error)

	go Child(errCh)

	err := <-errCh
	fmt.Println("Parent received error:", err)
}
