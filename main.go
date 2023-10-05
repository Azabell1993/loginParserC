package main

/*
#cgo CFLAGS: -I/mnt/e/system/loginParserC
#cgo LDFLAGS: -L/mnt/e/system/loginParserC -lconnector
#include "connector.c"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
    // C 코드에서 export한 C 함수를 Go에서 호출
    dbHost := C.get_tokens(C.token_db_host)
    dbUser := C.get_tokens(C.token_db_user)
    dbPass := C.get_tokens(C.token_db_pass)
    dbName := C.get_tokens(C.token_db_name)

    // C 문자열을 Go 문자열로 변환
    goDbHost := C.GoString(dbHost)
    goDbUser := C.GoString(dbUser)
    goDbPass := C.GoString(dbPass)
    goDbName := C.GoString(dbName)

    fmt.Printf("db_host: %s\n", goDbHost)
    fmt.Printf("db_user: %s\n", goDbUser)
    fmt.Printf("db_pass: %s\n", goDbPass)
    fmt.Printf("db_name: %s\n", goDbName)

    // 메모리 누수 방지를 위해 C 문자열 메모리 해제
    C.free(unsafe.Pointer(dbHost))
    C.free(unsafe.Pointer(dbUser))
    C.free(unsafe.Pointer(dbPass))
    C.free(unsafe.Pointer(dbName))
}
