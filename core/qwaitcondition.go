package core

//#include "core.h"
import "C"
import (
	"log"
	"unsafe"
)

type QWaitCondition struct {
	ptr unsafe.Pointer
}

type QWaitCondition_ITF interface {
	QWaitCondition_PTR() *QWaitCondition
}

func (p *QWaitCondition) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QWaitCondition) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQWaitCondition(ptr QWaitCondition_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWaitCondition_PTR().Pointer()
	}
	return nil
}

func NewQWaitConditionFromPointer(ptr unsafe.Pointer) *QWaitCondition {
	var n = new(QWaitCondition)
	n.SetPointer(ptr)
	return n
}

func (ptr *QWaitCondition) QWaitCondition_PTR() *QWaitCondition {
	return ptr
}

func NewQWaitCondition() *QWaitCondition {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWaitCondition::QWaitCondition")
		}
	}()

	return NewQWaitConditionFromPointer(C.QWaitCondition_NewQWaitCondition())
}

func (ptr *QWaitCondition) WakeAll() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWaitCondition::wakeAll")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWaitCondition_WakeAll(ptr.Pointer())
	}
}

func (ptr *QWaitCondition) WakeOne() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWaitCondition::wakeOne")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWaitCondition_WakeOne(ptr.Pointer())
	}
}

func (ptr *QWaitCondition) DestroyQWaitCondition() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWaitCondition::~QWaitCondition")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWaitCondition_DestroyQWaitCondition(ptr.Pointer())
	}
}
