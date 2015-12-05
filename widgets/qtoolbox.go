package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"log"
	"unsafe"
)

type QToolBox struct {
	QFrame
}

type QToolBox_ITF interface {
	QFrame_ITF
	QToolBox_PTR() *QToolBox
}

func PointerFromQToolBox(ptr QToolBox_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QToolBox_PTR().Pointer()
	}
	return nil
}

func NewQToolBoxFromPointer(ptr unsafe.Pointer) *QToolBox {
	var n = new(QToolBox)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QToolBox_") {
		n.SetObjectName("QToolBox_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QToolBox) QToolBox_PTR() *QToolBox {
	return ptr
}

func (ptr *QToolBox) Count() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBox::count")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QToolBox_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QToolBox) CurrentIndex() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBox::currentIndex")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QToolBox_CurrentIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QToolBox) SetCurrentIndex(index int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBox::setCurrentIndex")
		}
	}()

	if ptr.Pointer() != nil {
		C.QToolBox_SetCurrentIndex(ptr.Pointer(), C.int(index))
	}
}

func NewQToolBox(parent QWidget_ITF, f core.Qt__WindowType) *QToolBox {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBox::QToolBox")
		}
	}()

	return NewQToolBoxFromPointer(C.QToolBox_NewQToolBox(PointerFromQWidget(parent), C.int(f)))
}

func (ptr *QToolBox) AddItem2(w QWidget_ITF, text string) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBox::addItem")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QToolBox_AddItem2(ptr.Pointer(), PointerFromQWidget(w), C.CString(text)))
	}
	return 0
}

func (ptr *QToolBox) AddItem(widget QWidget_ITF, iconSet gui.QIcon_ITF, text string) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBox::addItem")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QToolBox_AddItem(ptr.Pointer(), PointerFromQWidget(widget), gui.PointerFromQIcon(iconSet), C.CString(text)))
	}
	return 0
}

func (ptr *QToolBox) ConnectCurrentChanged(f func(index int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBox::currentChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QToolBox_ConnectCurrentChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentChanged", f)
	}
}

func (ptr *QToolBox) DisconnectCurrentChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBox::currentChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QToolBox_DisconnectCurrentChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentChanged")
	}
}

//export callbackQToolBoxCurrentChanged
func callbackQToolBoxCurrentChanged(ptrName *C.char, index C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBox::currentChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "currentChanged").(func(int))(int(index))
}

func (ptr *QToolBox) CurrentWidget() *QWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBox::currentWidget")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QToolBox_CurrentWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QToolBox) IndexOf(widget QWidget_ITF) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBox::indexOf")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QToolBox_IndexOf(ptr.Pointer(), PointerFromQWidget(widget)))
	}
	return 0
}

func (ptr *QToolBox) InsertItem(index int, widget QWidget_ITF, icon gui.QIcon_ITF, text string) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBox::insertItem")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QToolBox_InsertItem(ptr.Pointer(), C.int(index), PointerFromQWidget(widget), gui.PointerFromQIcon(icon), C.CString(text)))
	}
	return 0
}

func (ptr *QToolBox) InsertItem2(index int, widget QWidget_ITF, text string) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBox::insertItem")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QToolBox_InsertItem2(ptr.Pointer(), C.int(index), PointerFromQWidget(widget), C.CString(text)))
	}
	return 0
}

func (ptr *QToolBox) IsItemEnabled(index int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBox::isItemEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QToolBox_IsItemEnabled(ptr.Pointer(), C.int(index)) != 0
	}
	return false
}

func (ptr *QToolBox) ItemText(index int) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBox::itemText")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QToolBox_ItemText(ptr.Pointer(), C.int(index)))
	}
	return ""
}

func (ptr *QToolBox) ItemToolTip(index int) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBox::itemToolTip")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QToolBox_ItemToolTip(ptr.Pointer(), C.int(index)))
	}
	return ""
}

func (ptr *QToolBox) RemoveItem(index int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBox::removeItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QToolBox_RemoveItem(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QToolBox) SetCurrentWidget(widget QWidget_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBox::setCurrentWidget")
		}
	}()

	if ptr.Pointer() != nil {
		C.QToolBox_SetCurrentWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QToolBox) SetItemEnabled(index int, enabled bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBox::setItemEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		C.QToolBox_SetItemEnabled(ptr.Pointer(), C.int(index), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QToolBox) SetItemIcon(index int, icon gui.QIcon_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBox::setItemIcon")
		}
	}()

	if ptr.Pointer() != nil {
		C.QToolBox_SetItemIcon(ptr.Pointer(), C.int(index), gui.PointerFromQIcon(icon))
	}
}

func (ptr *QToolBox) SetItemText(index int, text string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBox::setItemText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QToolBox_SetItemText(ptr.Pointer(), C.int(index), C.CString(text))
	}
}

func (ptr *QToolBox) SetItemToolTip(index int, toolTip string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBox::setItemToolTip")
		}
	}()

	if ptr.Pointer() != nil {
		C.QToolBox_SetItemToolTip(ptr.Pointer(), C.int(index), C.CString(toolTip))
	}
}

func (ptr *QToolBox) Widget(index int) *QWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBox::widget")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QToolBox_Widget(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QToolBox) DestroyQToolBox() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBox::~QToolBox")
		}
	}()

	if ptr.Pointer() != nil {
		C.QToolBox_DestroyQToolBox(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
