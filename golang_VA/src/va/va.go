package va

/*
#cgo LDFLAGS: -LG:/golang_VA/lib/VADllInterface/ -lgolang_DllInterface
#cgo CFLAGS: -IG:/golang_VA/lib/VADllInterface/ -D_GNU_SOURCE=1

#include <stdio.h>
#include <stdlib.h>
#include "golang_DllInterface.h"
*/
import "C"

import (
	"fmt"
	//"syscall"
	"unsafe"
	//"va"
	//"web"
)

func PrintVa() {
	fmt.Println("VAA")

	s := "3.bmp"
	cs := C.CString(s)

	C.VA_fd(cs, 320, 240)

	//var p *C.int

	defer C.free(unsafe.Pointer(cs))

}

//void foo(char const *buf, size_t n);
//C.foo(&b[0], C.size_t(n))
//(*C.char)(unsafe.Pointer(&b[0]))
