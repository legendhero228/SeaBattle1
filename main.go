package main

import "fmt"

func main() {
	boardSize := 11
	arr := [11][11]string{}
	for i := range arr {
		for j := range arr[i] {

			arr[i][j] = "О"
		}
	}
	/*for i := range arr {
		for j := range arr[i] {
			fmt.Printf(arr[i][j])
		}
		fmt.Printf("\n")
	}*/
	fmt.Println("Выберите координаты расположения и размер вашего корабля (Формат ввода X Y Размер)")
	x := 0
	y := 0
	size := 0
	var position string
	//position := "Право"
	//fmt.Scan(&x, &y, &size)
	//fmt.Scan(&position)
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
	4 на 1 клетку
	*/

	ship := Ship{size, x, y, position, boardSize, &arr}
	AddShip(ship)
	x = 1
	y = 3
	size = 3
	position = "Вниз"
	ship = Ship{size, x, y, position, boardSize, &arr}
	AddShip(ship)
	x = 7
	y = 7
	size = 3
	position = "Лево"
	ship = Ship{size, x, y, position, boardSize, &arr}
	AddShip(ship)
	x = 6
	y = 1
	size = 2
	position = "Право"
	ship = Ship{size, x, y, position, boardSize, &arr}
	AddShip(ship)
	x = 5
	y = 3
	size = 2
	position = "Вниз"
	ship = Ship{size, x, y, position, boardSize, &arr}
	AddShip(ship)
	x = 10
	y = 3
	size = 2
	position = "Вниз"
	ship = Ship{size, x, y, position, boardSize, &arr}
	AddShip(ship)
	x = 10
	y = 10
	size = 1
	position = "Вниз"
	ship = Ship{size, x, y, position, boardSize, &arr}
	AddShip(ship)
	x = 7
	y = 10
	size = 1
	position = "Вниз"
	ship = Ship{size, x, y, position, boardSize, &arr}
	AddShip(ship)
	x = 4
	y = 6
	size = 1
	position = "Вниз"
	ship = Ship{size, x, y, position, boardSize, &arr}
	AddShip(ship)
	x = 8
	y = 4
	size = 1
	position = "Вниз"
	ship = Ship{size, x, y, position, boardSize, &arr}
	AddShip(ship)
	attack(&ship)

}
