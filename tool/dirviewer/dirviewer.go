package main

import (
	"github.com/gizak/termui"
	"io/ioutil"
	"os"
	"path"
)

type DirViewManager struct {
	SelectedItem   int
	Items          []os.FileInfo
	CurrentDirPath string
}

func (m *DirViewManager) HandleUpKey() {
	m.SelectedItem--
	if m.SelectedItem < 0 {
		m.SelectedItem = 0
	}
}

func (m *DirViewManager) HandleDownKey() {
	m.SelectedItem++
	if m.SelectedItem == len(m.Items) {
		m.SelectedItem = len(m.Items) - 1
	}
}

func (m *DirViewManager) HandleEnterKey() (err error) {
	selectedItem := m.Items[m.SelectedItem]
	if selectedItem.IsDir() {
		m.CurrentDirPath = path.Join(m.CurrentDirPath, selectedItem.Name())
		err = m.getDirItems()
		m.SelectedItem = 0
	}
	return
}

func (m *DirViewManager) GetItems() (strs []string) {
	for i, item := range m.Items {
		if item.IsDir() {
			if i == m.SelectedItem {
				strs = append(strs, addColor(item.Name(), "fg-blue,bg-yellow"))
			} else {
				strs = append(strs, addColor(item.Name(), "fg-blue"))
			}
		} else {
			if i == m.SelectedItem {
				strs = append(strs, addColor(item.Name(), "bg-yellow"))
			} else {
				strs = append(strs, item.Name())
			}
		}
	}
	return
}

func (m *DirViewManager) getDirItems() (err error) {
	m.Items = []os.FileInfo{}

	parentDir, err := os.Open("..")
	if err != nil {
		return
	}

	parentDirInfo, err := parentDir.Stat()
	if err != nil {
		return
	}
	m.Items = append(m.Items, parentDirInfo)

	items, err := ioutil.ReadDir(m.CurrentDirPath)
	if err != nil {
		return
	}
	m.Items = append(m.Items, items...)
	return
}

func NewDirViewManager(dirpath string) (mgr *DirViewManager, err error) {
	mgr = &DirViewManager{
		SelectedItem:   0,
		CurrentDirPath: dirpath,
	}
	err = mgr.getDirItems()
	return
}

func addColor(str, color string) string {
	return "[" + str + "](" + color + ")"
}

func main() {
	err := termui.Init()
	if err != nil {
		panic(err)
	}
	defer termui.Close()

	mgr, err := NewDirViewManager(".")
	if err != nil {
		panic(err)
	}

	ls := termui.NewList()
	ls.Items = mgr.GetItems()
	ls.ItemFgColor = termui.ColorWhite
	ls.BorderLabel = "Directory View"
	ls.Height = 40
	ls.Width = 80
	ls.Y = 0

	termui.Render(ls)
	termui.Handle("/sys/kbd/q", func(termui.Event) {
		termui.StopLoop()
	})
	termui.Handle("/sys/kbd/<up>", func(termui.Event) {
		mgr.HandleUpKey()
		ls.Items = mgr.GetItems()
		termui.Render(ls)
	})
	termui.Handle("/sys/kbd/<down>", func(termui.Event) {
		mgr.HandleDownKey()
		ls.Items = mgr.GetItems()
		termui.Render(ls)
	})
	termui.Handle("/sys/kbd/<enter>", func(termui.Event) {
		err = mgr.HandleEnterKey()
		if err != nil {
			panic(err)
		}
		ls.Items = mgr.GetItems()
		termui.Render(ls)
	})

	termui.Loop()
}
