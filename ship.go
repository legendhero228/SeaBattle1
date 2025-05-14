package main

import "fmt"

type Ship struct {
	size      int
	X         int
	Y         int
	position  string
	boardSize int
	arr       *[11][11]string
}

func AddShip(ship Ship) {
	canPlace := true
	switch ship.position {
	case "Право":
		if ship.X+ship.size > ship.boardSize {
			canPlace = false
		} else {
			for i := 0; i < ship.size; i++ {
				if ship.arr[ship.Y][ship.X+i] == "+" {
					canPlace = false
					break
				}
			}
		}
		if canPlace {
			for i := 0; i < ship.size; i++ {
				ship.arr[ship.Y][ship.X+i] = "+"
			}
		} else {
			fmt.Println("Не влазит ебалай")
		}
	case "Лево":
		if ship.X-ship.size <= 0 {
			canPlace = false
		} else {
			for i := 0; i < ship.size; i++ {
				if ship.arr[ship.Y][ship.X-i] == "+" {
					canPlace = false
					break
				}
			}
		}
		if canPlace {
			for i := 0; i < ship.size; i++ {
				ship.arr[ship.Y][ship.X-i] = "+"
			}
		} else {
			fmt.Println("Не влазит ебалай")
		}

	case "Вверх":
		if ship.Y+ship.size > ship.boardSize {
			canPlace = false
		} else {
			for i := 0; i < ship.size; i++ {
				if ship.arr[ship.Y+i][ship.X] == "+" {
					canPlace = false
					break
				}
			}
		}
		if canPlace {
			for i := 0; i < ship.size; i++ {
				ship.arr[ship.Y+i][ship.X] = "+"
			}
		} else {
			fmt.Println("Не влазит ебалай")
		}

	case "Вниз":
		if ship.Y-ship.size > ship.boardSize {
			canPlace = false
		} else {
			for i := 0; i < ship.size; i++ {
				if ship.arr[ship.Y-i][ship.X] == "+" {
					canPlace = false
					break
				}
			}
		}
		if canPlace {
			for i := 0; i < ship.size; i++ {
				ship.arr[ship.Y-i][ship.X] = "+"
			}
		} else {
			fmt.Println("Не влазит ебалай")
		}

	default:
		fmt.Println("Хуйню не неси")
	}

	for i := range ship.arr {
		for j := range ship.arr[i] {
			fmt.Printf(ship.arr[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Println(" ")

	//fmt.Println("Тут уже есть корабль. Выберите новые координаты")

}

//return ship

func attack(ship *Ship) {
	fmt.Println("Совершите выстрел. Введите координаты")
	// тут типа скан
	x := 1
	y := 1

	for {
		switch ship.arr[y][x] {
		case "+":
			fmt.Println("Попал. Сделай еще один выстрел")
			ship.arr[y][x] = "X"
			// тут еще скан
			x = 3
			y = 1

		case "O":
			fmt.Println("Ха-ха-ха долбаеб не попал")
			ship.arr[y][x] = "."
			return

		case "X":
			fmt.Println("Ты уже стрелял сюда")
			return

		case ".":
			fmt.Println("Ты уже стрелял сюда")
			return
		}
	}

}
