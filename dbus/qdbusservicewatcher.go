package dbus

//#include "dbus.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"strings"
	"unsafe"
)

type QDBusServiceWatcher struct {
	core.QObject
}

type QDBusServiceWatcher_ITF interface {
	core.QObject_ITF
	QDBusServiceWatcher_PTR() *QDBusServiceWatcher
}

func PointerFromQDBusServiceWatcher(ptr QDBusServiceWatcher_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusServiceWatcher_PTR().Pointer()
	}
	return nil
}

func NewQDBusServiceWatcherFromPointer(ptr unsafe.Pointer) *QDBusServiceWatcher {
	var n = new(QDBusServiceWatcher)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QDBusServiceWatcher_") {
		n.SetObjectName("QDBusServiceWatcher_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QDBusServiceWatcher) QDBusServiceWatcher_PTR() *QDBusServiceWatcher {
	return ptr
}

//QDBusServiceWatcher::WatchModeFlag
type QDBusServiceWatcher__WatchModeFlag int64

const (
	QDBusServiceWatcher__WatchForRegistration   = QDBusServiceWatcher__WatchModeFlag(0x01)
	QDBusServiceWatcher__WatchForUnregistration = QDBusServiceWatcher__WatchModeFlag(0x02)
	QDBusServiceWatcher__WatchForOwnerChange    = QDBusServiceWatcher__WatchModeFlag(0x03)
)

func (ptr *QDBusServiceWatcher) SetWatchMode(mode QDBusServiceWatcher__WatchModeFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusServiceWatcher::setWatchMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_SetWatchMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QDBusServiceWatcher) WatchMode() QDBusServiceWatcher__WatchModeFlag {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusServiceWatcher::watchMode")
		}
	}()

	if ptr.Pointer() != nil {
		return QDBusServiceWatcher__WatchModeFlag(C.QDBusServiceWatcher_WatchMode(ptr.Pointer()))
	}
	return 0
}

func NewQDBusServiceWatcher(parent core.QObject_ITF) *QDBusServiceWatcher {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusServiceWatcher::QDBusServiceWatcher")
		}
	}()

	return NewQDBusServiceWatcherFromPointer(C.QDBusServiceWatcher_NewQDBusServiceWatcher(core.PointerFromQObject(parent)))
}

func NewQDBusServiceWatcher2(service string, connection QDBusConnection_ITF, watchMode QDBusServiceWatcher__WatchModeFlag, parent core.QObject_ITF) *QDBusServiceWatcher {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusServiceWatcher::QDBusServiceWatcher")
		}
	}()

	return NewQDBusServiceWatcherFromPointer(C.QDBusServiceWatcher_NewQDBusServiceWatcher2(C.CString(service), PointerFromQDBusConnection(connection), C.int(watchMode), core.PointerFromQObject(parent)))
}

func (ptr *QDBusServiceWatcher) AddWatchedService(newService string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusServiceWatcher::addWatchedService")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_AddWatchedService(ptr.Pointer(), C.CString(newService))
	}
}

func (ptr *QDBusServiceWatcher) RemoveWatchedService(service string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusServiceWatcher::removeWatchedService")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QDBusServiceWatcher_RemoveWatchedService(ptr.Pointer(), C.CString(service)) != 0
	}
	return false
}

func (ptr *QDBusServiceWatcher) ConnectServiceOwnerChanged(f func(serviceName string, oldOwner string, newOwner string)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusServiceWatcher::serviceOwnerChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_ConnectServiceOwnerChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "serviceOwnerChanged", f)
	}
}

func (ptr *QDBusServiceWatcher) DisconnectServiceOwnerChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusServiceWatcher::serviceOwnerChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_DisconnectServiceOwnerChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "serviceOwnerChanged")
	}
}

//export callbackQDBusServiceWatcherServiceOwnerChanged
func callbackQDBusServiceWatcherServiceOwnerChanged(ptrName *C.char, serviceName *C.char, oldOwner *C.char, newOwner *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusServiceWatcher::serviceOwnerChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "serviceOwnerChanged").(func(string, string, string))(C.GoString(serviceName), C.GoString(oldOwner), C.GoString(newOwner))
}

func (ptr *QDBusServiceWatcher) ConnectServiceRegistered(f func(serviceName string)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusServiceWatcher::serviceRegistered")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_ConnectServiceRegistered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "serviceRegistered", f)
	}
}

func (ptr *QDBusServiceWatcher) DisconnectServiceRegistered() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusServiceWatcher::serviceRegistered")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_DisconnectServiceRegistered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "serviceRegistered")
	}
}

//export callbackQDBusServiceWatcherServiceRegistered
func callbackQDBusServiceWatcherServiceRegistered(ptrName *C.char, serviceName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusServiceWatcher::serviceRegistered")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "serviceRegistered").(func(string))(C.GoString(serviceName))
}

func (ptr *QDBusServiceWatcher) ConnectServiceUnregistered(f func(serviceName string)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusServiceWatcher::serviceUnregistered")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_ConnectServiceUnregistered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "serviceUnregistered", f)
	}
}

func (ptr *QDBusServiceWatcher) DisconnectServiceUnregistered() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusServiceWatcher::serviceUnregistered")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_DisconnectServiceUnregistered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "serviceUnregistered")
	}
}

//export callbackQDBusServiceWatcherServiceUnregistered
func callbackQDBusServiceWatcherServiceUnregistered(ptrName *C.char, serviceName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusServiceWatcher::serviceUnregistered")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "serviceUnregistered").(func(string))(C.GoString(serviceName))
}

func (ptr *QDBusServiceWatcher) SetConnection(connection QDBusConnection_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusServiceWatcher::setConnection")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_SetConnection(ptr.Pointer(), PointerFromQDBusConnection(connection))
	}
}

func (ptr *QDBusServiceWatcher) SetWatchedServices(services []string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusServiceWatcher::setWatchedServices")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_SetWatchedServices(ptr.Pointer(), C.CString(strings.Join(services, ",,,")))
	}
}

func (ptr *QDBusServiceWatcher) WatchedServices() []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusServiceWatcher::watchedServices")
		}
	}()

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QDBusServiceWatcher_WatchedServices(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QDBusServiceWatcher) DestroyQDBusServiceWatcher() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusServiceWatcher::~QDBusServiceWatcher")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_DestroyQDBusServiceWatcher(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
