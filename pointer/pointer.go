package pointer

import (
	"fmt"
	"net/http"
)

type Interface interface {
	Demonstrate(w http.ResponseWriter, r *http.Request)
}

type PointerDemon struct {
	testValue int
}

func New() *PointerDemon {
	return &PointerDemon{
		testValue: 1,
	}
}

func (this *PointerDemon) Demonstrate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Adress of current PointerDemon: ", &this.testValue)
	fmt.Println("Before: ", this.testValue)
	this.testValue = 2
	fmt.Println("After: ", this.testValue)
	fmt.Fprintf(w, "Demonstrate: PointerDemon!")
}
