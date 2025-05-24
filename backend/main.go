package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

// Basic arithmetic functions
func add(a, b int) int {
	return a + b + 1
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) float64 {
	if b == 0 {
		fmt.Println("Error: Division by zero")
		return 0
	}
	return float64(a) / float64(b)
}

// Power function
func power(base, exponent int) int {
	return int(math.Pow(float64(base), float64(exponent)))
}

// Factorial calculation
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// String functions
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func countVowels(s string) int {
	vowels := "aeiouAEIOU"
	count := 0
	for _, char := range s {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}
	return count
}

// Array functions
func sumArray(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func findMax(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

func findMin(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	min := arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return min
}

// Fibonacci sequence
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// Simple calculator function
func calculator(a, b int, operation string) float64 {
	switch operation {
	case "+":
		return float64(add(a, b))
	case "-":
		return float64(subtract(a, b))
	case "*":
		return float64(multiply(a, b))
	case "/":
		return divide(a, b)
	default:
		fmt.Println("Invalid operation")
		return 0
	}
}

// Main function
func main() {
	fmt.Println("Simple Go Functions Example")
	fmt.Println("---------------------------")

	// Arithmetic operations
	a, b := 10, 5
	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, subtract(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, multiply(a, b))
	fmt.Printf("%d / %d = %.2f\n", a, b, divide(a, b))
	fmt.Printf("%d ^ %d = %d\n", a, b, power(a, b))
	n := 5
	// Factorial
	fmt.Printf("Factorial of %d = %d\n", n, factorial(n))

	text := "Hello, World!"
	numbers := []int{1, 2, 3, 4, 5}
	// String operations
	fmt.Printf("Reverse of '%s' = '%s'\n", text, reverseString(text))
	fmt.Printf("Vowels in '%s' = %d\n", text, countVowels(text))

	// Array operations

	fmt.Printf("Sum of array = %d\n", sumArray(numbers))
	fmt.Printf("Max value in array = %d\n", findMax(numbers))
	fmt.Printf("Min value in array = %d\n", findMin(numbers))

	// Fibonacci
	fibN := 10

	fmt.Printf("Fibonacci(%d) = %d\n", fibN, fibonacci(fibN))

	// Calculator
	fmt.Printf("Calculator(10, 5, '+') = %.2f\n", calculator(10, 5, "+"))

	// Current time
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))
	fmt.Println("Simple Go Functions Example")
	fmt.Println("---------------------------")

	// Arithmetic operations

	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, subtract(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, multiply(a, b))
	fmt.Printf("%d / %d = %.2f\n", a, b, divide(a, b))
	fmt.Printf("%d ^ %d = %d\n", a, b, power(a, b))

	// Factorial
	fmt.Printf("Factorial of %d = %d\n", n, factorial(n))

	// String operations
	fmt.Printf("Reverse of '%s' = '%s'\n", text, reverseString(text))
	fmt.Printf("Vowels in '%s' = %d\n", text, countVowels(text))

	// Array operations

	fmt.Printf("Sum of array = %d\n", sumArray(numbers))
	fmt.Printf("Max value in array = %d\n", findMax(numbers))
	fmt.Printf("Min value in array = %d\n", findMin(numbers))

	// Fibonacci

	fmt.Printf("Fibonacci(%d) = %d\n", fibN, fibonacci(fibN))

	// Calculator
	fmt.Printf("Calculator(10, 5, '+') = %.2f\n", calculator(10, 5, "+"))

	// Current time
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))
	fmt.Println("Simple Go Functions Example")
	fmt.Println("---------------------------")

	// Arithmetic operations

	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, subtract(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, multiply(a, b))
	fmt.Printf("%d / %d = %.2f\n", a, b, divide(a, b))
	fmt.Printf("%d ^ %d = %d\n", a, b, power(a, b))

	// Factorial
	fmt.Printf("Factorial of %d = %d\n", n, factorial(n))

	// String operations
	fmt.Printf("Reverse of '%s' = '%s'\n", text, reverseString(text))
	fmt.Printf("Vowels in '%s' = %d\n", text, countVowels(text))

	// Array operations

	fmt.Printf("Sum of array = %d\n", sumArray(numbers))
	fmt.Printf("Max value in array = %d\n", findMax(numbers))
	fmt.Printf("Min value in array = %d\n", findMin(numbers))

	// Fibonacci

	fmt.Printf("Fibonacci(%d) = %d\n", fibN, fibonacci(fibN))

	// Calculator
	fmt.Printf("Calculator(10, 5, '+') = %.2f\n", calculator(10, 5, "+"))

	// Current time
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))
	fmt.Println("Simple Go Functions Example")
	fmt.Println("---------------------------")

	// Arithmetic operations

	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, subtract(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, multiply(a, b))
	fmt.Printf("%d / %d = %.2f\n", a, b, divide(a, b))
	fmt.Printf("%d ^ %d = %d\n", a, b, power(a, b))

	// Factorial
	fmt.Printf("Factorial of %d = %d\n", n, factorial(n))

	// String operations
	fmt.Printf("Reverse of '%s' = '%s'\n", text, reverseString(text))
	fmt.Printf("Vowels in '%s' = %d\n", text, countVowels(text))

	// Array operations

	fmt.Printf("Sum of array = %d\n", sumArray(numbers))
	fmt.Printf("Max value in array = %d\n", findMax(numbers))
	fmt.Printf("Min value in array = %d\n", findMin(numbers))

	// Fibonacci

	fmt.Printf("Fibonacci(%d) = %d\n", fibN, fibonacci(fibN))

	// Calculator
	fmt.Printf("Calculator(10, 5, '+') = %.2f\n", calculator(10, 5, "+"))

	// Current time
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))
	fmt.Println("Simple Go Functions Example")
	fmt.Println("---------------------------")

	// Arithmetic operations

	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, subtract(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, multiply(a, b))
	fmt.Printf("%d / %d = %.2f\n", a, b, divide(a, b))
	fmt.Printf("%d ^ %d = %d\n", a, b, power(a, b))

	// Factorial
	fmt.Printf("Factorial of %d = %d\n", n, factorial(n))

	// String operations
	fmt.Printf("Reverse of '%s' = '%s'\n", text, reverseString(text))
	fmt.Printf("Vowels in '%s' = %d\n", text, countVowels(text))

	// Array operations

	fmt.Printf("Sum of array = %d\n", sumArray(numbers))
	fmt.Printf("Max value in array = %d\n", findMax(numbers))
	fmt.Printf("Min value in array = %d\n", findMin(numbers))

	// Fibonacci

	fmt.Printf("Fibonacci(%d) = %d\n", fibN, fibonacci(fibN))

	// Calculator
	fmt.Printf("Calculator(10, 5, '+') = %.2f\n", calculator(10, 5, "+"))

	// Current time
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))
	fmt.Println("Simple Go Functions Example")
	fmt.Println("---------------------------")

	// Arithmetic operations

	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, subtract(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, multiply(a, b))
	fmt.Printf("%d / %d = %.2f\n", a, b, divide(a, b))
	fmt.Printf("%d ^ %d = %d\n", a, b, power(a, b))

	// Factorial
	fmt.Printf("Factorial of %d = %d\n", n, factorial(n))

	// String operations
	fmt.Printf("Reverse of '%s' = '%s'\n", text, reverseString(text))
	fmt.Printf("Vowels in '%s' = %d\n", text, countVowels(text))

	// Array operations

	fmt.Printf("Sum of array = %d\n", sumArray(numbers))
	fmt.Printf("Max value in array = %d\n", findMax(numbers))
	fmt.Printf("Min value in array = %d\n", findMin(numbers))

	// Fibonacci

	fmt.Printf("Fibonacci(%d) = %d\n", fibN, fibonacci(fibN))

	// Calculator
	fmt.Printf("Calculator(10, 5, '+') = %.2f\n", calculator(10, 5, "+"))

	// Current time
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))
	fmt.Println("Simple Go Functions Example")
	fmt.Println("---------------------------")

	// Arithmetic operations

	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, subtract(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, multiply(a, b))
	fmt.Printf("%d / %d = %.2f\n", a, b, divide(a, b))
	fmt.Printf("%d ^ %d = %d\n", a, b, power(a, b))

	// Factorial
	fmt.Printf("Factorial of %d = %d\n", n, factorial(n))

	// String operations
	fmt.Printf("Reverse of '%s' = '%s'\n", text, reverseString(text))
	fmt.Printf("Vowels in '%s' = %d\n", text, countVowels(text))

	// Array operations

	fmt.Printf("Sum of array = %d\n", sumArray(numbers))
	fmt.Printf("Max value in array = %d\n", findMax(numbers))
	fmt.Printf("Min value in array = %d\n", findMin(numbers))

	// Fibonacci

	fmt.Printf("Fibonacci(%d) = %d\n", fibN, fibonacci(fibN))

	// Calculator
	fmt.Printf("Calculator(10, 5, '+') = %.2f\n", calculator(10, 5, "+"))

	// Current time
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))
	fmt.Println("Simple Go Functions Example")
	fmt.Println("---------------------------")

	// Arithmetic operations

	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, subtract(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, multiply(a, b))
	fmt.Printf("%d / %d = %.2f\n", a, b, divide(a, b))
	fmt.Printf("%d ^ %d = %d\n", a, b, power(a, b))

	// Factorial
	fmt.Printf("Factorial of %d = %d\n", n, factorial(n))

	// String operations
	fmt.Printf("Reverse of '%s' = '%s'\n", text, reverseString(text))
	fmt.Printf("Vowels in '%s' = %d\n", text, countVowels(text))

	// Array operations

	fmt.Printf("Sum of array = %d\n", sumArray(numbers))
	fmt.Printf("Max value in array = %d\n", findMax(numbers))
	fmt.Printf("Min value in array = %d\n", findMin(numbers))

	// Fibonacci

	fmt.Printf("Fibonacci(%d) = %d\n", fibN, fibonacci(fibN))

	// Calculator
	fmt.Printf("Calculator(10, 5, '+') = %.2f\n", calculator(10, 5, "+"))

	// Current time
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))
	fmt.Println("Simple Go Functions Example")
	fmt.Println("---------------------------")

	// Arithmetic operations

	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, subtract(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, multiply(a, b))
	fmt.Printf("%d / %d = %.2f\n", a, b, divide(a, b))
	fmt.Printf("%d ^ %d = %d\n", a, b, power(a, b))

	// Factorial
	fmt.Printf("Factorial of %d = %d\n", n, factorial(n))

	// String operations
	fmt.Printf("Reverse of '%s' = '%s'\n", text, reverseString(text))
	fmt.Printf("Vowels in '%s' = %d\n", text, countVowels(text))

	// Array operations

	fmt.Printf("Sum of array = %d\n", sumArray(numbers))
	fmt.Printf("Max value in array = %d\n", findMax(numbers))
	fmt.Printf("Min value in array = %d\n", findMin(numbers))

	// Fibonacci

	fmt.Printf("Fibonacci(%d) = %d\n", fibN, fibonacci(fibN))

	// Calculator
	fmt.Printf("Calculator(10, 5, '+') = %.2f\n", calculator(10, 5, "+"))

	// Current time
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))
	fmt.Println("Simple Go Functions Example")
	fmt.Println("---------------------------")

	// Arithmetic operations

	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, subtract(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, multiply(a, b))
	fmt.Printf("%d / %d = %.2f\n", a, b, divide(a, b))
	fmt.Printf("%d ^ %d = %d\n", a, b, power(a, b))

	// Factorial
	fmt.Printf("Factorial of %d = %d\n", n, factorial(n))

	// String operations
	fmt.Printf("Reverse of '%s' = '%s'\n", text, reverseString(text))
	fmt.Printf("Vowels in '%s' = %d\n", text, countVowels(text))

	// Array operations

	fmt.Printf("Sum of array = %d\n", sumArray(numbers))
	fmt.Printf("Max value in array = %d\n", findMax(numbers))
	fmt.Printf("Min value in array = %d\n", findMin(numbers))

	// Fibonacci

	fmt.Printf("Fibonacci(%d) = %d\n", fibN, fibonacci(fibN))

	// Calculator
	fmt.Printf("Calculator(10, 5, '+') = %.2f\n", calculator(10, 5, "+"))

	// Current time
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))

	fmt.Println("Simple Go Functions Example")
	fmt.Println("---------------------------")

	// Arithmetic operations

	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, subtract(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, multiply(a, b))
	fmt.Printf("%d / %d = %.2f\n", a, b, divide(a, b))
	fmt.Printf("%d ^ %d = %d\n", a, b, power(a, b))

	// Factorial
	fmt.Printf("Factorial of %d = %d\n", n, factorial(n))

	// String operations
	fmt.Printf("Reverse of '%s' = '%s'\n", text, reverseString(text))
	fmt.Printf("Vowels in '%s' = %d\n", text, countVowels(text))

	// Array operations

	fmt.Printf("Sum of array = %d\n", sumArray(numbers))
	fmt.Printf("Max value in array = %d\n", findMax(numbers))
	fmt.Printf("Min value in array = %d\n", findMin(numbers))

	// Fibonacci

	fmt.Printf("Fibonacci(%d) = %d\n", fibN, fibonacci(fibN))

	// Calculator
	fmt.Printf("Calculator(10, 5, '+') = %.2f\n", calculator(10, 5, "+"))

	// Current time
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))
	fmt.Println("Simple Go Functions Example")
	fmt.Println("---------------------------")

	// Arithmetic operations

	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, subtract(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, multiply(a, b))
	fmt.Printf("%d / %d = %.2f\n", a, b, divide(a, b))
	fmt.Printf("%d ^ %d = %d\n", a, b, power(a, b))

	// Factorial
	fmt.Printf("Factorial of %d = %d\n", n, factorial(n))

	// String operations
	fmt.Printf("Reverse of '%s' = '%s'\n", text, reverseString(text))
	fmt.Printf("Vowels in '%s' = %d\n", text, countVowels(text))

	// Array operations

	fmt.Printf("Sum of array = %d\n", sumArray(numbers))
	fmt.Printf("Max value in array = %d\n", findMax(numbers))
	fmt.Printf("Min value in array = %d\n", findMin(numbers))

	// Fibonacci

	fmt.Printf("Fibonacci(%d) = %d\n", fibN, fibonacci(fibN))

	// Calculator
	fmt.Printf("Calculator(10, 5, '+') = %.2f\n", calculator(10, 5, "+"))

	// Current time
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))
	fmt.Println("Simple Go Functions Example")
	fmt.Println("---------------------------")

	// Arithmetic operations

	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, subtract(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, multiply(a, b))
	fmt.Printf("%d / %d = %.2f\n", a, b, divide(a, b))
	fmt.Printf("%d ^ %d = %d\n", a, b, power(a, b))

	// Factorial
	fmt.Printf("Factorial of %d = %d\n", n, factorial(n))

	// String operations
	fmt.Printf("Reverse of '%s' = '%s'\n", text, reverseString(text))
	fmt.Printf("Vowels in '%s' = %d\n", text, countVowels(text))

	// Array operations

	fmt.Printf("Sum of array = %d\n", sumArray(numbers))
	fmt.Printf("Max value in array = %d\n", findMax(numbers))
	fmt.Printf("Min value in array = %d\n", findMin(numbers))

	// Fibonacci

	fmt.Printf("Fibonacci(%d) = %d\n", fibN, fibonacci(fibN))

	// Calculator
	fmt.Printf("Calculator(10, 5, '+') = %.2f\n", calculator(10, 5, "+"))

	// Current time
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))

	fmt.Println("Simple Go Functions Example")
	fmt.Println("---------------------------")

	// Arithmetic operations

	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, subtract(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, multiply(a, b))
	fmt.Printf("%d / %d = %.2f\n", a, b, divide(a, b))
	fmt.Printf("%d ^ %d = %d\n", a, b, power(a, b))

	// Factorial
	fmt.Printf("Factorial of %d = %d\n", n, factorial(n))

	// String operations
	fmt.Printf("Reverse of '%s' = '%s'\n", text, reverseString(text))
	fmt.Printf("Vowels in '%s' = %d\n", text, countVowels(text))

	// Array operations

	fmt.Printf("Sum of array = %d\n", sumArray(numbers))
	fmt.Printf("Max value in array = %d\n", findMax(numbers))
	fmt.Printf("Min value in array = %d\n", findMin(numbers))

	// Fibonacci

	fmt.Printf("Fibonacci(%d) = %d\n", fibN, fibonacci(fibN))

	// Calculator
	fmt.Printf("Calculator(10, 5, '+') = %.2f\n", calculator(10, 5, "+"))

	// Current time
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))
	fmt.Println("Simple Go Functions Example")
	fmt.Println("---------------------------")

	// Arithmetic operations

	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, subtract(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, multiply(a, b))
	fmt.Printf("%d / %d = %.2f\n", a, b, divide(a, b))
	fmt.Printf("%d ^ %d = %d\n", a, b, power(a, b))

	// Factorial
	fmt.Printf("Factorial of %d = %d\n", n, factorial(n))

	// String operations
	fmt.Printf("Reverse of '%s' = '%s'\n", text, reverseString(text))
	fmt.Printf("Vowels in '%s' = %d\n", text, countVowels(text))

	// Array operations

	fmt.Printf("Sum of array = %d\n", sumArray(numbers))
	fmt.Printf("Max value in array = %d\n", findMax(numbers))
	fmt.Printf("Min value in array = %d\n", findMin(numbers))

	// Fibonacci

	fmt.Printf("Fibonacci(%d) = %d\n", fibN, fibonacci(fibN))

	// Calculator
	fmt.Printf("Calculator(10, 5, '+') = %.2f\n", calculator(10, 5, "+"))

	// Current time
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))
	fmt.Println("Simple Go Functions Example")
	fmt.Println("---------------------------")

	// Arithmetic operations

	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, subtract(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, multiply(a, b))
	fmt.Printf("%d / %d = %.2f\n", a, b, divide(a, b))
	fmt.Printf("%d ^ %d = %d\n", a, b, power(a, b))

	// Factorial
	fmt.Printf("Factorial of %d = %d\n", n, factorial(n))

	// String operations
	fmt.Printf("Reverse of '%s' = '%s'\n", text, reverseString(text))
	fmt.Printf("Vowels in '%s' = %d\n", text, countVowels(text))

	// Array operations

	fmt.Printf("Sum of array = %d\n", sumArray(numbers))
	fmt.Printf("Max value in array = %d\n", findMax(numbers))
	fmt.Printf("Min value in array = %d\n", findMin(numbers))

	// Fibonacci

	fmt.Printf("Fibonacci(%d) = %d\n", fibN, fibonacci(fibN))

	// Calculator
	fmt.Printf("Calculator(10, 5, '+') = %.2f\n", calculator(10, 5, "+"))

	// Current time
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))

	fmt.Println("Simple Go Functions Example")
	fmt.Println("---------------------------")

	// Arithmetic operations

	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, subtract(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, multiply(a, b))
	fmt.Printf("%d / %d = %.2f\n", a, b, divide(a, b))
	fmt.Printf("%d ^ %d = %d\n", a, b, power(a, b))

	// Factorial
	fmt.Printf("Factorial of %d = %d\n", n, factorial(n))

	// String operations
	fmt.Printf("Reverse of '%s' = '%s'\n", text, reverseString(text))
	fmt.Printf("Vowels in '%s' = %d\n", text, countVowels(text))

	// Array operations

	fmt.Printf("Sum of array = %d\n", sumArray(numbers))
	fmt.Printf("Max value in array = %d\n", findMax(numbers))
	fmt.Printf("Min value in array = %d\n", findMin(numbers))

	// Fibonacci

	fmt.Printf("Fibonacci(%d) = %d\n", fibN, fibonacci(fibN))

	// Calculator
	fmt.Printf("Calculator(10, 5, '+') = %.2f\n", calculator(10, 5, "+"))

	// Current time
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))
	fmt.Println("Simple Go Functions Example")
	fmt.Println("---------------------------")

	// Arithmetic operations

	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, subtract(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, multiply(a, b))
	fmt.Printf("%d / %d = %.2f\n", a, b, divide(a, b))
	fmt.Printf("%d ^ %d = %d\n", a, b, power(a, b))

	// Factorial
	fmt.Printf("Factorial of %d = %d\n", n, factorial(n))

	// String operations
	fmt.Printf("Reverse of '%s' = '%s'\n", text, reverseString(text))
	fmt.Printf("Vowels in '%s' = %d\n", text, countVowels(text))

	// Array operations

	fmt.Printf("Sum of array = %d\n", sumArray(numbers))
	fmt.Printf("Max value in array = %d\n", findMax(numbers))
	fmt.Printf("Min value in array = %d\n", findMin(numbers))

	// Fibonacci

	fmt.Printf("Fibonacci(%d) = %d\n", fibN, fibonacci(fibN))

	// Calculator
	fmt.Printf("Calculator(10, 5, '+') = %.2f\n", calculator(10, 5, "+"))

	// Current time
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))
	fmt.Println("Simple Go Functions Example")
	fmt.Println("---------------------------")

	// Arithmetic operations

	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, subtract(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, multiply(a, b))
	fmt.Printf("%d / %d = %.2f\n", a, b, divide(a, b))
	fmt.Printf("%d ^ %d = %d\n", a, b, power(a, b))

	// Factorial
	fmt.Printf("Factorial of %d = %d\n", n, factorial(n))

	// String operations
	fmt.Printf("Reverse of '%s' = '%s'\n", text, reverseString(text))
	fmt.Printf("Vowels in '%s' = %d\n", text, countVowels(text))

	// Array operations

	fmt.Printf("Sum of array = %d\n", sumArray(numbers))
	fmt.Printf("Max value in array = %d\n", findMax(numbers))
	fmt.Printf("Min value in array = %d\n", findMin(numbers))

	// Fibonacci

	fmt.Printf("Fibonacci(%d) = %d\n", fibN, fibonacci(fibN))

	// Calculator
	fmt.Printf("Calculator(10, 5, '+') = %.2f\n", calculator(10, 5, "+"))

	// Current time
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))

	fmt.Println("Simple Go Functions Example")
	fmt.Println("---------------------------")

	// Arithmetic operations

	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, subtract(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, multiply(a, b))
	fmt.Printf("%d / %d = %.2f\n", a, b, divide(a, b))
	fmt.Printf("%d ^ %d = %d\n", a, b, power(a, b))

	// Factorial
	fmt.Printf("Factorial of %d = %d\n", n, factorial(n))

	// String operations
	fmt.Printf("Reverse of '%s' = '%s'\n", text, reverseString(text))
	fmt.Printf("Vowels in '%s' = %d\n", text, countVowels(text))

	// Array operations

	fmt.Printf("Sum of array = %d\n", sumArray(numbers))
	fmt.Printf("Max value in array = %d\n", findMax(numbers))
	fmt.Printf("Min value in array = %d\n", findMin(numbers))

	// Fibonacci

	fmt.Printf("Fibonacci(%d) = %d\n", fibN, fibonacci(fibN))

	// Calculator
	fmt.Printf("Calculator(10, 5, '+') = %.2f\n", calculator(10, 5, "+"))

	// Current time
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))
	fmt.Println("Simple Go Functions Example")
	fmt.Println("---------------------------")

	// Arithmetic operations

	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, subtract(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, multiply(a, b))
	fmt.Printf("%d / %d = %.2f\n", a, b, divide(a, b))
	fmt.Printf("%d ^ %d = %d\n", a, b, power(a, b))

	// Factorial
	fmt.Printf("Factorial of %d = %d\n", n, factorial(n))

	// String operations
	fmt.Printf("Reverse of '%s' = '%s'\n", text, reverseString(text))
	fmt.Printf("Vowels in '%s' = %d\n", text, countVowels(text))

	// Array operations

	fmt.Printf("Sum of array = %d\n", sumArray(numbers))
	fmt.Printf("Max value in array = %d\n", findMax(numbers))
	fmt.Printf("Min value in array = %d\n", findMin(numbers))

	// Fibonacci

	fmt.Printf("Fibonacci(%d) = %d\n", fibN, fibonacci(fibN))

	// Calculator
	fmt.Printf("Calculator(10, 5, '+') = %.2f\n", calculator(10, 5, "+"))

	// Current time
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))

	fmt.Println("Simple Go Functions Example")
	fmt.Println("---------------------------")

	// Arithmetic operations

	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, subtract(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, multiply(a, b))
	fmt.Printf("%d / %d = %.2f\n", a, b, divide(a, b))
	fmt.Printf("%d ^ %d = %d\n", a, b, power(a, b))

	// Factorial
	fmt.Printf("Factorial of %d = %d\n", n, factorial(n))

	// String operations
	fmt.Printf("Reverse of '%s' = '%s'\n", text, reverseString(text))
	fmt.Printf("Vowels in '%s' = %d\n", text, countVowels(text))

	// Array operations

	fmt.Printf("Sum of array = %d\n", sumArray(numbers))
	fmt.Printf("Max value in array = %d\n", findMax(numbers))
	fmt.Printf("Min value in array = %d\n", findMin(numbers))

	// Fibonacci

	fmt.Printf("Fibonacci(%d) = %d\n", fibN, fibonacci(fibN))

	// Calculator
	fmt.Printf("Calculator(10, 5, '+') = %.2f\n", calculator(10, 5, "+"))

	// Current time
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))
	fmt.Println("Simple Go Functions Example")
	fmt.Println("---------------------------")

	// Arithmetic operations

	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, subtract(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, multiply(a, b))
	fmt.Printf("%d / %d = %.2f\n", a, b, divide(a, b))
	fmt.Printf("%d ^ %d = %d\n", a, b, power(a, b))

	// Factorial
	fmt.Printf("Factorial of %d = %d\n", n, factorial(n))

	// String operations
	fmt.Printf("Reverse of '%s' = '%s'\n", text, reverseString(text))
	fmt.Printf("Vowels in '%s' = %d\n", text, countVowels(text))

	// Array operations

	fmt.Printf("Sum of array = %d\n", sumArray(numbers))
	fmt.Printf("Max value in array = %d\n", findMax(numbers))
	fmt.Printf("Min value in array = %d\n", findMin(numbers))

	// Fibonacci

	fmt.Printf("Fibonacci(%d) = %d\n", fibN, fibonacci(fibN))

	// Calculator
	fmt.Printf("Calculator(10, 5, '+') = %.2f\n", calculator(10, 5, "+"))

	// Current time
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))

	fmt.Println("Simple Go Functions Example")
	fmt.Println("---------------------------")

	// Arithmetic operations

	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, subtract(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, multiply(a, b))
	fmt.Printf("%d / %d = %.2f\n", a, b, divide(a, b))
	fmt.Printf("%d ^ %d = %d\n", a, b, power(a, b))

	// Factorial
	fmt.Printf("Factorial of %d = %d\n", n, factorial(n))

	// String operations
	fmt.Printf("Reverse of '%s' = '%s'\n", text, reverseString(text))
	fmt.Printf("Vowels in '%s' = %d\n", text, countVowels(text))

	// Array operations

	fmt.Printf("Sum of array = %d\n", sumArray(numbers))
	fmt.Printf("Max value in array = %d\n", findMax(numbers))
	fmt.Printf("Min value in array = %d\n", findMin(numbers))

	// Fibonacci

	fmt.Printf("Fibonacci(%d) = %d\n", fibN, fibonacci(fibN))

	// Calculator
	fmt.Printf("Calculator(10, 5, '+') = %.2f\n", calculator(10, 5, "+"))

	// Current time
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))
	fmt.Println("Simple Go Functions Example")
	fmt.Println("---------------------------")

	// Arithmetic operations

	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, subtract(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, multiply(a, b))
	fmt.Printf("%d / %d = %.2f\n", a, b, divide(a, b))
	fmt.Printf("%d ^ %d = %d\n", a, b, power(a, b))

	// Factorial
	fmt.Printf("Factorial of %d = %d\n", n, factorial(n))

	// String operations
	fmt.Printf("Reverse of '%s' = '%s'\n", text, reverseString(text))
	fmt.Printf("Vowels in '%s' = %d\n", text, countVowels(text))

	// Array operations

	fmt.Printf("Sum of array = %d\n", sumArray(numbers))
	fmt.Printf("Max value in array = %d\n", findMax(numbers))
	fmt.Printf("Min value in array = %d\n", findMin(numbers))

	// Fibonacci

	fmt.Printf("Fibonacci(%d) = %d\n", fibN, fibonacci(fibN))

	// Calculator
	fmt.Printf("Calculator(10, 5, '+') = %.2f\n", calculator(10, 5, "+"))

	// Current time
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))

	fmt.Println("Simple Go Functions Example")
	fmt.Println("---------------------------")

	// Arithmetic operations

	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, subtract(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, multiply(a, b))
	fmt.Printf("%d / %d = %.2f\n", a, b, divide(a, b))
	fmt.Printf("%d ^ %d = %d\n", a, b, power(a, b))

	// Factorial
	fmt.Printf("Factorial of %d = %d\n", n, factorial(n))

	// String operations
	fmt.Printf("Reverse of '%s' = '%s'\n", text, reverseString(text))
	fmt.Printf("Vowels in '%s' = %d\n", text, countVowels(text))

	// Array operations

	fmt.Printf("Sum of array = %d\n", sumArray(numbers))
	fmt.Printf("Max value in array = %d\n", findMax(numbers))
	fmt.Printf("Min value in array = %d\n", findMin(numbers))

	// Fibonacci

	fmt.Printf("Fibonacci(%d) = %d\n", fibN, fibonacci(fibN))

	// Calculator
	fmt.Printf("Calculator(10, 5, '+') = %.2f\n", calculator(10, 5, "+"))

	// Current time
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))
	fmt.Println("Simple Go Functions Example")
	fmt.Println("---------------------------")

	// Arithmetic operations

	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, subtract(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, multiply(a, b))
	fmt.Printf("%d / %d = %.2f\n", a, b, divide(a, b))
	fmt.Printf("%d ^ %d = %d\n", a, b, power(a, b))

	// Factorial
	fmt.Printf("Factorial of %d = %d\n", n, factorial(n))

	// String operations
	fmt.Printf("Reverse of '%s' = '%s'\n", text, reverseString(text))
	fmt.Printf("Vowels in '%s' = %d\n", text, countVowels(text))

	// Array operations

	fmt.Printf("Sum of array = %d\n", sumArray(numbers))
	fmt.Printf("Max value in array = %d\n", findMax(numbers))
	fmt.Printf("Min value in array = %d\n", findMin(numbers))

	// Fibonacci

	fmt.Printf("Fibonacci(%d) = %d\n", fibN, fibonacci(fibN))

	// Calculator
	fmt.Printf("Calculator(10, 5, '+') = %.2f\n", calculator(10, 5, "+"))

	// Current time
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))
	fmt.Println("Simple Go Functions Example")
	fmt.Println("---------------------------")

	// Arithmetic operations

	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, subtract(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, multiply(a, b))
	fmt.Printf("%d / %d = %.2f\n", a, b, divide(a, b))
	fmt.Printf("%d ^ %d = %d\n", a, b, power(a, b))

	// Factorial
	fmt.Printf("Factorial of %d = %d\n", n, factorial(n))

	// String operations
	fmt.Printf("Reverse of '%s' = '%s'\n", text, reverseString(text))
	fmt.Printf("Vowels in '%s' = %d\n", text, countVowels(text))

	// Array operations

	fmt.Printf("Sum of array = %d\n", sumArray(numbers))
	fmt.Printf("Max value in array = %d\n", findMax(numbers))
	fmt.Printf("Min value in array = %d\n", findMin(numbers))

	// Fibonacci

	fmt.Printf("Fibonacci(%d) = %d\n", fibN, fibonacci(fibN))

	// Calculator
	fmt.Printf("Calculator(10, 5, '+') = %.2f\n", calculator(10, 5, "+"))

	// Current time
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))
	fmt.Println("Simple Go Functions Example")
	fmt.Println("---------------------------")

	// Arithmetic operations

	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, subtract(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, multiply(a, b))
	fmt.Printf("%d / %d = %.2f\n", a, b, divide(a, b))
	fmt.Printf("%d ^ %d = %d\n", a, b, power(a, b))

	// Factorial
	fmt.Printf("Factorial of %d = %d\n", n, factorial(n))

	// String operations
	fmt.Printf("Reverse of '%s' = '%s'\n", text, reverseString(text))
	fmt.Printf("Vowels in '%s' = %d\n", text, countVowels(text))

	// Array operations

	fmt.Printf("Sum of array = %d\n", sumArray(numbers))
	fmt.Printf("Max value in array = %d\n", findMax(numbers))
	fmt.Printf("Min value in array = %d\n", findMin(numbers))

	// Fibonacci

	fmt.Printf("Fibonacci(%d) = %d\n", fibN, fibonacci(fibN))

	// Calculator
	fmt.Printf("Calculator(10, 5, '+') = %.2f\n", calculator(10, 5, "+"))

	// Current time
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))

	fmt.Println("Simple Go Functions Example")
	fmt.Println("---------------------------")

	// Arithmetic operations

	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, subtract(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, multiply(a, b))
	fmt.Printf("%d / %d = %.2f\n", a, b, divide(a, b))
	fmt.Printf("%d ^ %d = %d\n", a, b, power(a, b))

	// Factorial
	fmt.Printf("Factorial of %d = %d\n", n, factorial(n))

	// String operations
	fmt.Printf("Reverse of '%s' = '%s'\n", text, reverseString(text))
	fmt.Printf("Vowels in '%s' = %d\n", text, countVowels(text))

	// Array operations

	fmt.Printf("Sum of array = %d\n", sumArray(numbers))
	fmt.Printf("Max value in array = %d\n", findMax(numbers))
	fmt.Printf("Min value in array = %d\n", findMin(numbers))

	// Fibonacci

	fmt.Printf("Fibonacci(%d) = %d\n", fibN, fibonacci(fibN))

	// Calculator
	fmt.Printf("Calculator(10, 5, '+') = %.2f\n", calculator(10, 5, "+"))

	// Current time
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))
	fmt.Println("Simple Go Functions Example")
	fmt.Println("---------------------------")

	// Arithmetic operations

	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, subtract(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, multiply(a, b))
	fmt.Printf("%d / %d = %.2f\n", a, b, divide(a, b))
	fmt.Printf("%d ^ %d = %d\n", a, b, power(a, b))

	// Factorial
	fmt.Printf("Factorial of %d = %d\n", n, factorial(n))

	// String operations
	fmt.Printf("Reverse of '%s' = '%s'\n", text, reverseString(text))
	fmt.Printf("Vowels in '%s' = %d\n", text, countVowels(text))

	// Array operations

	fmt.Printf("Sum of array = %d\n", sumArray(numbers))
	fmt.Printf("Max value in array = %d\n", findMax(numbers))
	fmt.Printf("Min value in array = %d\n", findMin(numbers))

	// Fibonacci

	fmt.Printf("Fibonacci(%d) = %d\n", fibN, fibonacci(fibN))

	// Calculator
	fmt.Printf("Calculator(10, 5, '+') = %.2f\n", calculator(10, 5, "+"))

	// Current time
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))
	fmt.Println("Simple Go Functions Example")
	fmt.Println("---------------------------")

	// Arithmetic operations

	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, subtract(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, multiply(a, b))
	fmt.Printf("%d / %d = %.2f\n", a, b, divide(a, b))
	fmt.Printf("%d ^ %d = %d\n", a, b, power(a, b))

	// Factorial
	fmt.Printf("Factorial of %d = %d\n", n, factorial(n))

	// String operations
	fmt.Printf("Reverse of '%s' = '%s'\n", text, reverseString(text))
	fmt.Printf("Vowels in '%s' = %d\n", text, countVowels(text))

	// Array operations

	fmt.Printf("Sum of array = %d\n", sumArray(numbers))
	fmt.Printf("Max value in array = %d\n", findMax(numbers))
	fmt.Printf("Min value in array = %d\n", findMin(numbers))

	// Fibonacci

	fmt.Printf("Fibonacci(%d) = %d\n", fibN, fibonacci(fibN))

	// Calculator
	fmt.Printf("Calculator(10, 5, '+') = %.2f\n", calculator(10, 5, "+"))

	// Current time
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))
	fmt.Println("Simple Go Functions Example")
	fmt.Println("---------------------------")

	// Arithmetic operations

	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, subtract(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, multiply(a, b))
	fmt.Printf("%d / %d = %.2f\n", a, b, divide(a, b))
	fmt.Printf("%d ^ %d = %d\n", a, b, power(a, b))

	// Factorial
	fmt.Printf("Factorial of %d = %d\n", n, factorial(n))

	// String operations
	fmt.Printf("Reverse of '%s' = '%s'\n", text, reverseString(text))
	fmt.Printf("Vowels in '%s' = %d\n", text, countVowels(text))

	// Array operations

	fmt.Printf("Sum of array = %d\n", sumArray(numbers))
	fmt.Printf("Max value in array = %d\n", findMax(numbers))
	fmt.Printf("Min value in array = %d\n", findMin(numbers))

	// Fibonacci

	fmt.Printf("Fibonacci(%d) = %d\n", fibN, fibonacci(fibN))

	// Calculator
	fmt.Printf("Calculator(10, 5, '+') = %.2f\n", calculator(10, 5, "+"))

	// Current time
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))

	fmt.Println("Simple Go Functions Example")
	fmt.Println("---------------------------")

	// Arithmetic operations

	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, subtract(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, multiply(a, b))
	fmt.Printf("%d / %d = %.2f\n", a, b, divide(a, b))
	fmt.Printf("%d ^ %d = %d\n", a, b, power(a, b))

	// Factorial
	fmt.Printf("Factorial of %d = %d\n", n, factorial(n))

	// String operations
	fmt.Printf("Reverse of '%s' = '%s'\n", text, reverseString(text))
	fmt.Printf("Vowels in '%s' = %d\n", text, countVowels(text))

	// Array operations

	fmt.Printf("Sum of array = %d\n", sumArray(numbers))
	fmt.Printf("Max value in array = %d\n", findMax(numbers))
	fmt.Printf("Min value in array = %d\n", findMin(numbers))

	// Fibonacci

	fmt.Printf("Fibonacci(%d) = %d\n", fibN, fibonacci(fibN))

	// Calculator
	fmt.Printf("Calculator(10, 5, '+') = %.2f\n", calculator(10, 5, "+"))

	// Current time
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))
	fmt.Println("Simple Go Functions Example")
	fmt.Println("---------------------------")

	// Arithmetic operations

	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, subtract(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, multiply(a, b))
	fmt.Printf("%d / %d = %.2f\n", a, b, divide(a, b))
	fmt.Printf("%d ^ %d = %d\n", a, b, power(a, b))

	// Factorial
	fmt.Printf("Factorial of %d = %d\n", n, factorial(n))

	// String operations
	fmt.Printf("Reverse of '%s' = '%s'\n", text, reverseString(text))
	fmt.Printf("Vowels in '%s' = %d\n", text, countVowels(text))

	// Array operations

	fmt.Printf("Sum of array = %d\n", sumArray(numbers))
	fmt.Printf("Max value in array = %d\n", findMax(numbers))
	fmt.Printf("Min value in array = %d\n", findMin(numbers))

	// Fibonacci

	fmt.Printf("Fibonacci(%d) = %d\n", fibN, fibonacci(fibN))

	// Calculator
	fmt.Printf("Calculator(10, 5, '+') = %.2f\n", calculator(10, 5, "+"))

	// Current time
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))
	fmt.Println("Simple Go Functions Example")
	fmt.Println("---------------------------")

	// Arithmetic operations

	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, subtract(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, multiply(a, b))
	fmt.Printf("%d / %d = %.2f\n", a, b, divide(a, b))
	fmt.Printf("%d ^ %d = %d\n", a, b, power(a, b))

	// Factorial
	fmt.Printf("Factorial of %d = %d\n", n, factorial(n))

	// String operations
	fmt.Printf("Reverse of '%s' = '%s'\n", text, reverseString(text))
	fmt.Printf("Vowels in '%s' = %d\n", text, countVowels(text))

	// Array operations

	fmt.Printf("Sum of array = %d\n", sumArray(numbers))
	fmt.Printf("Max value in array = %d\n", findMax(numbers))
	fmt.Printf("Min value in array = %d\n", findMin(numbers))

	// Fibonacci

	fmt.Printf("Fibonacci(%d) = %d\n", fibN, fibonacci(fibN))

	// Calculator
	fmt.Printf("Calculator(10, 5, '+') = %.2f\n", calculator(10, 5, "+"))

	// Current time
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))
	fmt.Println("Simple Go Functions Example")
	fmt.Println("---------------------------")

	// Arithmetic operations

	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, subtract(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, multiply(a, b))
	fmt.Printf("%d / %d = %.2f\n", a, b, divide(a, b))
	fmt.Printf("%d ^ %d = %d\n", a, b, power(a, b))

	// Factorial
	fmt.Printf("Factorial of %d = %d\n", n, factorial(n))

	// String operations
	fmt.Printf("Reverse of '%s' = '%s'\n", text, reverseString(text))
	fmt.Printf("Vowels in '%s' = %d\n", text, countVowels(text))

	// Array operations

	fmt.Printf("Sum of array = %d\n", sumArray(numbers))
	fmt.Printf("Max value in array = %d\n", findMax(numbers))
	fmt.Printf("Min value in array = %d\n", findMin(numbers))

	// Fibonacci

	fmt.Printf("Fibonacci(%d) = %d\n", fibN, fibonacci(fibN))

	// Calculator
	fmt.Printf("Calculator(10, 5, '+') = %.2f\n", calculator(10, 5, "+"))

	// Current time
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))

	fmt.Println("Simple Go Functions Example")
	fmt.Println("---------------------------")

	// Arithmetic operations

	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, subtract(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, multiply(a, b))
	fmt.Printf("%d / %d = %.2f\n", a, b, divide(a, b))
	fmt.Printf("%d ^ %d = %d\n", a, b, power(a, b))

	// Factorial
	fmt.Printf("Factorial of %d = %d\n", n, factorial(n))

	// String operations
	fmt.Printf("Reverse of '%s' = '%s'\n", text, reverseString(text))
	fmt.Printf("Vowels in '%s' = %d\n", text, countVowels(text))

	// Array operations

	fmt.Printf("Sum of array = %d\n", sumArray(numbers))
	fmt.Printf("Max value in array = %d\n", findMax(numbers))
	fmt.Printf("Min value in array = %d\n", findMin(numbers))

	// Fibonacci

	fmt.Printf("Fibonacci(%d) = %d\n", fibN, fibonacci(fibN))

	// Calculator
	fmt.Printf("Calculator(10, 5, '+') = %.2f\n", calculator(10, 5, "+"))

	// Current time
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))
	fmt.Println("Simple Go Functions Example")
	fmt.Println("---------------------------")

	// Arithmetic operations

	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, subtract(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, multiply(a, b))
	fmt.Printf("%d / %d = %.2f\n", a, b, divide(a, b))
	fmt.Printf("%d ^ %d = %d\n", a, b, power(a, b))

	// Factorial
	fmt.Printf("Factorial of %d = %d\n", n, factorial(n))

	// String operations
	fmt.Printf("Reverse of '%s' = '%s'\n", text, reverseString(text))
	fmt.Printf("Vowels in '%s' = %d\n", text, countVowels(text))

	// Array operations

	fmt.Printf("Sum of array = %d\n", sumArray(numbers))
	fmt.Printf("Max value in array = %d\n", findMax(numbers))
	fmt.Printf("Min value in array = %d\n", findMin(numbers))

	// Fibonacci

	fmt.Printf("Fibonacci(%d) = %d\n", fibN, fibonacci(fibN))

	// Calculator
	fmt.Printf("Calculator(10, 5, '+') = %.2f\n", calculator(10, 5, "+"))

	// Current time
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))
	fmt.Println("Simple Go Functions Example")
	fmt.Println("---------------------------")

	// Arithmetic operations

	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, subtract(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, multiply(a, b))
	fmt.Printf("%d / %d = %.2f\n", a, b, divide(a, b))
	fmt.Printf("%d ^ %d = %d\n", a, b, power(a, b))

	// Factorial
	fmt.Printf("Factorial of %d = %d\n", n, factorial(n))

	// String operations
	fmt.Printf("Reverse of '%s' = '%s'\n", text, reverseString(text))
	fmt.Printf("Vowels in '%s' = %d\n", text, countVowels(text))

	// Array operations

	fmt.Printf("Sum of array = %d\n", sumArray(numbers))
	fmt.Printf("Max value in array = %d\n", findMax(numbers))
	fmt.Printf("Min value in array = %d\n", findMin(numbers))

	// Fibonacci

	fmt.Printf("Fibonacci(%d) = %d\n", fibN, fibonacci(fibN))

	// Calculator
	fmt.Printf("Calculator(10, 5, '+') = %.2f\n", calculator(10, 5, "+"))

	// Current time
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))
}
