package widgets

//#include "qmainwindow.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QMainWindow struct {
	QWidget
}

type QMainWindow_ITF interface {
	QWidget_ITF
	QMainWindow_PTR() *QMainWindow
}

func PointerFromQMainWindow(ptr QMainWindow_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMainWindow_PTR().Pointer()
	}
	return nil
}

func NewQMainWindowFromPointer(ptr unsafe.Pointer) *QMainWindow {
	var n = new(QMainWindow)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QMainWindow_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMainWindow) QMainWindow_PTR() *QMainWindow {
	return ptr
}

//QMainWindow::DockOption
type QMainWindow__DockOption int64

const (
	QMainWindow__AnimatedDocks    = QMainWindow__DockOption(0x01)
	QMainWindow__AllowNestedDocks = QMainWindow__DockOption(0x02)
	QMainWindow__AllowTabbedDocks = QMainWindow__DockOption(0x04)
	QMainWindow__ForceTabbedDocks = QMainWindow__DockOption(0x08)
	QMainWindow__VerticalTabs     = QMainWindow__DockOption(0x10)
)

func (ptr *QMainWindow) DockOptions() QMainWindow__DockOption {
	if ptr.Pointer() != nil {
		return QMainWindow__DockOption(C.QMainWindow_DockOptions(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMainWindow) DocumentMode() bool {
	if ptr.Pointer() != nil {
		return C.QMainWindow_DocumentMode(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMainWindow) IsAnimated() bool {
	if ptr.Pointer() != nil {
		return C.QMainWindow_IsAnimated(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMainWindow) IsDockNestingEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QMainWindow_IsDockNestingEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMainWindow) SetAnimated(enabled bool) {
	if ptr.Pointer() != nil {
		C.QMainWindow_SetAnimated(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QMainWindow) SetDockNestingEnabled(enabled bool) {
	if ptr.Pointer() != nil {
		C.QMainWindow_SetDockNestingEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QMainWindow) SetDockOptions(options QMainWindow__DockOption) {
	if ptr.Pointer() != nil {
		C.QMainWindow_SetDockOptions(ptr.Pointer(), C.int(options))
	}
}

func (ptr *QMainWindow) SetDocumentMode(enabled bool) {
	if ptr.Pointer() != nil {
		C.QMainWindow_SetDocumentMode(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QMainWindow) SetIconSize(iconSize core.QSize_ITF) {
	if ptr.Pointer() != nil {
		C.QMainWindow_SetIconSize(ptr.Pointer(), core.PointerFromQSize(iconSize))
	}
}

func (ptr *QMainWindow) SetTabShape(tabShape QTabWidget__TabShape) {
	if ptr.Pointer() != nil {
		C.QMainWindow_SetTabShape(ptr.Pointer(), C.int(tabShape))
	}
}

func (ptr *QMainWindow) SetToolButtonStyle(toolButtonStyle core.Qt__ToolButtonStyle) {
	if ptr.Pointer() != nil {
		C.QMainWindow_SetToolButtonStyle(ptr.Pointer(), C.int(toolButtonStyle))
	}
}

func (ptr *QMainWindow) SetUnifiedTitleAndToolBarOnMac(set bool) {
	if ptr.Pointer() != nil {
		C.QMainWindow_SetUnifiedTitleAndToolBarOnMac(ptr.Pointer(), C.int(qt.GoBoolToInt(set)))
	}
}

func (ptr *QMainWindow) SplitDockWidget(first QDockWidget_ITF, second QDockWidget_ITF, orientation core.Qt__Orientation) {
	if ptr.Pointer() != nil {
		C.QMainWindow_SplitDockWidget(ptr.Pointer(), PointerFromQDockWidget(first), PointerFromQDockWidget(second), C.int(orientation))
	}
}

func (ptr *QMainWindow) TabShape() QTabWidget__TabShape {
	if ptr.Pointer() != nil {
		return QTabWidget__TabShape(C.QMainWindow_TabShape(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMainWindow) TabifyDockWidget(first QDockWidget_ITF, second QDockWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QMainWindow_TabifyDockWidget(ptr.Pointer(), PointerFromQDockWidget(first), PointerFromQDockWidget(second))
	}
}

func (ptr *QMainWindow) ToolButtonStyle() core.Qt__ToolButtonStyle {
	if ptr.Pointer() != nil {
		return core.Qt__ToolButtonStyle(C.QMainWindow_ToolButtonStyle(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMainWindow) UnifiedTitleAndToolBarOnMac() bool {
	if ptr.Pointer() != nil {
		return C.QMainWindow_UnifiedTitleAndToolBarOnMac(ptr.Pointer()) != 0
	}
	return false
}

func NewQMainWindow(parent QWidget_ITF, flags core.Qt__WindowType) *QMainWindow {
	return NewQMainWindowFromPointer(C.QMainWindow_NewQMainWindow(PointerFromQWidget(parent), C.int(flags)))
}

func (ptr *QMainWindow) AddDockWidget(area core.Qt__DockWidgetArea, dockwidget QDockWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QMainWindow_AddDockWidget(ptr.Pointer(), C.int(area), PointerFromQDockWidget(dockwidget))
	}
}

func (ptr *QMainWindow) AddDockWidget2(area core.Qt__DockWidgetArea, dockwidget QDockWidget_ITF, orientation core.Qt__Orientation) {
	if ptr.Pointer() != nil {
		C.QMainWindow_AddDockWidget2(ptr.Pointer(), C.int(area), PointerFromQDockWidget(dockwidget), C.int(orientation))
	}
}

func (ptr *QMainWindow) AddToolBar3(title string) *QToolBar {
	if ptr.Pointer() != nil {
		return NewQToolBarFromPointer(C.QMainWindow_AddToolBar3(ptr.Pointer(), C.CString(title)))
	}
	return nil
}

func (ptr *QMainWindow) AddToolBar2(toolbar QToolBar_ITF) {
	if ptr.Pointer() != nil {
		C.QMainWindow_AddToolBar2(ptr.Pointer(), PointerFromQToolBar(toolbar))
	}
}

func (ptr *QMainWindow) AddToolBar(area core.Qt__ToolBarArea, toolbar QToolBar_ITF) {
	if ptr.Pointer() != nil {
		C.QMainWindow_AddToolBar(ptr.Pointer(), C.int(area), PointerFromQToolBar(toolbar))
	}
}

func (ptr *QMainWindow) AddToolBarBreak(area core.Qt__ToolBarArea) {
	if ptr.Pointer() != nil {
		C.QMainWindow_AddToolBarBreak(ptr.Pointer(), C.int(area))
	}
}

func (ptr *QMainWindow) CentralWidget() *QWidget {
	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QMainWindow_CentralWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMainWindow) Corner(corner core.Qt__Corner) core.Qt__DockWidgetArea {
	if ptr.Pointer() != nil {
		return core.Qt__DockWidgetArea(C.QMainWindow_Corner(ptr.Pointer(), C.int(corner)))
	}
	return 0
}

func (ptr *QMainWindow) CreatePopupMenu() *QMenu {
	if ptr.Pointer() != nil {
		return NewQMenuFromPointer(C.QMainWindow_CreatePopupMenu(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMainWindow) DockWidgetArea(dockwidget QDockWidget_ITF) core.Qt__DockWidgetArea {
	if ptr.Pointer() != nil {
		return core.Qt__DockWidgetArea(C.QMainWindow_DockWidgetArea(ptr.Pointer(), PointerFromQDockWidget(dockwidget)))
	}
	return 0
}

func (ptr *QMainWindow) InsertToolBar(before QToolBar_ITF, toolbar QToolBar_ITF) {
	if ptr.Pointer() != nil {
		C.QMainWindow_InsertToolBar(ptr.Pointer(), PointerFromQToolBar(before), PointerFromQToolBar(toolbar))
	}
}

func (ptr *QMainWindow) InsertToolBarBreak(before QToolBar_ITF) {
	if ptr.Pointer() != nil {
		C.QMainWindow_InsertToolBarBreak(ptr.Pointer(), PointerFromQToolBar(before))
	}
}

func (ptr *QMainWindow) MenuBar() *QMenuBar {
	if ptr.Pointer() != nil {
		return NewQMenuBarFromPointer(C.QMainWindow_MenuBar(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMainWindow) MenuWidget() *QWidget {
	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QMainWindow_MenuWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMainWindow) RemoveDockWidget(dockwidget QDockWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QMainWindow_RemoveDockWidget(ptr.Pointer(), PointerFromQDockWidget(dockwidget))
	}
}

func (ptr *QMainWindow) RemoveToolBar(toolbar QToolBar_ITF) {
	if ptr.Pointer() != nil {
		C.QMainWindow_RemoveToolBar(ptr.Pointer(), PointerFromQToolBar(toolbar))
	}
}

func (ptr *QMainWindow) RemoveToolBarBreak(before QToolBar_ITF) {
	if ptr.Pointer() != nil {
		C.QMainWindow_RemoveToolBarBreak(ptr.Pointer(), PointerFromQToolBar(before))
	}
}

func (ptr *QMainWindow) RestoreDockWidget(dockwidget QDockWidget_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QMainWindow_RestoreDockWidget(ptr.Pointer(), PointerFromQDockWidget(dockwidget)) != 0
	}
	return false
}

func (ptr *QMainWindow) RestoreState(state core.QByteArray_ITF, version int) bool {
	if ptr.Pointer() != nil {
		return C.QMainWindow_RestoreState(ptr.Pointer(), core.PointerFromQByteArray(state), C.int(version)) != 0
	}
	return false
}

func (ptr *QMainWindow) SaveState(version int) *core.QByteArray {
	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QMainWindow_SaveState(ptr.Pointer(), C.int(version)))
	}
	return nil
}

func (ptr *QMainWindow) SetCentralWidget(widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QMainWindow_SetCentralWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QMainWindow) SetCorner(corner core.Qt__Corner, area core.Qt__DockWidgetArea) {
	if ptr.Pointer() != nil {
		C.QMainWindow_SetCorner(ptr.Pointer(), C.int(corner), C.int(area))
	}
}

func (ptr *QMainWindow) SetMenuBar(menuBar QMenuBar_ITF) {
	if ptr.Pointer() != nil {
		C.QMainWindow_SetMenuBar(ptr.Pointer(), PointerFromQMenuBar(menuBar))
	}
}

func (ptr *QMainWindow) SetMenuWidget(menuBar QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QMainWindow_SetMenuWidget(ptr.Pointer(), PointerFromQWidget(menuBar))
	}
}

func (ptr *QMainWindow) SetStatusBar(statusbar QStatusBar_ITF) {
	if ptr.Pointer() != nil {
		C.QMainWindow_SetStatusBar(ptr.Pointer(), PointerFromQStatusBar(statusbar))
	}
}

func (ptr *QMainWindow) SetTabPosition(areas core.Qt__DockWidgetArea, tabPosition QTabWidget__TabPosition) {
	if ptr.Pointer() != nil {
		C.QMainWindow_SetTabPosition(ptr.Pointer(), C.int(areas), C.int(tabPosition))
	}
}

func (ptr *QMainWindow) StatusBar() *QStatusBar {
	if ptr.Pointer() != nil {
		return NewQStatusBarFromPointer(C.QMainWindow_StatusBar(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMainWindow) TabPosition(area core.Qt__DockWidgetArea) QTabWidget__TabPosition {
	if ptr.Pointer() != nil {
		return QTabWidget__TabPosition(C.QMainWindow_TabPosition(ptr.Pointer(), C.int(area)))
	}
	return 0
}

func (ptr *QMainWindow) TakeCentralWidget() *QWidget {
	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QMainWindow_TakeCentralWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMainWindow) ToolBarArea(toolbar QToolBar_ITF) core.Qt__ToolBarArea {
	if ptr.Pointer() != nil {
		return core.Qt__ToolBarArea(C.QMainWindow_ToolBarArea(ptr.Pointer(), PointerFromQToolBar(toolbar)))
	}
	return 0
}

func (ptr *QMainWindow) ToolBarBreak(toolbar QToolBar_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QMainWindow_ToolBarBreak(ptr.Pointer(), PointerFromQToolBar(toolbar)) != 0
	}
	return false
}

func (ptr *QMainWindow) ConnectToolButtonStyleChanged(f func(toolButtonStyle core.Qt__ToolButtonStyle)) {
	if ptr.Pointer() != nil {
		C.QMainWindow_ConnectToolButtonStyleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "toolButtonStyleChanged", f)
	}
}

func (ptr *QMainWindow) DisconnectToolButtonStyleChanged() {
	if ptr.Pointer() != nil {
		C.QMainWindow_DisconnectToolButtonStyleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "toolButtonStyleChanged")
	}
}

//export callbackQMainWindowToolButtonStyleChanged
func callbackQMainWindowToolButtonStyleChanged(ptrName *C.char, toolButtonStyle C.int) {
	qt.GetSignal(C.GoString(ptrName), "toolButtonStyleChanged").(func(core.Qt__ToolButtonStyle))(core.Qt__ToolButtonStyle(toolButtonStyle))
}

func (ptr *QMainWindow) DestroyQMainWindow() {
	if ptr.Pointer() != nil {
		C.QMainWindow_DestroyQMainWindow(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
