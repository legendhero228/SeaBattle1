package main

import "fmt"

func createField(arr *[11][11]string) {
	for i := range arr {
		for j := range arr[i] {

			arr[i][j] = "О"

		}
	}
}
func ezField() {
	fieldPlayer := [11][11]string{}
	createField(&fieldPlayer)
	fieldBot1 := [11][11]string{}
	createField(&fieldBot1)

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

	player := Ship{size, x, y, position, boardSize, &fieldPlayer}
	AddShip(player)

	x = 1
	y = 3
	size = 3
	position = "Вниз"
	player = Ship{size, x, y, position, boardSize, &fieldPlayer}
	AddShip(player)

	x = 7
	y = 7
	size = 3
	position = "Лево"
	player = Ship{size, x, y, position, boardSize, &fieldPlayer}
	AddShip(player)

	x = 6
	y = 1
	size = 2
	position = "Право"
	player = Ship{size, x, y, position, boardSize, &fieldPlayer}
	AddShip(player)

	x = 5
	y = 3
	size = 2
	position = "Вниз"
	player = Ship{size, x, y, position, boardSize, &fieldPlayer}
	AddShip(player)

	x = 10
	y = 3
	size = 2
	position = "Вниз"
	player = Ship{size, x, y, position, boardSize, &fieldPlayer}
	AddShip(player)

	x = 10
	y = 10
	size = 1
	position = "Вниз"
	player = Ship{size, x, y, position, boardSize, &fieldPlayer}
	AddShip(player)

	x = 7
	y = 10
	size = 1
	position = "Вниз"
	player = Ship{size, x, y, position, boardSize, &fieldPlayer}
	AddShip(player)

	x = 4
	y = 6
	size = 1
	position = "Вниз"
	player = Ship{size, x, y, position, boardSize, &fieldPlayer}
	AddShip(player)

	x = 8
	y = 4
	size = 1
	position = "Вниз"
	player = Ship{size, x, y, position, boardSize, &fieldPlayer}

	AddShip(player)
	attack(&player)

	bot1 := Ship{size, x, y, position, boardSize, &fieldBot1}
	x = 1
	y = 5
	size = 3
	position = "Вниз"
	AddShip(bot1)

	return

}
