package xml

//#include "xml.h"
import "C"
import (
	"log"
	"unsafe"
)

type QXmlLexicalHandler struct {
	ptr unsafe.Pointer
}

type QXmlLexicalHandler_ITF interface {
	QXmlLexicalHandler_PTR() *QXmlLexicalHandler
}

func (p *QXmlLexicalHandler) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlLexicalHandler) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlLexicalHandler(ptr QXmlLexicalHandler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlLexicalHandler_PTR().Pointer()
	}
	return nil
}

func NewQXmlLexicalHandlerFromPointer(ptr unsafe.Pointer) *QXmlLexicalHandler {
	var n = new(QXmlLexicalHandler)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlLexicalHandler) QXmlLexicalHandler_PTR() *QXmlLexicalHandler {
	return ptr
}

func (ptr *QXmlLexicalHandler) Comment(ch string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlLexicalHandler::comment")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QXmlLexicalHandler_Comment(ptr.Pointer(), C.CString(ch)) != 0
	}
	return false
}

func (ptr *QXmlLexicalHandler) EndCDATA() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlLexicalHandler::endCDATA")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QXmlLexicalHandler_EndCDATA(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlLexicalHandler) EndDTD() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlLexicalHandler::endDTD")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QXmlLexicalHandler_EndDTD(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlLexicalHandler) EndEntity(name string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlLexicalHandler::endEntity")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QXmlLexicalHandler_EndEntity(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QXmlLexicalHandler) ErrorString() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlLexicalHandler::errorString")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlLexicalHandler_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlLexicalHandler) StartCDATA() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlLexicalHandler::startCDATA")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QXmlLexicalHandler_StartCDATA(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlLexicalHandler) StartDTD(name string, publicId string, systemId string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlLexicalHandler::startDTD")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QXmlLexicalHandler_StartDTD(ptr.Pointer(), C.CString(name), C.CString(publicId), C.CString(systemId)) != 0
	}
	return false
}

func (ptr *QXmlLexicalHandler) StartEntity(name string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlLexicalHandler::startEntity")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QXmlLexicalHandler_StartEntity(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QXmlLexicalHandler) DestroyQXmlLexicalHandler() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlLexicalHandler::~QXmlLexicalHandler")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlLexicalHandler_DestroyQXmlLexicalHandler(ptr.Pointer())
	}
}
