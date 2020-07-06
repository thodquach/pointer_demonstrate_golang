# pointer_demonstrate_golang
Demonstrate pointer change value in http request Golang

# Demonstrate how value of pointer change in a http server request 

First, run the http server:
```console
go run main.go
```

* With pointer:
```code
func New() *PointerDemon {
	return &PointerDemon{
		testValue: 1,
	}
}

func (this *PointerDemon) Demonstrate(w http.ResponseWriter, r *http.Request) {
    ...
}
```
* **curl http://localhost:8080/pointer_demonstrate**
Pointer will only init a int pointer point into variable **testValue**, and if the process still exist, the value of this variable will only store in address which the pointer point to. When we set another value to this variable, process will only update the value into the same address **0xc000018330**
So the **before** value (init value) changes continuously depent on our process run.
The terminal will print like that:
```console
╰─ go run main.go 
Adress of current PointerDemon:  0xc000018330
Before:  1
After:  2
Adress of current PointerDemon:  0xc000018330
Before:  2
After:  2
Adress of current PointerDemon:  0xc000018330
Before:  2
After:  2
```
* **curl http://localhost:8080/nonpointer_demonstrate**
When we do not use pointer, in every http request, process will create a different address to store value of varibale **testValue** 
**0xc000098040** ==> **0xc000098090** ==> **0xc000018480**
So the **before** value (init value) always equal to **1**
The terminal will print like that:
```
Adress of current PointerDemon:  0xc000098040
Before:  1
After:  2
Adress of current PointerDemon:  0xc000098090
Before:  1
After:  2
Adress of current PointerDemon:  0xc000018480
Before:  1
After:  2
```
