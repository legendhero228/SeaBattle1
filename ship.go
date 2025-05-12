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

func AddShip(ship Ship) Ship {

	switch ship.position {
	case "Право":
		for i := 0; i < ship.size; i++ {
			if ship.boardSize >= ship.size-ship.X {
				ship.arr[ship.X][ship.Y+i] = "+"
			} else {
				fmt.Println("НЕ ВЛАЗИТ ЕБАЛАЙ")
			}
		}
	case "Лево":
		for i := 0; i < ship.size; i++ {
			if ship.boardSize >= ship.size+ship.X {
				ship.arr[ship.X][ship.Y-i] = "+"
			} else {
				fmt.Println("НЕ ВЛАЗИТ ЕБАЛАЙ")
			}
		}

	case "Вверх":
		for i := 0; i < ship.size; i++ {
			if ship.boardSize-ship.Y < ship.size+ship.X {
				ship.arr[ship.X-i][ship.Y] = "+"
			} else {
				fmt.Println("НЕ ВЛАЗИТ ЕБАЛАЙ")
			}
		}

	case "Вниз":
		for i := 0; i < ship.size; i++ {
			if ship.boardSize >= ship.size+ship.X {
				ship.arr[ship.X+i][ship.Y] = "+"
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

	return ship
}
