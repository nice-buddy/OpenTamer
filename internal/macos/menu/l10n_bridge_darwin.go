//go:build darwin && cgo

package menu

/*
#include <stdlib.h>
*/
import "C"

import (
	"unsafe"

	"github.com/Digital-Shane/open-tamer/internal/ui"
)

func copyCString(value string) *C.char {
	return C.CString(value)
}

//export OpenTamerL10nTranslate
func OpenTamerL10nTranslate(language *C.char, key *C.char) *C.char {
	if language == nil || key == nil {
		return nil
	}
	return copyCString(ui.BridgeString(C.GoString(language), C.GoString(key)))
}

//export OpenTamerL10nStatus
func OpenTamerL10nStatus(language *C.char, status *C.char) *C.char {
	if language == nil || status == nil {
		return nil
	}
	return copyCString(ui.StatusLabel(C.GoString(language), C.GoString(status)))
}

//export OpenTamerL10nSafety
func OpenTamerL10nSafety(language *C.char, reason *C.char) *C.char {
	if language == nil || reason == nil {
		return nil
	}
	return copyCString(ui.SafetyLabel(C.GoString(language), C.GoString(reason)))
}

//export OpenTamerL10nFree
func OpenTamerL10nFree(value *C.char) {
	if value == nil {
		return
	}
	C.free(unsafe.Pointer(value))
}
