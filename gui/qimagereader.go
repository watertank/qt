package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"strings"
	"unsafe"
)

type QImageReader struct {
	ptr unsafe.Pointer
}

type QImageReader_ITF interface {
	QImageReader_PTR() *QImageReader
}

func (p *QImageReader) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QImageReader) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQImageReader(ptr QImageReader_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QImageReader_PTR().Pointer()
	}
	return nil
}

func NewQImageReaderFromPointer(ptr unsafe.Pointer) *QImageReader {
	var n = new(QImageReader)
	n.SetPointer(ptr)
	return n
}

func (ptr *QImageReader) QImageReader_PTR() *QImageReader {
	return ptr
}

//QImageReader::ImageReaderError
type QImageReader__ImageReaderError int64

const (
	QImageReader__UnknownError           = QImageReader__ImageReaderError(0)
	QImageReader__FileNotFoundError      = QImageReader__ImageReaderError(1)
	QImageReader__DeviceError            = QImageReader__ImageReaderError(2)
	QImageReader__UnsupportedFormatError = QImageReader__ImageReaderError(3)
	QImageReader__InvalidDataError       = QImageReader__ImageReaderError(4)
)

func NewQImageReader() *QImageReader {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageReader::QImageReader")
		}
	}()

	return NewQImageReaderFromPointer(C.QImageReader_NewQImageReader())
}

func NewQImageReader2(device core.QIODevice_ITF, format core.QByteArray_ITF) *QImageReader {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageReader::QImageReader")
		}
	}()

	return NewQImageReaderFromPointer(C.QImageReader_NewQImageReader2(core.PointerFromQIODevice(device), core.PointerFromQByteArray(format)))
}

func NewQImageReader3(fileName string, format core.QByteArray_ITF) *QImageReader {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageReader::QImageReader")
		}
	}()

	return NewQImageReaderFromPointer(C.QImageReader_NewQImageReader3(C.CString(fileName), core.PointerFromQByteArray(format)))
}

func (ptr *QImageReader) AutoDetectImageFormat() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageReader::autoDetectImageFormat")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QImageReader_AutoDetectImageFormat(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QImageReader) AutoTransform() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageReader::autoTransform")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QImageReader_AutoTransform(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QImageReader) BackgroundColor() *QColor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageReader::backgroundColor")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQColorFromPointer(C.QImageReader_BackgroundColor(ptr.Pointer()))
	}
	return nil
}

func (ptr *QImageReader) CanRead() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageReader::canRead")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QImageReader_CanRead(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QImageReader) CurrentImageNumber() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageReader::currentImageNumber")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QImageReader_CurrentImageNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QImageReader) DecideFormatFromContent() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageReader::decideFormatFromContent")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QImageReader_DecideFormatFromContent(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QImageReader) Device() *core.QIODevice {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageReader::device")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QImageReader_Device(ptr.Pointer()))
	}
	return nil
}

func (ptr *QImageReader) Error() QImageReader__ImageReaderError {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageReader::error")
		}
	}()

	if ptr.Pointer() != nil {
		return QImageReader__ImageReaderError(C.QImageReader_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QImageReader) ErrorString() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageReader::errorString")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QImageReader_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QImageReader) FileName() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageReader::fileName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QImageReader_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QImageReader) Format() *core.QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageReader::format")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QImageReader_Format(ptr.Pointer()))
	}
	return nil
}

func (ptr *QImageReader) ImageCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageReader::imageCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QImageReader_ImageCount(ptr.Pointer()))
	}
	return 0
}

func QImageReader_ImageFormat3(device core.QIODevice_ITF) *core.QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageReader::imageFormat")
		}
	}()

	return core.NewQByteArrayFromPointer(C.QImageReader_QImageReader_ImageFormat3(core.PointerFromQIODevice(device)))
}

func QImageReader_ImageFormat2(fileName string) *core.QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageReader::imageFormat")
		}
	}()

	return core.NewQByteArrayFromPointer(C.QImageReader_QImageReader_ImageFormat2(C.CString(fileName)))
}

func (ptr *QImageReader) ImageFormat() QImage__Format {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageReader::imageFormat")
		}
	}()

	if ptr.Pointer() != nil {
		return QImage__Format(C.QImageReader_ImageFormat(ptr.Pointer()))
	}
	return 0
}

func (ptr *QImageReader) JumpToImage(imageNumber int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageReader::jumpToImage")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QImageReader_JumpToImage(ptr.Pointer(), C.int(imageNumber)) != 0
	}
	return false
}

