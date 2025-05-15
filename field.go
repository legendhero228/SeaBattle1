package main

import (
	"fmt"
	"math/rand"
)

func createField() *[11][11]string {
	var arr [11][11]string
	for i := range arr {
		for j := range arr[i] {
			arr[i][j] = "O"
		}
	}
	return &arr
}

func ezField() {
	fieldPlayer := createField()
	fieldBot1 := createField()

	boardSize := 11
	fmt.Println("Выберите координаты расположения и размер вашего корабля (Формат ввода X Y Размер)")
	position := ""
	x := 0
	y := 0
	size := 0
	x = 1
	y = 1
	size = 4
	position = "Право"
	fmt.Println("Ваш выбор", x, y, size)
	fmt.Println("Куда вы хотите направить корабль (Лево,Право,Вверх,Вниз)")
	/* Всего 10 кораблей
	1 на 4 клетки
	2 на 3 клетки
	3 на 2 клетки
	4 на 1 клетку*/

	player := Ship{size, x, y}
	AddShip(player, position, boardSize, fieldPlayer)

	x = 1
	y = 3
	size = 3
	position = "Вниз"
	player = Ship{size, x, y}
	AddShip(player, position, boardSize, fieldPlayer)

	x = 7
	y = 7
	size = 3
	position = "Лево"
	player = Ship{size, x, y}
	AddShip(player, position, boardSize, fieldPlayer)

	x = 6
	y = 1
	size = 2
	position = "Право"
	player = Ship{size, x, y}
	AddShip(player, position, boardSize, fieldPlayer)

	x = 5
	y = 3
	size = 2
	position = "Вниз"
	player = Ship{size, x, y}
	AddShip(player, position, boardSize, fieldPlayer)

	x = 10
	y = 3
	size = 2
	position = "Вниз"
	player = Ship{size, x, y}
	AddShip(player, position, boardSize, fieldPlayer)

	x = 10
	y = 10
	size = 1
	position = "Вниз"
	player = Ship{size, x, y}
	AddShip(player, position, boardSize, fieldPlayer)

	x = 7
	y = 10
	size = 1
	position = "Вниз"
	player = Ship{size, x, y}
	AddShip(player, position, boardSize, fieldPlayer)

	x = 4
	y = 6
	size = 1
	position = "Вниз"
	player = Ship{size, x, y}
	AddShip(player, position, boardSize, fieldPlayer)

	x = 8
	y = 4
	size = 1
	player = Ship{size, x, y}
	AddShip(player, position, boardSize, fieldPlayer)

	fmt.Println("Поле бота ")

	shipMap := map[int]int{
		4: 1,
		3: 2,
		2: 3,
		1: 4,
	}
	arrPosition := [4]string{"Лево", "Право", "Вниз", "Вверх"}
	for size, count := range shipMap {
		for i := 0; i < count; {
			x = rand.Intn(10)
			y = rand.Intn(10)
			j := rand.Intn(4)
			position = arrPosition[j]

			bot1 := Ship{size, x, y}

			temp := AddShip(bot1, position, boardSize, fieldBot1)
			if temp {
				i++
			} else {
				fmt.Println("БОТ ТВОРИТ ХУЙНЮ")
			}
		}
	}
	playerShoot := true
	counterBot := 0
	counterPlayer := 0
	allShip := 5

	for {
		if playerShoot {
			x = rand.Intn(10)
			y = rand.Intn(10)
			playerShoot, counterPlayer = attack(fieldBot1, x, y, counterPlayer)
		} else {
			x = rand.Intn(10)
			y = rand.Intn(10)
			playerShoot, counterBot = attack(fieldPlayer, x, y, counterBot)
		}
		if counterPlayer >= allShip {
			fmt.Println("Вы победили")
			break
		}
		if counterBot >= allShip {
			fmt.Println("Бот выиграл")
			break
		}
	}
	//return

}
