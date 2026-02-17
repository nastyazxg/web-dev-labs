package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type UniqueGrid struct {
	RowsCount int
	ColsCount int
	Data      [][]int
	Values    map[int]bool
}

func CreateUniqueGrid(rows, cols int) *UniqueGrid {
	data := make([][]int, rows)
	for i := range data {
		data[i] = make([]int, cols)
	}

	return &UniqueGrid{
		RowsCount: rows,
		ColsCount: cols,
		Data:      data,
		Values:    make(map[int]bool),
	}
}

func (g *UniqueGrid) FillRandom(min, max int) error {
	totalCells := g.RowsCount * g.ColsCount
	availableRange := max - min + 1

	if availableRange < totalCells {
		return errors.New("не хватает уникальных чисел в заданном диапазоне")
	}

	rand.Seed(time.Now().UnixNano())

	allNumbers := make([]int, availableRange)
	for i := 0; i < availableRange; i++ {
		allNumbers[i] = min + i
	}

	for i := len(allNumbers) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		allNumbers[i], allNumbers[j] = allNumbers[j], allNumbers[i]
	}

	idx := 0
	for i := 0; i < g.RowsCount; i++ {
		for j := 0; j < g.ColsCount; j++ {
			g.Data[i][j] = allNumbers[idx]
			g.Values[allNumbers[idx]] = true
			idx++
		}
	}

	return nil
}

func (g *UniqueGrid) FillRandomAlternative(min, max int) error {
	totalCells := g.RowsCount * g.ColsCount
	availableRange := max - min + 1

	if availableRange < totalCells {
		return errors.New("не хватает уникальных чисел в заданном диапазоне")
	}

	rand.Seed(time.Now().UnixNano())
	g.Values = make(map[int]bool)

	for i := 0; i < g.RowsCount; i++ {
		for j := 0; j < g.ColsCount; j++ {
			for {
				num := min + rand.Intn(availableRange)
				if !g.Values[num] {
					g.Data[i][j] = num
					g.Values[num] = true
					break
				}
			}
		}
	}

	return nil
}

func (g *UniqueGrid) At(row, col int) (int, error) {
	if row < 0 || row >= g.RowsCount || col < 0 || col >= g.ColsCount {
		return 0, errors.New("индекс выходит за границы")
	}
	return g.Data[row][col], nil
}

func (g *UniqueGrid) Update(row, col, val int) error {
	if row < 0 || row >= g.RowsCount || col < 0 || col >= g.ColsCount {
		return errors.New("индекс выходит за границы")
	}

	if g.Values[val] {
		return errors.New("такое число уже присутствует в сетке")
	}

	oldVal := g.Data[row][col]
	delete(g.Values, oldVal)

	g.Data[row][col] = val
	g.Values[val] = true

	return nil
}

func (g *UniqueGrid) Display() {
	for i := 0; i < g.RowsCount; i++ {
		for j := 0; j < g.ColsCount; j++ {
			fmt.Printf("%4d ", g.Data[i][j])
		}
		fmt.Println()
	}
}

func (g *UniqueGrid) ExtractRow(row int) ([]int, error) {
	if row < 0 || row >= g.RowsCount {
		return nil, errors.New("номер строки вне допустимого диапазона")
	}

	result := make([]int, g.ColsCount)
	copy(result, g.Data[row])
	return result, nil
}

func (g *UniqueGrid) ExtractColumn(col int) ([]int, error) {
	if col < 0 || col >= g.ColsCount {
		return nil, errors.New("номер колонки вне допустимого диапазона")
	}

	result := make([]int, g.RowsCount)
	for i := 0; i < g.RowsCount; i++ {
		result[i] = g.Data[i][col]
	}
	return result, nil
}

func (g *UniqueGrid) Locate(val int) (row, col int, exists bool) {
	for i := 0; i < g.RowsCount; i++ {
		for j := 0; j < g.ColsCount; j++ {
			if g.Data[i][j] == val {
				return i, j, true
			}
		}
	}
	return -1, -1, false
}

func (g *UniqueGrid) Total() int {
	sum := 0
	for i := 0; i < g.RowsCount; i++ {
		for j := 0; j < g.ColsCount; j++ {
			sum += g.Data[i][j]
		}
	}
	return sum
}

func (g *UniqueGrid) MinValue() int {
	minimum := g.Data[0][0]
	for i := 0; i < g.RowsCount; i++ {
		for j := 0; j < g.ColsCount; j++ {
			if g.Data[i][j] < minimum {
				minimum = g.Data[i][j]
			}
		}
	}
	return minimum
}

func (g *UniqueGrid) MaxValue() int {
	maximum := g.Data[0][0]
	for i := 0; i < g.RowsCount; i++ {
		for j := 0; j < g.ColsCount; j++ {
			if g.Data[i][j] > maximum {
				maximum = g.Data[i][j]
			}
		}
	}
	return maximum
}
