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
	fmt.Scan(&x, &y, &size)
	fmt.Println("Ваш выбор", x, y, size)
	fmt.Println("Куда вы хотите направить корабль (Лево,Право,Вверх,Вниз)")
	fmt.Scan(&position)
	ship := Ship{size, x, y, position, boardSize, &arr}
	AddShip(ship)

}
