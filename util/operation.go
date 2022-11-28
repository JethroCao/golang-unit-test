package util

import (
	"errors"
	"log"
)

func Add(a, b int) int {
	log.Printf("执行Add")
	return a + b
}

func Subtract(a, b int) int {
	log.Printf("执行Subtract")
	return a - b
}

func Multiply(a, b int) int {
	log.Printf("执行Multiply")
	return a * b
}

func Division(a, b int) (int, error) {
	log.Println("执行Division")
	if b == 0 {
		return 0, errors.New("被除数不能为 0")
	}
	return a / b, nil
}
