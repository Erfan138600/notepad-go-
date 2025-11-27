package main

import (
	"encoding/json"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Note struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type NotesTab struct {
	notesList    *widget.List
	notes        []Note
	titleEntry   *widget.Entry
	contentEntry *widget.Entry
	currentIndex int
}

func NewNotesTab() fyne.CanvasObject {
	notesTab := &NotesTab{
		notes:        []Note{},
		currentIndex: -1,
	}

	// ورودی عنوان
	titleLabel := widget.NewLabel("Title:")
	notesTab.titleEntry = widget.NewEntry()
	notesTab.titleEntry.SetPlaceHolder("Enter note title...")

	// ورودی محتوا
	contentLabel := widget.NewLabel("Content:")
	notesTab.contentEntry = widget.NewMultiLineEntry()
	notesTab.contentEntry.SetPlaceHolder("Write your note here...")
	notesTab.contentEntry.Wrapping = fyne.TextWrapWord

	// دکمه‌های عملیات
	saveBtn := widget.NewButton("Save", func() {
		notesTab.saveNote()
	})
	saveBtn.Importance = widget.HighImportance

	newBtn := widget.NewButton("New", func() {
		notesTab.newNote()
	})

	deleteBtn := widget.NewButton("Delete", func() {
		notesTab.deleteNote()
	})

	buttonBox := container.NewHBox(newBtn, saveBtn, deleteBtn)

	// لیست یادداشت‌ها
	notesTab.notesList = widget.NewList(
		func() int {
			return len(notesTab.notes)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("عنوان")
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			if id < len(notesTab.notes) {
				obj.(*widget.Label).SetText(notesTab.notes[id].Title)
			}
		})

	notesTab.notesList.OnSelected = func(id widget.ListItemID) {
		if id < len(notesTab.notes) {
			notesTab.currentIndex = id
			notesTab.titleEntry.SetText(notesTab.notes[id].Title)
			notesTab.contentEntry.SetText(notesTab.notes[id].Content)
		}
	}

	// بارگذاری یادداشت‌ها
	notesTab.loadNotes()

	// بهبود چیدمان با padding
	titleContainer := container.NewVBox(titleLabel, notesTab.titleEntry)
	contentContainer := container.NewVBox(contentLabel, notesTab.contentEntry)
	
	formContainer := container.NewBorder(
		nil,
		container.NewPadded(buttonBox),
		nil,
		nil,
		container.NewVBox(
			container.NewPadded(titleContainer),
			container.NewPadded(contentContainer),
		),
	)

	// لیست با padding
	listContainer := container.NewPadded(
		container.NewBorder(
			widget.NewLabel("Notes List"),
			nil,
			nil,
			nil,
			notesTab.notesList,
		),
	)

	split := container.NewHSplit(listContainer, formContainer)
	split.SetOffset(0.25)

	return split
}

func (nt *NotesTab) saveNote() {
	title := nt.titleEntry.Text
	content := nt.contentEntry.Text

	if title == "" {
		return
	}

	if nt.currentIndex >= 0 && nt.currentIndex < len(nt.notes) {
		nt.notes[nt.currentIndex] = Note{Title: title, Content: content}
	} else {
		nt.notes = append(nt.notes, Note{Title: title, Content: content})
		nt.currentIndex = len(nt.notes) - 1
	}

	nt.notesList.Refresh()
	nt.saveNotes()
}

func (nt *NotesTab) newNote() {
	nt.currentIndex = -1
	nt.titleEntry.SetText("")
	nt.contentEntry.SetText("")
	nt.notesList.UnselectAll()
}

func (nt *NotesTab) deleteNote() {
	if nt.currentIndex >= 0 && nt.currentIndex < len(nt.notes) {
		nt.notes = append(nt.notes[:nt.currentIndex], nt.notes[nt.currentIndex+1:]...)
		nt.currentIndex = -1
		nt.titleEntry.SetText("")
		nt.contentEntry.SetText("")
		nt.notesList.Refresh()
		nt.saveNotes()
	}
}

func (nt *NotesTab) saveNotes() {
	homeDir, _ := os.UserHomeDir()
	notesDir := filepath.Join(homeDir, ".notepad-app")
	os.MkdirAll(notesDir, 0755)
	notesFile := filepath.Join(notesDir, "notes.json")

	// استفاده از MarshalIndent برای خوانایی بهتر و UTF-8
	data, err := json.MarshalIndent(nt.notes, "", "  ")
	if err != nil {
		return
	}

	// ذخیره با UTF-8 encoding
	os.WriteFile(notesFile, data, 0644)
}

func (nt *NotesTab) loadNotes() {
	homeDir, _ := os.UserHomeDir()
	notesFile := filepath.Join(homeDir, ".notepad-app", "notes.json")

	data, err := os.ReadFile(notesFile)
	if err != nil {
		return
	}

	json.Unmarshal(data, &nt.notes)
	nt.notesList.Refresh()
}

