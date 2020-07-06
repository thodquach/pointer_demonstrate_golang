package nonpointer

import (
	"fmt"
	"net/http"
)

type Interface interface {
	Demonstrate(w http.ResponseWriter, r *http.Request)
}

type NonPointerDemon struct {
	testValue int
}

func New() NonPointerDemon {
	return NonPointerDemon{
		testValue: 1,
	}
}

func (this NonPointerDemon) Demonstrate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Adress of current PointerDemon: ", &this.testValue)
	fmt.Println("Before: ", this.testValue)
	this.testValue = 2
	fmt.Println("After: ", this.testValue)
	fmt.Fprintf(w, "Demonstrate: PointerDemon!")
}
