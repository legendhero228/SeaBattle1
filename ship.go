package main

import (
	"fmt"
)

type Ship struct {
	size int
	X    int
	Y    int
	//position  string
	//boardSize int
	//arr       *[11][11]string
}

func AddShip(ship Ship, position string, boardSize int, arr *[11][11]string) bool {
	canPlace := true
	switch position {
	case "Право":
		if ship.X+ship.size > boardSize {
			canPlace = false
		} else {
			for i := 0; i < ship.size; i++ {
				if arr[ship.Y][ship.X+i] == "+" {
					canPlace = false
					break
				}
			}
		}
		if canPlace {
			for i := 0; i < ship.size; i++ {
				arr[ship.Y][ship.X+i] = "+"
			}
		} else {
			//fmt.Println("Не влазит ебалай")
		}
	case "Лево":
		if ship.X-ship.size <= 0 {
			canPlace = false
		} else {
			for i := 0; i < ship.size; i++ {
				if arr[ship.Y][ship.X-i] == "+" {
					canPlace = false
					break
				}
			}
		}
		if canPlace {
			for i := 0; i < ship.size; i++ {
				arr[ship.Y][ship.X-i] = "+"
			}
		} else {
			fmt.Println("Не влазит ебалай")
		}

	case "Вверх":
		if ship.Y+ship.size > boardSize {
			canPlace = false
		} else {
			for i := 0; i < ship.size; i++ {
				if arr[ship.Y+i][ship.X] == "+" {
					canPlace = false
					break
				}
			}
		}
		if canPlace {
			for i := 0; i < ship.size; i++ {
				arr[ship.Y+i][ship.X] = "+"
			}
		} else {
			fmt.Println("Не влазит ебалай")
		}

	case "Вниз":
		if ship.Y-ship.size+1 < 0 {
			canPlace = false
		} else {
			for i := 0; i < ship.size; i++ {
				if arr[ship.Y-i][ship.X] == "+" {
					canPlace = false
					break
				}
			}
		}
		if canPlace {
			for i := 0; i < ship.size; i++ {
				arr[ship.Y-i][ship.X] = "+"
			}
		} else {
			fmt.Println("Не влазит ебалай")
		}

	default:
		fmt.Println("Хуйню не неси")
		return false
	}

	return canPlace
}

func attack(arr *[11][11]string, x int, y int, counter int, name string) (bool, int) {

	switch arr[y][x] {
	case "+":
		fmt.Printf("%s, попал. Сделай еще один выстрел\n", name)
		arr[y][x] = "X"
		counter++

		return true, counter

	case "O":
		fmt.Printf("Ха-ха-ха %s не попал\n", name)
		arr[y][x] = "."
		return false, counter

	case "X":
		fmt.Printf("%s, ты уже стрелял сюда\n", name)
		return false, counter

	case ".":
		fmt.Printf("%s, ты уже стрелял сюда\n", name)
		return false, counter

	default:
		fmt.Printf("%s, че то не то\n", name)
		return false, counter
	}
	//return false, counter
}
