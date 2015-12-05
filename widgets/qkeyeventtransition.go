package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QKeyEventTransition struct {
	core.QEventTransition
}

type QKeyEventTransition_ITF interface {
	core.QEventTransition_ITF
	QKeyEventTransition_PTR() *QKeyEventTransition
}

func PointerFromQKeyEventTransition(ptr QKeyEventTransition_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QKeyEventTransition_PTR().Pointer()
	}
	return nil
}

func NewQKeyEventTransitionFromPointer(ptr unsafe.Pointer) *QKeyEventTransition {
	var n = new(QKeyEventTransition)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QKeyEventTransition_") {
		n.SetObjectName("QKeyEventTransition_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QKeyEventTransition) QKeyEventTransition_PTR() *QKeyEventTransition {
	return ptr
}

func NewQKeyEventTransition2(object core.QObject_ITF, ty core.QEvent__Type, key int, sourceState core.QState_ITF) *QKeyEventTransition {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QKeyEventTransition::QKeyEventTransition")
		}
	}()

	return NewQKeyEventTransitionFromPointer(C.QKeyEventTransition_NewQKeyEventTransition2(core.PointerFromQObject(object), C.int(ty), C.int(key), core.PointerFromQState(sourceState)))
}

func NewQKeyEventTransition(sourceState core.QState_ITF) *QKeyEventTransition {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QKeyEventTransition::QKeyEventTransition")
		}
	}()

	return NewQKeyEventTransitionFromPointer(C.QKeyEventTransition_NewQKeyEventTransition(core.PointerFromQState(sourceState)))
}

func (ptr *QKeyEventTransition) Key() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QKeyEventTransition::key")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QKeyEventTransition_Key(ptr.Pointer()))
	}
	return 0
}

func (ptr *QKeyEventTransition) ModifierMask() core.Qt__KeyboardModifier {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QKeyEventTransition::modifierMask")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__KeyboardModifier(C.QKeyEventTransition_ModifierMask(ptr.Pointer()))
	}
	return 0
}

func (ptr *QKeyEventTransition) SetKey(key int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QKeyEventTransition::setKey")
		}
	}()

	if ptr.Pointer() != nil {
		C.QKeyEventTransition_SetKey(ptr.Pointer(), C.int(key))
	}
}

func (ptr *QKeyEventTransition) SetModifierMask(modifierMask core.Qt__KeyboardModifier) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QKeyEventTransition::setModifierMask")
		}
	}()

	if ptr.Pointer() != nil {
		C.QKeyEventTransition_SetModifierMask(ptr.Pointer(), C.int(modifierMask))
	}
}

func (ptr *QKeyEventTransition) DestroyQKeyEventTransition() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QKeyEventTransition::~QKeyEventTransition")
		}
	}()

	if ptr.Pointer() != nil {
		C.QKeyEventTransition_DestroyQKeyEventTransition(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
