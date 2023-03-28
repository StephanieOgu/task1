package task1

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

// Функція для знаходження суми компонентів файлу
func SumFileComponents(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	sum := 0

	for {
		line, _, err := reader.ReadLine()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return 0, err
		}
		num, err := strconv.Atoi(string(line))
		if err != nil {
			return 0, err
		}
		sum += num
	}

	return sum, nil
}

// вирішення завдання 1
func CreateSqrtFile(filename string, nums []int) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for _, num := range nums {
		if num < 0 {
			fmt.Printf("Число %d не було додано, бо не можна брати корінь з негативного числа\n", num)
		} else {
			sqrt := math.Sqrt(float64(num))
			_, err := fmt.Fprintln(writer, sqrt)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func Main() {
	//Написати пакет функцій для роботи з файлами, що містять числові дані.
	//Врахувати ситуацію помилок та панік.
	//Сформувати файл із квадратного коріння цілих чисел.
	//Знайти <Суму компонент файлу> ( n = 1, таблиця 2).

	name := "task1.txt"
	nums := []int{1, -2, 4, 16, 9, 25, -7, 8, 9, 10}

	// записати значення у файл
	fmt.Println("Записати значення у файл")
	fmt.Println()
	fmt.Println("Значення:")
	fmt.Println(nums)

	err := CreateSqrtFile(name, nums)
	if err != nil {
		log.Fatal(err)
	}

	// знайти суму комнонет
	fmt.Println()
	fmt.Println()
	fmt.Println("Знайти суму комнонет")
	sum, err := SumFileComponents(name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The sum of the components of the sqrt file is: %d\n", sum)
}
