package main

import (
	"fmt"
	"time"
)

type MyString string
type Permanent struct {
	EmpId    int
	Basicpay int
	Pf       int
}

type Contract struct {
	EmpId    int
	Basicpay int
}

type Address struct {
	City, Detail string
}

type VowelsFinder interface {
	FindVowels() []rune
}

type SalaryCalculator interface {
	CalculateSalary() int
}
type SalaryCalculator2 interface {
	CalculateSalary2() int
}

//嵌套接口
type ContractOperations interface {
	SalaryCalculator
	SalaryCalculator2
}

type Add interface {
	Add()
}

func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

//salary of permanent employee is sum of basic pay and pf
func (p Permanent) CalculateSalary() int {
	return p.Basicpay + p.Pf
}

//salary of contract employee is the basic pay alone
func (c Contract) CalculateSalary() int {
	return c.Basicpay
}

func (c Contract) CalculateSalary2() int {
	return c.Basicpay
}

/*
total expense is calculated by iterating though the SalaryCalculator slice and summing
the salaries of the individual employees
*/
func TotalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%d", expense)
}

//使用指针接收实现接口
func (a *Address) Add() {
	fmt.Println(a)
}

func hello() {
	fmt.Println("Hello world goroutine")
}

func numbers() {
	for i := 1; i < 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d", i)
	}
}

func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func main() {

	//启动一个go协程
	go hello()
	time.Sleep(1 * time.Second)
	fmt.Println("main function")
	//启动多个go协程
	go numbers()

}