func (ptr *QImageReader) JumpToNextImage() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageReader::jumpToNextImage")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QImageReader_JumpToNextImage(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QImageReader) LoopCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageReader::loopCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QImageReader_LoopCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QImageReader) NextImageDelay() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageReader::nextImageDelay")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QImageReader_NextImageDelay(ptr.Pointer()))
	}
	return 0
}

func (ptr *QImageReader) Quality() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageReader::quality")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QImageReader_Quality(ptr.Pointer()))
	}
	return 0
}

func (ptr *QImageReader) Read2(image QImage_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageReader::read")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QImageReader_Read2(ptr.Pointer(), PointerFromQImage(image)) != 0
	}
	return false
}

func (ptr *QImageReader) SetAutoDetectImageFormat(enabled bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageReader::setAutoDetectImageFormat")
		}
	}()

	if ptr.Pointer() != nil {
		C.QImageReader_SetAutoDetectImageFormat(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QImageReader) SetAutoTransform(enabled bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageReader::setAutoTransform")
		}
	}()

	if ptr.Pointer() != nil {
		C.QImageReader_SetAutoTransform(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QImageReader) SetBackgroundColor(color QColor_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageReader::setBackgroundColor")
		}
	}()

	if ptr.Pointer() != nil {
		C.QImageReader_SetBackgroundColor(ptr.Pointer(), PointerFromQColor(color))
	}
}

func (ptr *QImageReader) SetClipRect(rect core.QRect_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageReader::setClipRect")
		}
	}()

	if ptr.Pointer() != nil {
		C.QImageReader_SetClipRect(ptr.Pointer(), core.PointerFromQRect(rect))
	}
}

func (ptr *QImageReader) SetDecideFormatFromContent(ignored bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageReader::setDecideFormatFromContent")
		}
	}()

	if ptr.Pointer() != nil {
		C.QImageReader_SetDecideFormatFromContent(ptr.Pointer(), C.int(qt.GoBoolToInt(ignored)))
	}
}

func (ptr *QImageReader) SetDevice(device core.QIODevice_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageReader::setDevice")
		}
	}()

	if ptr.Pointer() != nil {
		C.QImageReader_SetDevice(ptr.Pointer(), core.PointerFromQIODevice(device))
	}
}

func (ptr *QImageReader) SetFileName(fileName string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageReader::setFileName")
		}
	}()

	if ptr.Pointer() != nil {
		C.QImageReader_SetFileName(ptr.Pointer(), C.CString(fileName))
	}
}

func (ptr *QImageReader) SetFormat(format core.QByteArray_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageReader::setFormat")
		}
	}()

	if ptr.Pointer() != nil {
		C.QImageReader_SetFormat(ptr.Pointer(), core.PointerFromQByteArray(format))
	}
}

func (ptr *QImageReader) SetQuality(quality int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageReader::setQuality")
		}
	}()

	if ptr.Pointer() != nil {
		C.QImageReader_SetQuality(ptr.Pointer(), C.int(quality))
	}
}

func (ptr *QImageReader) SetScaledClipRect(rect core.QRect_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageReader::setScaledClipRect")
		}
	}()

	if ptr.Pointer() != nil {
		C.QImageReader_SetScaledClipRect(ptr.Pointer(), core.PointerFromQRect(rect))
	}
}

func (ptr *QImageReader) SetScaledSize(size core.QSize_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageReader::setScaledSize")
		}
	}()

	if ptr.Pointer() != nil {
		C.QImageReader_SetScaledSize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QImageReader) SubType() *core.QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageReader::subType")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QImageReader_SubType(ptr.Pointer()))
	}
	return nil
}

func (ptr *QImageReader) SupportsAnimation() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageReader::supportsAnimation")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QImageReader_SupportsAnimation(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QImageReader) SupportsOption(option QImageIOHandler__ImageOption) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageReader::supportsOption")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QImageReader_SupportsOption(ptr.Pointer(), C.int(option)) != 0
	}
	return false
}

func (ptr *QImageReader) Text(key string) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageReader::text")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QImageReader_Text(ptr.Pointer(), C.CString(key)))
	}
	return ""
}

func (ptr *QImageReader) TextKeys() []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageReader::textKeys")
		}
	}()

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QImageReader_TextKeys(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QImageReader) Transformation() QImageIOHandler__Transformation {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageReader::transformation")
		}
	}()

	if ptr.Pointer() != nil {
		return QImageIOHandler__Transformation(C.QImageReader_Transformation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QImageReader) DestroyQImageReader() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageReader::~QImageReader")
		}
	}()

	if ptr.Pointer() != nil {
		C.QImageReader_DestroyQImageReader(ptr.Pointer())
	}
}
