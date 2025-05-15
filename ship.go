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
			fmt.Println("Не влазит ебалай")
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
	}

	for i := range arr {
		for j := range arr[i] {
			fmt.Printf(arr[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Println(" ")

	//fmt.Println("Тут уже есть корабль. Выберите новые координаты")
	return canPlace
}

//return ship

func attack(arr *[11][11]string, x int, y int, counter int) (bool, int) {
	//fmt.Println("Совершите выстрел. Введите координаты")
	switch arr[y][x] {
	case "+":
		fmt.Println("Попал. Сделай еще один выстрел")
		arr[y][x] = "X"
		// тут еще скан
		counter++
		return true, counter

	case "O":
		fmt.Println("Ха-ха-ха долбаеб не попал")
		arr[y][x] = "."
		return false, counter

	case "X":
		fmt.Println("Ты уже стрелял сюда")
		return false, counter

	case ".":
		fmt.Println("Ты уже стрелял сюда")
		return false, counter

	default:
		fmt.Println("Че то не то")
		return false, counter
	}
	//return false, counter
}
