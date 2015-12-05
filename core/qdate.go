package core

//#include "core.h"
import "C"
import (
	"log"
	"unsafe"
)

type QDate struct {
	ptr unsafe.Pointer
}

type QDate_ITF interface {
	QDate_PTR() *QDate
}

func (p *QDate) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDate) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDate(ptr QDate_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDate_PTR().Pointer()
	}
	return nil
}

func NewQDateFromPointer(ptr unsafe.Pointer) *QDate {
	var n = new(QDate)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDate) QDate_PTR() *QDate {
	return ptr
}

//QDate::MonthNameType
type QDate__MonthNameType int64

const (
	QDate__DateFormat       = QDate__MonthNameType(0)
	QDate__StandaloneFormat = QDate__MonthNameType(1)
)

func QDate_IsLeapYear(year int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDate::isLeapYear")
		}
	}()

	return C.QDate_QDate_IsLeapYear(C.int(year)) != 0
}

func (ptr *QDate) ToString2(format Qt__DateFormat) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDate::toString")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QDate_ToString2(ptr.Pointer(), C.int(format)))
	}
	return ""
}

func NewQDate() *QDate {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDate::QDate")
		}
	}()

	return NewQDateFromPointer(C.QDate_NewQDate())
}

func NewQDate3(y int, m int, d int) *QDate {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDate::QDate")
		}
	}()

	return NewQDateFromPointer(C.QDate_NewQDate3(C.int(y), C.int(m), C.int(d)))
}

func (ptr *QDate) Day() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDate::day")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QDate_Day(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDate) DayOfWeek() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDate::dayOfWeek")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QDate_DayOfWeek(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDate) DayOfYear() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDate::dayOfYear")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QDate_DayOfYear(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDate) DaysInMonth() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDate::daysInMonth")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QDate_DaysInMonth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDate) DaysInYear() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDate::daysInYear")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QDate_DaysInYear(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDate) GetDate(year int, month int, day int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDate::getDate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDate_GetDate(ptr.Pointer(), C.int(year), C.int(month), C.int(day))
	}
}

func (ptr *QDate) IsNull() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDate::isNull")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QDate_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func QDate_IsValid2(year int, month int, day int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDate::isValid")
		}
	}()

	return C.QDate_QDate_IsValid2(C.int(year), C.int(month), C.int(day)) != 0
}

func (ptr *QDate) IsValid() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDate::isValid")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QDate_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func QDate_LongDayName(weekday int, ty QDate__MonthNameType) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDate::longDayName")
		}
	}()

	return C.GoString(C.QDate_QDate_LongDayName(C.int(weekday), C.int(ty)))
}

func QDate_LongMonthName(month int, ty QDate__MonthNameType) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDate::longMonthName")
		}
	}()

	return C.GoString(C.QDate_QDate_LongMonthName(C.int(month), C.int(ty)))
}

func (ptr *QDate) Month() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDate::month")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QDate_Month(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDate) SetDate(year int, month int, day int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDate::setDate")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QDate_SetDate(ptr.Pointer(), C.int(year), C.int(month), C.int(day)) != 0
	}
	return false
}

func QDate_ShortDayName(weekday int, ty QDate__MonthNameType) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDate::shortDayName")
		}
	}()

	return C.GoString(C.QDate_QDate_ShortDayName(C.int(weekday), C.int(ty)))
}

func QDate_ShortMonthName(month int, ty QDate__MonthNameType) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDate::shortMonthName")
		}
	}()

	return C.GoString(C.QDate_QDate_ShortMonthName(C.int(month), C.int(ty)))
}

func (ptr *QDate) ToString(format string) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDate::toString")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QDate_ToString(ptr.Pointer(), C.CString(format)))
	}
	return ""
}

func (ptr *QDate) WeekNumber(yearNumber int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDate::weekNumber")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QDate_WeekNumber(ptr.Pointer(), C.int(yearNumber)))
	}
	return 0
}

func (ptr *QDate) Year() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDate::year")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QDate_Year(ptr.Pointer()))
	}
	return 0
}
