package geometric

import (
	"errors"
	"fmt"
	"math"
)

func Example() {
	gipotenuza, err := calculateGipotenuza(3, 4)
	if err != nil {
		fmt.Println(err)
		//panic(err)
	}
	fmt.Println("Ура, мы получили гипотенузу =", gipotenuza)

	// площадь круга Pi*R*R
	res, err := squareCircle(5)
	if err != nil {
		panic(err)
	}
	fmt.Println("площадь круга =", res)

	err = calcSquare(7)
	if err != nil {
		panic(err)
	}
}

// Для промяугольного треугольника
func calculateGipotenuza(catet1 int, catet2 int) (gipotenuza float64, err error) {
	if catet1 <= 0 {
		err = errors.New("invalid catet1")
		return
	}
	if catet2 <= 0 {
		err = errors.New("invalid catet2")
		return
	}
	gipotenuza2 := math.Pow(float64(catet1), 2) + math.Pow(float64(catet2), 2)
	gipotenuza = math.Sqrt(gipotenuza2)
	return
}

// Эта функция считает площади разных фигур одновременно
func calcSquare(v float64) (err error) {
	// calc circle (Площадь круга)
	res, err := squareCircle(v)
	if err != nil {
		return
	}
	fmt.Println("площадь круга =", res)

	// calc square (Площадь квадрата)
	res, err = squareSquare(v)
	if err != nil {
		return
	}
	fmt.Println("площадь квадрата =", res)
	return
}

func squareCircle(radius float64) (float64, error) {
	if radius <= 0 {
		err := errors.New("радиус меньше или равен нулю")
		// После return нужно поставить ноль и err, т.к. мы не именовали возвращаемые параметры.
		return 0, err
	}
	square := math.Pi * math.Pow(radius, 2)
	return square, nil
}

func squareSquare(side float64) (square float64, err error) {
	if side <= 0 {
		err = errors.New("сторона меньше или равно нулю")
		return
	}
	square = math.Pow(side, 2)
	return
}
