package main

import "fmt"

func run() {
	fmt.Println("ПРОВЕРКА РАБОТЫ")

	checkStack()
	fmt.Println()

	checkQueue()
	fmt.Println()

	checkList()
	fmt.Println()

	checkRomanNumbers()
	fmt.Println()

	checkRandomGrid()
}

func checkStack() {
	fmt.Println("СТЕК")
	s := CreateStack()

	s.Push(10)
	s.Push(20)
	s.Push(30)
	fmt.Println("Вставили: 10, 20, 30")

	val, _ := s.Top()
	fmt.Printf("Верхний элемент: %v\n", val)

	val, _ = s.Pop()
	fmt.Printf("Удалили: %v\n", val)
	val, _ = s.Pop()
	fmt.Printf("Удалили: %v\n", val)

	fmt.Printf("Стек пустой? %v\n", s.Empty())
}

func checkQueue() {
	fmt.Println("ОЧЕРЕДЬ")
	q := CreateQueue()

	q.Add("один")
	q.Add("два")
	q.Add("три")
	fmt.Println("Вставили: один, два, три")

	val, _ := q.Front()
	fmt.Printf("Первый элемент: %v\n", val)

	val, _ = q.Remove()
	fmt.Printf("Удалили: %v\n", val)
	val, _ = q.Remove()
	fmt.Printf("Удалили: %v\n", val)

	fmt.Printf("Очередь пустая? %v\n", q.Empty())
}

func checkList() {
	fmt.Println("СПИСОК")
	l := CreateList()

	l.Append(100)
	l.Append(200)
	l.Append(300)
	fmt.Printf("Список: %v\n", l.GetAll())
	fmt.Printf("Длина: %d\n", l.Count)

	val, _ := l.GetElement(1)
	fmt.Printf("Элемент на позиции 1: %v\n", val)

	l.Delete(1)
	fmt.Printf("После удаления позиции 1: %v\n", l.GetAll())

	l.Append(400)
	fmt.Printf("После добавления 400: %v\n", l.GetAll())
}

func checkRomanNumbers() {
	fmt.Println("РИМСКИЕ ЧИСЛА")
	conv := NewConverter()

	pairs := conv.GetExamplePairs()
	fmt.Println("Римские → Арабские:")
	for roman, arabic := range pairs {
		res, err := conv.ToArabic(roman)
		if err != nil {
			fmt.Printf("  Ошибка: %v\n", err)
		} else {
			fmt.Printf("  %s = %d (ожидалось %d) %v\n",
				roman, res, arabic, res == arabic)
		}
	}

	fmt.Println("\nАрабские → Римские:")
	nums := []int{1, 4, 9, 40, 90, 400, 900, 2024, 1950, 3999}
	for _, n := range nums {
		roman, err := conv.ToRoman(n)
		if err != nil {
			fmt.Printf("  Ошибка: %v\n", err)
		} else {
			fmt.Printf("  %d = %s\n", n, roman)
		}
	}

	fmt.Println("\nПроверка ошибок:")
	_, err := conv.ToArabic("IIII")
	fmt.Printf("  'IIII' → ошибка: %v\n", err)

	_, err = conv.ToRoman(4000)
	fmt.Printf("  4000 → ошибка: %v\n", err)
}

func checkRandomGrid() {
	fmt.Println("СЕТКА С УНИКАЛЬНЫМИ ЧИСЛАМИ")

	grid := NewNumberGrid(3, 3)

	err := grid.FillRandom(1, 20)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}

	fmt.Println("Сгенерированная сетка 3x3:")
	grid.Display()

	fmt.Printf("\nХарактеристики:\n")
	fmt.Printf("  Сумма элементов: %d\n", grid.Total())
	fmt.Printf("  Минимальное: %d\n", grid.MinValue())
	fmt.Printf("  Максимальное: %d\n", grid.MaxValue())

	num := grid.Data[1][1]
	rowIdx, colIdx, ok := grid.Locate(num)
	fmt.Printf("\nПоиск числа %d: строка %d, столбец %d (найдено: %v)\n",
		num, rowIdx, colIdx, ok)

	rowVals, _ := grid.GetRow(1)
	fmt.Printf("Вторая строка: %v\n", rowVals)

	colVals, _ := grid.GetColumn(1)
	fmt.Printf("Второй столбец: %v\n", colVals)

	fmt.Println("\nПроверка на уникальность:")
	seen := make(map[int]bool)
	allUnique := true
	for i := 0; i < grid.Rows; i++ {
		for j := 0; j < grid.Cols; j++ {
			if seen[grid.Data[i][j]] {
				allUnique = false
				fmt.Printf("  Найдено повторение: %d\n", grid.Data[i][j])
			}
			seen[grid.Data[i][j]] = true
		}
	}
	if allUnique {
		fmt.Println("  Все числа уникальны! ✓")
	}
}
