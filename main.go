package main

import (
	"fmt"
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"github.com/fatih/color"
	"os"
	"strconv"
)

var windows *ui.Window
var Selected *ui.Combobox



func CalcForm() ui.Control {
	
	box := ui.NewVerticalBox()
	box.SetPadded(true)

	group := ui.NewGroup(" ")
	group.SetMargined(true)
	box.Append(group, true)

	Form := ui.NewForm()
	Form.SetPadded(true)
	group.SetChild(Form)

	Num1 := ui.NewEntry()
	Num1.OnChanged(func(entry *ui.Entry){
		fmt.Scanln(Num1.Text())
	})
	Form.Append("Number: ", Num1, false)


	Symbol := ui.NewCombobox()
	Symbol.Append("+")
	Symbol.Append("-")
	Symbol.Append("x")
	Symbol.Append("÷")
	Symbol.SetSelected(0)
  Symbol.OnSelected(func(combobox *ui.Combobox){

    })
    Form.Append("Symbol: ", Symbol, false)


    Num2 := ui.NewEntry()
    Num2.OnChanged(func(entry *ui.Entry){

    })
    Form.Append("Number: ", Num2, false)

    answerRespond := ui.NewEntry()

    Answer := ui.NewButton("Answer :D")
	Progress := ui.NewProgressBar()
	Progress.SetValue(0)
	Form.Append("Progress: ", Progress, false)
    Answer.OnClicked(func(button *ui.Button){
			Selected := Symbol.Selected()
			switch Selected {
			case 0:
				Progress.SetValue(100)
				number1 := Num1.Text()
 			    number2 := Num2.Text()
 			    var Nb1 int = stringToInt(number1)
 			    var Nb2 int = stringToInt(number2)
 			    answer := Nb1 + Nb2
 			    answerstr := strconv.Itoa(answer)

 	            answerRespond.SetText(answerstr)
	 	 case 1:
			 Progress.SetValue(100)
		 	 number1 := Num1.Text()
			 number2 := Num2.Text()
			 var Nb1 int = stringToInt(number1)
			 var Nb2 int = stringToInt(number2)
			 answer := Nb1 - Nb2
			 answerstr := strconv.Itoa(answer)

	     answerRespond.SetText(answerstr)
		case 2:
			Progress.SetValue(100)
		 	 number1 := Num1.Text()
			 number2 := Num2.Text()
			 var Nb1 int = stringToInt(number1)
			 var Nb2 int = stringToInt(number2)
			 answer := Nb1 * Nb2
			 answerstr := strconv.Itoa(answer)

			 answerRespond.SetText(answerstr)
	 case 3:
		 Progress.SetValue(100)
		  number1 := Num1.Text()
		  number2 := Num2.Text()
		  var Nb1 int = stringToInt(number1)
		  var Nb2 int = stringToInt(number2)
		  answer := Nb1 / Nb2
		  answerstr := strconv.Itoa(answer)

		  answerRespond.SetText(answerstr)
			}

    })
    box.Append(Answer, false)

	Clear := ui.NewButton("Clear")
	box.Append(Clear, false)
	Clear.OnClicked(func(Button *ui.Button){
		Num1.SetText("")
		Num2.SetText("")
		answerRespond.SetText("")
		Progress.SetValue(0)
	})


    answerRespond.SetReadOnly(true)
    Form.Append("Answer: ", answerRespond, false)

	return box
}

func AboutForm() ui.Control {
	box := ui.NewVerticalBox()
	box.SetPadded(true)

	Label0 := ui.NewLabel("")
	box.Append(Label0, false)

	Label1 := ui.NewLabel("Name: Calculator")
	box.Append(Label1, false)

	Label2 := ui.NewLabel("Author: OwoNico")
	box.Append(Label2, false)

	Label3 := ui.NewLabel("Version: 1.0")
	box.Append(Label3, false)

	return box
}

func Calc() {
	windows := ui.NewWindow("Calculator", 400, 550, true)

	windows.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		windows.Destroy()
		return true
	})

	tabs := ui.NewTab()
	windows.SetChild(tabs)

	windows.SetMargined(true)
	tabs.Append("Calculator", CalcForm())
	tabs.SetMargined(0, true)

	tabs.Append("About", AboutForm())
	tabs.SetMargined(0, true)

	windows.Show()

}

func stringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	return i
}

func main() {
	color.Cyan("#Calculator by OwoNico")
	color.Blue("https://github.com/OwoNico")
	fmt.Println("")
	w := color.New(color.FgWhite).Add(color.Bold)
	y := color.New(color.FgYellow).Add(color.Bold)
	g := color.New(color.FgGreen).Add(color.Bold)
	w.Println("Loading Calculator...")
	w.Println("Loading API...")
	w.Println("Loading UI...")
	w.Println("Loaded all threads")
	y.Println("Launching Calculator...")
	g.Println("Launched")
	ui.Main(Calc)

}
