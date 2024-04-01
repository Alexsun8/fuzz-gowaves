//go:build ignore
// +build ignore

package main

import (
	"unsafe"

	fuzz_WE "github.com/wavesplatform/gowaves/fuzz-WE"
)

// #include <stdio.h>
// #include <errno.h>
// #include <stdint.h>
//#include <stdbool.h>
import "C"

var node_initialized = C.bool(true)

//export LLVMFuzzerTestOneInput
func LLVMFuzzerTestOneInput(data *C.char, size C.size_t) C.int {
	node_initialized = C.bool(fuzz_WE.NodeInitForFuzz())

	// TODO(mdempsky): Use unsafe.Slice once golang.org/issue/19367 is accepted.
	s := (*[1 << 30]byte)(unsafe.Pointer(data))[:size:size]
	// defer catchPanics()
	fuzz_WE.FuzzTransactionsBroadcast(s)
	return 0
}

func main() {
}

// func catchPanics() {
// 	if r := recover(); r != nil {
// 		var err string
// 		switch r.(type) {
// 		case string:
// 			err = r.(string)
// 		case runtime.Error:
// 			err = r.(runtime.Error).Error()
// 		case error:
// 			err = r.(error).Error()
// 		}
// 		if strings.Contains(err, "GO-FUZZ-BUILD-PANIC") {
// 			return
// 		} else {
// 			panic(err)
// 		}
// 	}
// }
