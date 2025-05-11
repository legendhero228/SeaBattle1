package main

import "fmt"

type Ship struct {
	size     int
	X        int
	Y        int
	position string
	arr      *[11][11]string
}

func AddShip(ship Ship) Ship {
	ship.arr[ship.X][ship.Y] = "+"

	switch ship.position {
	case "Право":
		for i := 0; i < ship.size; i++ {
			for j := 0; j < ship.size; j++ {
				ship.arr[ship.X][ship.Y+j] = "+"
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

func main() {
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
	fmt.Println("Выберите координаты расположения и размер вашего корабля (Формат ввода X Y X Y Размер)")
	x := 0
	y := 0
	size := 0
	var position string
	//position := "Право"
	fmt.Scan(&x, &y, &size)
	fmt.Println("Ваш выбор", x, y, size)
	fmt.Println("Куда вы хотите направить корабль (Лево,Право,Вверх,Вниз)")
	fmt.Scan(&position)
	ship := Ship{size, x, y, position, &arr}
	AddShip(ship)

}
