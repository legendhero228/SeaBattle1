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

	switch ship.position {
	case "Право":
		for i := 0; i < ship.size; i++ {
			if ship.boardSize >= ship.size-ship.Y {
				ship.arr[ship.Y][ship.X+i] = "+"
			} else {
				fmt.Println("НЕ ВЛАЗИТ ЕБАЛАЙ")
			}
		}
	case "Лево":
		for i := 0; i < ship.size; i++ {
			if ship.boardSize >= ship.size+ship.Y {
				ship.arr[ship.Y][ship.X-i] = "+"
			} else {
				fmt.Println("НЕ ВЛАЗИТ ЕБАЛАЙ")
			}
		}

	case "Вверх":
		for i := 0; i < ship.size; i++ {
			if ship.boardSize-ship.X < ship.size+ship.X {
				ship.arr[ship.Y-i][ship.X] = "+"
			} else {
				fmt.Println("НЕ ВЛАЗИТ ЕБАЛАЙ")
			}
		}

	case "Вниз":
		for i := 0; i < ship.size; i++ {
			if ship.boardSize >= ship.size+ship.Y {
				ship.arr[ship.Y+i][ship.X] = "+"
			} else {
				fmt.Println("НЕ ВЛАЗИТ ЕБАЛАЙ")
			}
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

	//return ship
}

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
