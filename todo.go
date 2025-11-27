package main

import (
	"encoding/json"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type TodoItem struct {
	Text     string `json:"text"`
	Done     bool   `json:"done"`
	Priority string `json:"priority"`
}

type TodoTab struct {
	todoList    *widget.List
	todos       []TodoItem
	entry       *widget.Entry
	priority    *widget.Select
	currentIndex int
}

func NewTodoTab() fyne.CanvasObject {
	todoTab := &TodoTab{
		todos:        []TodoItem{},
		currentIndex: -1,
	}

	// ورودی کار جدید
	inputLabel := widget.NewLabel("New Task:")
	todoTab.entry = widget.NewEntry()
	todoTab.entry.SetPlaceHolder("Enter new task...")
	todoTab.entry.OnSubmitted = func(text string) {
		todoTab.addTodo()
	}

	// انتخاب اولویت
	priorityLabel := widget.NewLabel("Priority:")
	todoTab.priority = widget.NewSelect([]string{"High", "Medium", "Low"}, func(s string) {})
	todoTab.priority.SetSelected("Medium")

	// دکمه‌ها
	addBtn := widget.NewButton("Add", func() {
		todoTab.addTodo()
	})
	addBtn.Importance = widget.HighImportance

	clearBtn := widget.NewButton("Clear Completed", func() {
		todoTab.clearCompleted()
	})

	buttonBox := container.NewHBox(addBtn, clearBtn)

	// لیست کارها
	todoTab.todoList = widget.NewList(
		func() int {
			return len(todoTab.todos)
		},
		func() fyne.CanvasObject {
			check := widget.NewCheck("", nil)
			label := widget.NewLabel("")
			priorityLabel := widget.NewLabel("")
			return container.NewHBox(check, label, priorityLabel)
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			if id < len(todoTab.todos) {
				box := obj.(*fyne.Container)
				objects := box.Objects
				if len(objects) >= 3 {
					check := objects[0].(*widget.Check)
					label := objects[1].(*widget.Label)
					priorityLabel := objects[2].(*widget.Label)

					todo := todoTab.todos[id]
					label.SetText(todo.Text)
					check.SetChecked(todo.Done)
					if todo.Done {
						label.Importance = widget.LowImportance
					} else {
						label.Importance = widget.MediumImportance
					}

					priorityLabel.SetText("[" + todo.Priority + "]")
					switch todo.Priority {
					case "High", "بالا":
						priorityLabel.Importance = widget.DangerImportance
					case "Medium", "متوسط":
						priorityLabel.Importance = widget.WarningImportance
					case "Low", "پایین":
						priorityLabel.Importance = widget.LowImportance
					}

					check.OnChanged = func(done bool) {
						todoTab.todos[id].Done = done
						todoTab.todoList.Refresh()
						todoTab.saveTodos()
					}
				}
			}
		})

	todoTab.todoList.OnSelected = func(id widget.ListItemID) {
		if id < len(todoTab.todos) {
			todoTab.currentIndex = id
			todoTab.entry.SetText(todoTab.todos[id].Text)
			todoTab.priority.SetSelected(todoTab.todos[id].Priority)
		}
	}

	// بارگذاری کارها
	todoTab.loadTodos()

	// بهبود چیدمان با padding و spacing بهتر
	inputSection := container.NewVBox(
		container.NewPadded(container.NewVBox(inputLabel, todoTab.entry)),
		container.NewPadded(container.NewVBox(priorityLabel, todoTab.priority)),
		container.NewPadded(buttonBox),
	)

	listSection := container.NewBorder(
		container.NewPadded(widget.NewLabel("Tasks List")),
		nil,
		nil,
		nil,
		container.NewPadded(todoTab.todoList),
	)

	mainContainer := container.NewBorder(
		container.NewPadded(inputSection),
		nil,
		nil,
		nil,
		listSection,
	)

	return mainContainer
}

func (tt *TodoTab) addTodo() {
	text := tt.entry.Text
	if text == "" {
		return
	}

	priority := tt.priority.Selected
	if priority == "" {
		priority = "Medium"
	}

	if tt.currentIndex >= 0 && tt.currentIndex < len(tt.todos) {
		tt.todos[tt.currentIndex] = TodoItem{
			Text:     text,
			Done:     tt.todos[tt.currentIndex].Done,
			Priority: priority,
		}
	} else {
		tt.todos = append(tt.todos, TodoItem{
			Text:     text,
			Done:     false,
			Priority: priority,
		})
	}

	tt.entry.SetText("")
	tt.currentIndex = -1
	tt.priority.SetSelected("Medium")
	tt.todoList.Refresh()
	tt.saveTodos()
}

func (tt *TodoTab) clearCompleted() {
	var newTodos []TodoItem
	for _, todo := range tt.todos {
		if !todo.Done {
			newTodos = append(newTodos, todo)
		}
	}
	tt.todos = newTodos
	tt.todoList.Refresh()
	tt.saveTodos()
}

func (tt *TodoTab) saveTodos() {
	homeDir, _ := os.UserHomeDir()
	todoDir := filepath.Join(homeDir, ".notepad-app")
	os.MkdirAll(todoDir, 0755)
	todoFile := filepath.Join(todoDir, "todos.json")

	// استفاده از MarshalIndent برای خوانایی بهتر و UTF-8
	data, err := json.MarshalIndent(tt.todos, "", "  ")
	if err != nil {
		return
	}

	// ذخیره با UTF-8 encoding
	os.WriteFile(todoFile, data, 0644)
}

func (tt *TodoTab) loadTodos() {
	homeDir, _ := os.UserHomeDir()
	todoFile := filepath.Join(homeDir, ".notepad-app", "todos.json")

	data, err := os.ReadFile(todoFile)
	if err != nil {
		return
	}

	json.Unmarshal(data, &tt.todos)
	tt.todoList.Refresh()
}

