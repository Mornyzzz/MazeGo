package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"maze/entity"
	"os"
	"strconv"
	"strings"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("A1_Maze_Go-1")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	vbox, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	button, _ := gtk.ButtonNewWithLabel("Open file")
	vbox.PackStart(button, false, false, 0)

	drawingArea, _ := gtk.DrawingAreaNew()
	drawingArea.SetSizeRequest(500, 500)
	vbox.PackStart(drawingArea, false, false, 10)

	rowAdjusment, _ := gtk.AdjustmentNew(5, 1, 51, 1, 1, 1)
	rowButton, _ := gtk.SpinButtonNew(rowAdjusment, 1, 2)
	vbox.PackStart(rowButton, false, false, 0)

	colAdjusment, _ := gtk.AdjustmentNew(5, 1, 51, 1, 1, 1)
	colButton, _ := gtk.SpinButtonNew(colAdjusment, 1, 2)
	vbox.PackStart(colButton, false, false, 0)

	generateButton, _ := gtk.ButtonNewWithLabel("Generate Matrix")
	startCalcButton, _ := gtk.ButtonNewWithLabel("Execute")

	vbox.PackStart(generateButton, false, false, 0)
	vbox.PackStart(startCalcButton, false, false, 0)

	var (
		maze                 *entity.Maze
		startX               int     = -1
		startY               int     = -1
		endX                 int     = -1
		endY                 int     = -1
		cellXsize, cellYsize float64 = 0, 0
		radius               float64 = 0
		solvedMaze           *entity.MazeSolving
	)

	generateButton.Connect("clicked", func() {
		if maze != nil {
			maze = nil
		}
		if solvedMaze != nil {
			solvedMaze = nil
		}
		startX, startY, endX, endY = -1, -1, -1, -1
		cellXsize, cellYsize, radius = 0, 0, 0
		drawingArea.QueueDraw()

		rows := rowButton.GetValueAsInt()
		cols := colButton.GetValueAsInt()

		maze, err = entity.GenerateMaze(rows, cols)
		if err != nil {
			return
		}

		cellXsize = 500 / float64(maze.ColSize())
		cellYsize = 500 / float64(maze.RowSize())
		radius = cellXsize / 4
		if cellXsize > cellYsize {
			radius = cellYsize / 4
		}

		drawingArea.QueueDraw()
	})

	drawingArea.Connect("draw", func(da *gtk.DrawingArea, cr *cairo.Context) {
		cr.SetSourceRGB(1, 1, 1)
		cr.Paint()

		if maze != nil {
			cr.SetSourceRGB(0, 0, 0)
			cr.Rectangle(0, 0, 500, 500)
			cr.SetLineWidth(2)
			cr.Stroke()

			for i := 0; i < maze.RowSize(); i++ {
				for j := 0; j < maze.ColSize()-1; j++ {
					if maze.RightBorderMatrix[i][j] == 1 {
						x1 := float64(j+1) * cellXsize
						y1 := float64(i) * cellYsize
						x2 := float64(j+1) * cellXsize
						y2 := float64(i+1) * cellYsize
						cr.MoveTo(x1, y1)
						cr.LineTo(x2, y2)
						cr.SetLineWidth(2)
						cr.Stroke()
					}
				}
			}

			for i := 0; i < maze.RowSize()-1; i++ {
				for j := 0; j < maze.ColSize(); j++ {
					if maze.LowBorderMatrix[i][j] == 1 {
						x1 := float64(j) * cellXsize
						y1 := float64(i+1) * cellYsize
						x2 := float64(j+1) * cellXsize
						y2 := float64(i+1) * cellYsize
						cr.MoveTo(x1, y1)
						cr.LineTo(x2, y2)
						cr.SetLineWidth(2)
						cr.Stroke()
					}
				}
			}
		}

		if startX != -1 && startY != -1 {
			cr.SetSourceRGB(0, 1, 0)
			cr.Arc(
				float64(startX)*cellXsize+cellXsize/2,
				float64(startY)*cellYsize+cellYsize/2,
				radius, 0, 2*math.Pi,
			)
			cr.Fill()
		}

		if endX != -1 && endY != -1 {
			cr.SetSourceRGB(1, 0, 0)
			cr.Arc(
				float64(endX)*cellXsize+cellXsize/2,
				float64(endY)*cellYsize+cellYsize/2,
				radius, 0, 2*math.Pi,
			)
			cr.Fill()
		}

		if solvedMaze != nil {
			cr.SetLineWidth(2)
			cr.SetSourceRGB(0, 1, 0)

			currentX := startX
			currentY := startY
			var prevX, prevY int = -1, -1

			for !(currentX == endX && currentY == endY) {
				x1 := float64(currentX)*cellXsize + cellXsize/2
				y1 := float64(currentY)*cellYsize + cellYsize/2

				var nextX, nextY int
				moved := false

				if currentX+1 < solvedMaze.ColSize() && solvedMaze.SolvingMatrix[currentY][currentX+1] > 0 && maze.RightBorderMatrix[currentY][currentX] == 0 && !(currentX+1 == prevX && currentY == prevY) {
					nextX, nextY = currentX+1, currentY
					moved = true
				} else if currentY+1 < solvedMaze.RowSize() && solvedMaze.SolvingMatrix[currentY+1][currentX] > 0 && maze.LowBorderMatrix[currentY][currentX] == 0 && !(currentY+1 == prevY && currentX == prevX) {
					nextX, nextY = currentX, currentY+1
					moved = true
				} else if currentX-1 >= 0 && solvedMaze.SolvingMatrix[currentY][currentX-1] > 0 && maze.RightBorderMatrix[currentY][currentX-1] == 0 && !(currentX-1 == prevX && currentY == prevY) {
					nextX, nextY = currentX-1, currentY
					moved = true
				} else if currentY-1 >= 0 && solvedMaze.SolvingMatrix[currentY-1][currentX] > 0 && maze.LowBorderMatrix[currentY-1][currentX] == 0 && !(currentY-1 == prevY && currentX == prevX) {
					nextX, nextY = currentX, currentY-1
					moved = true
				}

				if !moved {
					break
				}

				x2 := float64(nextX)*cellXsize + cellXsize/2
				y2 := float64(nextY)*cellYsize + cellYsize/2
				cr.MoveTo(x1, y1)
				cr.LineTo(x2, y2)
				cr.Stroke()

				prevX, prevY = currentX, currentY
				currentX, currentY = nextX, nextY
			}
		}
	})
	drawingArea.AddEvents(int(gdk.EventType(gdk.BUTTON_PRESS_MASK)))
	drawingArea.Connect("button-press-event", func(da *gtk.DrawingArea, event *gdk.Event) bool {
		if solvedMaze != nil {
			solvedMaze = nil
			startX, startY, endX, endY = -1, -1, -1, -1
		}

		mouseEvent := gdk.EventButtonNewFromEvent(event)
		x := mouseEvent.X()
		y := mouseEvent.Y()

		col := int(x / cellXsize)
		row := int(y / cellYsize)

		if startX == -1 {
			startX, startY = col, row
		} else if endX == -1 {
			endX, endY = col, row
		} else {
			startX, startY = endX, endY
			endX, endY = col, row
		}

		log.Printf("Updated Start Point: (%d, %d)\n", startX, startY)
		log.Printf("Updated End Point: (%d, %d)\n", endX, endY)

		drawingArea.QueueDraw()
		return true
	})
	startCalcButton.Connect("clicked", func() {
		if startX != -1 && startY != -1 && endX != -1 && endY != -1 {
			solvedMaze, err = entity.SolvingMaze(maze, startX, startY, endX, endY)
			if err != nil {
				return
			}
			drawingArea.QueueDraw()
		}
	})
	button.Connect("clicked", func() {
		fileDialog, _ := gtk.FileChooserDialogNewWith2Buttons(
			"Choose a maze file",
			win,
			gtk.FILE_CHOOSER_ACTION_OPEN,
			"Cancel", gtk.RESPONSE_CANCEL,
			"Open", gtk.RESPONSE_ACCEPT,
		)

		if response := fileDialog.Run(); response == gtk.RESPONSE_ACCEPT {
			filename := fileDialog.GetFilename()

			startX, startY, endX, endY = -1, -1, -1, -1
			solvedMaze = nil
			cellXsize = 0
			cellYsize = 0
			radius = 0

			var err error

			maze, err = loadMazeFromFile(filename)
			if err != nil {
				warnDlg := gtk.MessageDialogNew(win, gtk.DIALOG_MODAL, gtk.MESSAGE_INFO, gtk.BUTTONS_OK, "%s: %s", "Bad file", filename)
				warnDlg.Run()
				warnDlg.Destroy()
				log.Println("Error loading maze:", err)
				fileDialog.Destroy()
				maze = nil
				return
			}

			cellXsize = 500 / float64(maze.ColSize())
			cellYsize = 500 / float64(maze.RowSize())
			radius = cellXsize / 4
			if cellXsize > cellYsize {
				radius = cellYsize / 4
			}
		} else {
			cancelDlg := gtk.MessageDialogNew(win, gtk.DIALOG_MODAL, gtk.MESSAGE_INFO, gtk.BUTTONS_OK, "%s", "No file was selected")
			cancelDlg.Run()
			cancelDlg.Destroy()
			fileDialog.Destroy()
			return
		}
		fileDialog.Hide()
		fileDialog.Destroy()
		win.QueueDraw()
	})
	win.Add(vbox)
	win.SetDefaultSize(500, 500)
	win.SetResizable(false)
	win.ShowAll()
	gtk.Main()
}

