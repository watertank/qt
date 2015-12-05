package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"strings"
	"unsafe"
)

type QAbstractProxyModel struct {
	QAbstractItemModel
}

type QAbstractProxyModel_ITF interface {
	QAbstractItemModel_ITF
	QAbstractProxyModel_PTR() *QAbstractProxyModel
}

func PointerFromQAbstractProxyModel(ptr QAbstractProxyModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractProxyModel_PTR().Pointer()
	}
	return nil
}

func NewQAbstractProxyModelFromPointer(ptr unsafe.Pointer) *QAbstractProxyModel {
	var n = new(QAbstractProxyModel)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAbstractProxyModel_") {
		n.SetObjectName("QAbstractProxyModel_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAbstractProxyModel) QAbstractProxyModel_PTR() *QAbstractProxyModel {
	return ptr
}

func (ptr *QAbstractProxyModel) Buddy(index QModelIndex_ITF) *QModelIndex {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractProxyModel::buddy")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QAbstractProxyModel_Buddy(ptr.Pointer(), PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QAbstractProxyModel) CanDropMimeData(data QMimeData_ITF, action Qt__DropAction, row int, column int, parent QModelIndex_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractProxyModel::canDropMimeData")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAbstractProxyModel_CanDropMimeData(ptr.Pointer(), PointerFromQMimeData(data), C.int(action), C.int(row), C.int(column), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QAbstractProxyModel) CanFetchMore(parent QModelIndex_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractProxyModel::canFetchMore")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAbstractProxyModel_CanFetchMore(ptr.Pointer(), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QAbstractProxyModel) Data(proxyIndex QModelIndex_ITF, role int) *QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractProxyModel::data")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQVariantFromPointer(C.QAbstractProxyModel_Data(ptr.Pointer(), PointerFromQModelIndex(proxyIndex), C.int(role)))
	}
	return nil
}

func (ptr *QAbstractProxyModel) DropMimeData(data QMimeData_ITF, action Qt__DropAction, row int, column int, parent QModelIndex_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractProxyModel::dropMimeData")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAbstractProxyModel_DropMimeData(ptr.Pointer(), PointerFromQMimeData(data), C.int(action), C.int(row), C.int(column), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QAbstractProxyModel) FetchMore(parent QModelIndex_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractProxyModel::fetchMore")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractProxyModel_FetchMore(ptr.Pointer(), PointerFromQModelIndex(parent))
	}
}

func (ptr *QAbstractProxyModel) Flags(index QModelIndex_ITF) Qt__ItemFlag {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractProxyModel::flags")
		}
	}()

	if ptr.Pointer() != nil {
		return Qt__ItemFlag(C.QAbstractProxyModel_Flags(ptr.Pointer(), PointerFromQModelIndex(index)))
	}
	return 0
}

func (ptr *QAbstractProxyModel) HasChildren(parent QModelIndex_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractProxyModel::hasChildren")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAbstractProxyModel_HasChildren(ptr.Pointer(), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QAbstractProxyModel) HeaderData(section int, orientation Qt__Orientation, role int) *QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractProxyModel::headerData")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQVariantFromPointer(C.QAbstractProxyModel_HeaderData(ptr.Pointer(), C.int(section), C.int(orientation), C.int(role)))
	}
	return nil
}

func (ptr *QAbstractProxyModel) MapFromSource(sourceIndex QModelIndex_ITF) *QModelIndex {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractProxyModel::mapFromSource")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QAbstractProxyModel_MapFromSource(ptr.Pointer(), PointerFromQModelIndex(sourceIndex)))
	}
	return nil
}

func (ptr *QAbstractProxyModel) MapToSource(proxyIndex QModelIndex_ITF) *QModelIndex {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractProxyModel::mapToSource")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QAbstractProxyModel_MapToSource(ptr.Pointer(), PointerFromQModelIndex(proxyIndex)))
	}
	return nil
}

func (ptr *QAbstractProxyModel) MimeTypes() []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractProxyModel::mimeTypes")
		}
	}()

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QAbstractProxyModel_MimeTypes(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QAbstractProxyModel) Revert() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractProxyModel::revert")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractProxyModel_Revert(ptr.Pointer())
	}
}

func (ptr *QAbstractProxyModel) SetData(index QModelIndex_ITF, value QVariant_ITF, role int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractProxyModel::setData")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAbstractProxyModel_SetData(ptr.Pointer(), PointerFromQModelIndex(index), PointerFromQVariant(value), C.int(role)) != 0
	}
	return false
}

func (ptr *QAbstractProxyModel) SetHeaderData(section int, orientation Qt__Orientation, value QVariant_ITF, role int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractProxyModel::setHeaderData")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAbstractProxyModel_SetHeaderData(ptr.Pointer(), C.int(section), C.int(orientation), PointerFromQVariant(value), C.int(role)) != 0
	}
	return false
}

func (ptr *QAbstractProxyModel) SetSourceModel(sourceModel QAbstractItemModel_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractProxyModel::setSourceModel")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractProxyModel_SetSourceModel(ptr.Pointer(), PointerFromQAbstractItemModel(sourceModel))
	}
}

func (ptr *QAbstractProxyModel) Sibling(row int, column int, idx QModelIndex_ITF) *QModelIndex {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractProxyModel::sibling")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QAbstractProxyModel_Sibling(ptr.Pointer(), C.int(row), C.int(column), PointerFromQModelIndex(idx)))
	}
	return nil
}

func (ptr *QAbstractProxyModel) Sort(column int, order Qt__SortOrder) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractProxyModel::sort")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractProxyModel_Sort(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QAbstractProxyModel) SourceModel() *QAbstractItemModel {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractProxyModel::sourceModel")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQAbstractItemModelFromPointer(C.QAbstractProxyModel_SourceModel(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractProxyModel) ConnectSourceModelChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractProxyModel::sourceModelChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractProxyModel_ConnectSourceModelChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sourceModelChanged", f)
	}
}

func (ptr *QAbstractProxyModel) DisconnectSourceModelChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractProxyModel::sourceModelChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractProxyModel_DisconnectSourceModelChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sourceModelChanged")
	}
}

//export callbackQAbstractProxyModelSourceModelChanged
func callbackQAbstractProxyModelSourceModelChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractProxyModel::sourceModelChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "sourceModelChanged").(func())()
}

func (ptr *QAbstractProxyModel) Submit() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractProxyModel::submit")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAbstractProxyModel_Submit(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractProxyModel) SupportedDragActions() Qt__DropAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractProxyModel::supportedDragActions")
		}
	}()

	if ptr.Pointer() != nil {
		return Qt__DropAction(C.QAbstractProxyModel_SupportedDragActions(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractProxyModel) SupportedDropActions() Qt__DropAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractProxyModel::supportedDropActions")
		}
	}()

	if ptr.Pointer() != nil {
		return Qt__DropAction(C.QAbstractProxyModel_SupportedDropActions(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractProxyModel) DestroyQAbstractProxyModel() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractProxyModel::~QAbstractProxyModel")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractProxyModel_DestroyQAbstractProxyModel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
