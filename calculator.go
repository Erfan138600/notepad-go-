package main

import (
	"fmt"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type CalculatorTab struct {
	display    *widget.Entry
	expression string
	result     float64
	isNewInput bool
}

func NewCalculatorTab() fyne.CanvasObject {
	calc := &CalculatorTab{
		isNewInput: true,
	}

	calc.display = widget.NewEntry()
	calc.display.SetText("0")
	calc.display.Disable()
	calc.display.TextStyle = fyne.TextStyle{Monospace: true}

	// دکمه‌های ماشین حساب
	buttons := calc.createButtons()

	// ردیف اول
	row1 := container.NewGridWithColumns(4,
		buttons["clear"], buttons["±"], buttons["%"], buttons["÷"],
	)
	
	// ردیف دوم
	row2 := container.NewGridWithColumns(4,
		buttons["7"], buttons["8"], buttons["9"], buttons["×"],
	)
	
	// ردیف سوم
	row3 := container.NewGridWithColumns(4,
		buttons["4"], buttons["5"], buttons["6"], buttons["-"],
	)
	
	// ردیف چهارم
	row4 := container.NewGridWithColumns(4,
		buttons["1"], buttons["2"], buttons["3"], buttons["+"],
	)
	
	// ردیف پنجم - دکمه صفر بزرگتر
	zeroRowLeft := container.NewGridWithColumns(2, buttons["0"], buttons["."])
	zeroRow := container.NewGridWithColumns(3, zeroRowLeft, buttons["="])

	// بهبود layout با padding
	calcContainer := container.NewVBox(
		container.NewPadded(calc.display),
		container.NewPadded(row1),
		container.NewPadded(row2),
		container.NewPadded(row3),
		container.NewPadded(row4),
		container.NewPadded(zeroRow),
	)

	mainContainer := container.NewPadded(
		container.NewBorder(
			container.NewPadded(widget.NewLabel("Calculator")),
			nil,
			nil,
			nil,
			calcContainer,
		),
	)

	return mainContainer
}

func (c *CalculatorTab) createButtons() map[string]*widget.Button {
	buttons := make(map[string]*widget.Button)

	// اعداد
	for i := 0; i <= 9; i++ {
		num := strconv.Itoa(i)
		buttons[num] = widget.NewButton(num, func(n string) func() {
			return func() {
				c.appendNumber(n)
			}
		}(num))
	}

	// عملگرها
	operators := map[string]func(){
		"+": func() { c.setOperator("+") },
		"-": func() { c.setOperator("-") },
		"×": func() { c.setOperator("×") },
		"÷": func() { c.setOperator("÷") },
		"=": func() { c.calculate() },
		".": func() { c.appendDecimal() },
		"%": func() { c.percentage() },
		"±": func() { c.toggleSign() },
		"clear": func() { c.clear() },
	}

	for op, fn := range operators {
		btn := widget.NewButton(op, fn)
		if op == "=" {
			btn.Importance = widget.HighImportance
		} else if op == "clear" {
			btn.Importance = widget.DangerImportance
		} else if strings.Contains("+-×÷", op) {
			btn.Importance = widget.MediumImportance
		}
		buttons[op] = btn
	}

	return buttons
}

func (c *CalculatorTab) appendNumber(num string) {
	current := c.display.Text
	if c.isNewInput {
		current = "0"
		c.isNewInput = false
	}

	if current == "0" && num != "." {
		current = num
	} else {
		current += num
	}

	c.display.SetText(current)
}

func (c *CalculatorTab) appendDecimal() {
	current := c.display.Text
	if c.isNewInput {
		current = "0"
		c.isNewInput = false
	}

	if !strings.Contains(current, ".") {
		current += "."
		c.display.SetText(current)
	}
}

func (c *CalculatorTab) setOperator(op string) {
	if c.expression == "" {
		c.expression = c.display.Text + " " + op
		c.result, _ = strconv.ParseFloat(c.display.Text, 64)
	} else {
		c.calculate()
		c.expression = c.display.Text + " " + op
	}
	c.isNewInput = true
}

func (c *CalculatorTab) calculate() {
	if c.expression == "" {
		return
	}

	current, err := strconv.ParseFloat(c.display.Text, 64)
	if err != nil {
		return
	}

	parts := strings.Fields(c.expression)
	if len(parts) != 2 {
		return
	}

	prev, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return
	}

	op := parts[1]

	switch op {
	case "+":
		c.result = prev + current
	case "-":
		c.result = prev - current
	case "×":
		c.result = prev * current
	case "÷":
		if current != 0 {
			c.result = prev / current
		} else {
			c.display.SetText("Error")
			c.expression = ""
			c.isNewInput = true
			return
		}
	}

	// نمایش نتیجه
	if c.result == float64(int64(c.result)) {
		c.display.SetText(strconv.Itoa(int(c.result)))
	} else {
		c.display.SetText(fmt.Sprintf("%.10g", c.result))
	}

	c.expression = ""
	c.isNewInput = true
}

func (c *CalculatorTab) clear() {
	c.display.SetText("0")
	c.expression = ""
	c.result = 0
	c.isNewInput = true
}

func (c *CalculatorTab) toggleSign() {
	current, err := strconv.ParseFloat(c.display.Text, 64)
	if err != nil {
		return
	}

	current = -current
	if current == float64(int64(current)) {
		c.display.SetText(strconv.Itoa(int(current)))
	} else {
		c.display.SetText(fmt.Sprintf("%.10g", current))
	}
}

func (c *CalculatorTab) percentage() {
	current, err := strconv.ParseFloat(c.display.Text, 64)
	if err != nil {
		return
	}

	result := current / 100
	if result == float64(int64(result)) {
		c.display.SetText(strconv.Itoa(int(result)))
	} else {
		c.display.SetText(fmt.Sprintf("%.10g", result))
	}
	c.isNewInput = true
}

