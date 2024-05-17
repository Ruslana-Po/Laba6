package main

import (
"fmt"
"math/rand"
)

const (
m = 8
n = 7
)

func main() {
// Random number generation
randInt := func(min, max int) int {
return rand.Intn(max-min+1) + min
}

// Task 1
fmt.Println("\tTask 1:")
arr1 := make([][]int, m)
for i := range arr1 {
arr1[i] = make([]int, n)
for j := range arr1[i] {
arr1[i][j] = randInt(-9, 9)
fmt.Print(arr1[i][j], " ")
}
fmt.Println()
}

maxNegativeCount := 0
maxNegativeRow := 0
for i := range arr1 {
negativeCount := 0
for j := range arr1[i] {
if arr1[i][j] < 0 {
negativeCount++
}
}
if negativeCount > maxNegativeCount {
maxNegativeCount = negativeCount
maxNegativeRow = i
}
}

fmt.Printf("Most negative numbers in row %d\n", maxNegativeRow)
for _, v := range arr1[maxNegativeRow] {
fmt.Print(v, " ")
}
fmt.Println()

// Task 2
fmt.Println("\tTask 2:")
arr2 := make([][]int, m)
for i := range arr2 {
arr2[i] = make([]int, m)
for j := range arr2[i] {
arr2[i][j] = randInt(10, 100)
fmt.Print(arr2[i][j], " ")
}
fmt.Println()
}

// Sort the diagonals in descending order
for i := range arr2 {
for j := i + 1; j < m; j++ {
if arr2[i][i] < arr2[j][j] {
arr2[i][i], arr2[j][j] = arr2[j][j], arr2[i][i]
}
if arr2[i][m-i-1] < arr2[j][m-j-1] {
arr2[i][m-i-1], arr2[j][m-j-1] = arr2[j][m-j-1], arr2[i][m-i-1]
}
}
}

fmt.Println("Sorted:")
for i := range arr2 {
for j := range arr2[i] {
fmt.Print(arr2[i][j], " ")
}
fmt.Println()
}

sum1, sum2 := 0, 0
for i := range arr2 {
sum1 += arr2[i][i]
sum2 += arr2[i][m-i-1]
}

if sum1 > sum2 {
fmt.Printf("Sum of first diagonal %d is greater than sum of second diagonal %d\n", sum1, sum2)
} else {
fmt.Printf("Sum of second diagonal %d is greater than sum of first diagonal %d\n", sum2, sum1)
}

// Task 3
fmt.Println("\tTask 3:")
arr3 := [][]byte{
{'.', 'X', 'X', 'X', '.', '.', '.', '.', '.', 'X'},
{'.', '.', '.', '.', '.', 'X', '.', '.', '.', 'X'},
{'.', '.', '.', '.', '.', 'X', '.', '.', '.', '.'},
{'.', '.', '.', '.', '.', 'X', '.', '.', 'X', '.'},
{'X', 'X', 'X', '.', '.', 'X', '.', '.', 'X', '.'},
{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
{'.', '.', '.', '.', 'X', '.', '.', '.', '.', '.'},
{'.', '.', '.', '.', '.', '.', 'X', '.', '.', '.'},
{'X', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
{'.', '.', '.', 'X', 'X', '.', '.', 'X', '.', '.'},
}

for i := range arr3 {
for j := range arr3[i] {
fmt.Print(string(arr3[i][j]), " ")
}
fmt.Println()
}

count := 0
for i := range arr3 {
for j := range arr3[i] {
if arr3[i][j] == 'X' {
if (i == 0 || arr3[i-1][j] == '.') && (j == 0 || arr3[i][j-1] == '.') {
count++
}
}
}
}

fmt.Println("Number of ships:", count)
}