func loadMazeFromFile(filename string) (*entity.Maze, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed load maze from file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sizes []string

	if scanner.Scan() {
		line := scanner.Text()
		sizes = strings.Fields(line)
	}

	if len(sizes) != 2 {
		return nil, fmt.Errorf("bad sizes for matrix")
	}

	m, err := strconv.Atoi(sizes[0])
	if err != nil {
		return nil, err
	}

	n, err := strconv.Atoi(sizes[1])
	if err != nil {
		return nil, err
	}

	maze := entity.NewMaze(m, n)
	for i := 0; i < m && scanner.Scan(); i++ {
		line := scanner.Text()
		digits := strings.Fields(line)
		for j := 0; j < n; j++ {
			digit, err := strconv.Atoi(digits[j])
			if err != nil {
				return nil, fmt.Errorf("scan digit to matrix: %ws", err)
			}
			maze.RightBorderMatrix[i][j] = digit
		}
	}

	if scanner.Scan() {
		_ = scanner.Text()
	}

	for i := 0; i < m && scanner.Scan(); i++ {
		line := scanner.Text()
		digits := strings.Fields(line)
		for j := 0; j < n; j++ {
			digit, err := strconv.Atoi(digits[j])
			if err != nil {
				return nil, fmt.Errorf("scan digit to matrix: %ws", err)
			}
			maze.LowBorderMatrix[i][j] = digit
		}
	}

	return maze, nil
}
