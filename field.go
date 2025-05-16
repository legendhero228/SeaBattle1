package main

import (
	"fmt"
	"math/rand"
)

func createField() *[11][11]string { // Создает массив
	var arr [11][11]string
	for i := range arr {
		for j := range arr[i] {
			arr[i][j] = "O"
		}
	}
	return &arr
}

func printField(arr *[11][11]string) { // Печатает массив
	for i := range arr {
		for j := range arr[i] {
			fmt.Printf(arr[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Println(" ")
}

func ezField() {
	name := "" // Имя игрока
	fmt.Printf("Введите свое имя\n")
	fmt.Scan(&name)
	bot := "Bot"
	fieldPlayer := createField()
	fieldBot1 := createField()
	fieldResultPlayer := createField()

	boardSize := 11 // Размер поля
	position := ""  // Переменная для поворота корабля
	x := 0
	y := 0
	size := 0 // Размер корабля

	shipMap := map[int]int{ // Мапа с кораблями
		4: 1,
		3: 2,
		2: 3,
		1: 4,
	}

	player := Ship{size, x, y}

	arrPosition := [4]string{"Лево", "Право", "Вниз", "Вверх"} // Массив с поворотами

	for size, count := range shipMap {
		for i := 0; i < count; {
			fmt.Printf("Введите координаты Х У и Поворот корабля. Размерность корабля %d", size)
			_, err := fmt.Scan(&x, &y, &position)

			player = Ship{size, x, y}

			temp := AddShip(player, position, boardSize, fieldPlayer)

			if err != nil {
				fmt.Printf("Ошибка ввода%v\n", err)
				temp = false
			}
			if temp {
				i++
				printField(fieldPlayer)
			} else {
				fmt.Printf("Нельзя добавить корабль\n")
			}

		}
	}
	fmt.Println("Поле бота ")
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
				//fmt.Printf("БОТ ТВОРИТ ХУЙНЮ\n")
			}
		}
	}
	fmt.Printf("Вы расставили свои корабли. Игра начинается\n")
	fmt.Printf("Введите координаты выстрела\n")
	playerShoot := true
	botShoot := false
	counterBot := 0
	counterPlayer := 0
	var allShip = 20

	for {

		if playerShoot { // Реализация атаки игрока

			fmt.Scan(&x, &y)
			playerShoot, counterPlayer = attack(fieldBot1, x, y, counterPlayer, name)
			botShoot = !playerShoot
			if playerShoot {
				fieldResultPlayer[y][x] = "X"
			} else {
				fieldResultPlayer[y][x] = "."
			}
			printField(fieldResultPlayer)
		}
		if botShoot { // Реализация атаки бота
			x = rand.Intn(11)
			y = rand.Intn(11)
			botShoot, counterBot = attack(fieldPlayer, x, y, counterBot, bot)
			playerShoot = !botShoot
		}
		// fmt.Printf("Каунтер бот %d, Каунтер игрок %d\n", counterBot, counterPlayer)
		if counterPlayer == allShip { // Проверка победы игрока
			fmt.Println("Вы победили")
			break
		}
		if counterBot == allShip { // Проверка победы бота
			fmt.Printf("Бот выиграл")
			break
		}

	}

}
