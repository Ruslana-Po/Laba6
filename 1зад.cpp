#include <iostream>
#include <vector>
#include <Windows.h>
#include <random>
using namespace std;
const int m = 8;
const int n = 7;
int main() {
	SetConsoleCP(1251);
	SetConsoleOutputCP(1251);
	//рандом 
	random_device ran;
	uniform_int_distribution<> dist(-9, 9);
	//1 задание
	int arr1[m][n];
	cout << endl << "	Задание1:" << endl;
	//заполняем массив -10 до 10
	for (int i = 0; i < m; i++) {
		for (int j = 0; j < n; j++) {
			arr1[i][j] = dist(ran);
			cout << arr1[i][j]<<" ";
		}
		cout << endl;
	}
	int max = 0;
	int number = 0;
	//поиск строки с наиб кол-вом отрицательных чисел
	for (int i = 0; i < m; i++) {
		int kol = 0;
		for (int j = 0; j < n; j++) {
			if (arr1[i][j] < 0) {
				kol++;
			}
		}
		if (max < kol) {
			max = kol;
			number = i;
		}
	}
	vector<int> str;
	cout << " Наибольшее кол-во отриц чисел в " << number << " строоке" << endl;
	for (int i = 0; i < n; i++) {
		str.push_back(arr1[number][i]);
	}
	for (int i : str) {
		cout << i << " ";
	}
	//задание 2
	cout << endl << "	Задание2:" << endl;
	int arr2[m][m];
	random_device ran2;
	uniform_int_distribution<> dist2(10,100);
	//заполняем массив 10 до 100
	for (int i = 0; i < m; i++) {
		for (int j = 0; j < m; j++) {
			arr2[i][j] = dist2(ran2);
			cout << arr2[i][j] << " ";
		}
		cout << endl;
	}
	//сортировка по убыванию
	//1 диагональ
	for (int i = 0; i < m; ++i) {
		for (int j = i + 1; j < m; ++j) {
			if (arr2[i][i] < arr2[j][j]) {
				int temp = arr2[i][i];
				arr2[i][i] = arr2[j][j];
				arr2[j][j] = temp;
			}
		}
	}
	//2
	for (int i = 0; i < m; ++i) {
		for (int j = i + 1; j < m; ++j) {
			if (arr2[i][m-i-1] < arr2[j][m-j-1]) {
				int temp = arr2[i][m-i-1];
				arr2[i][m-i-1] = arr2[j][m-j-1];
				arr2[j][m-j-1] = temp;
			}
		}
	}
	cout << "Отсортированный: " << endl;
	for (int i = 0; i < m; i++) {
		for (int j = 0; j < m; j++) {
			cout << arr2[i][j] << " ";
		}
		cout << endl;
	}
	//Cумма
	int sum1 = 0;
	int sum2 = 0;
	for (int i = 0; i < m; i++) {
		sum1 += arr2[i][i];
		sum2 += arr2[i][m - i - 1];
	}
	if (sum1 > sum2) {
		cout << "Сумма первой " << sum1 << " больше суммы второй " << sum2 << endl;
	}
	else {
		cout << "Сумма второй " << sum2 << " больше суммы первой " << sum1 << endl;
	}
	//задание 3
	char arr3[10][10]= {
		{'.', 'X', 'X', 'X','.','.','.','.','.','X'},
		{'.', '.', '.', '.','.','X','.','.','.','X'},
		{'.', '.', '.', '.','.','X','.','.','.','.'},
		{'.', '.', '.', '.','.','X','.','.','X','.'},
		{'X', 'X', 'X', '.','.','X','.','.','X','.'},
		{'.', '.', '.', '.','.','.','.','.','.','.'},
		{'.', '.', '.', '.','X','.','.','.','.','.'},
		{'.', '.', '.', '.','.','.','X','.','.','.'},
		{'X', '.', '.', '.','.','.','.','.','.','.'},
		{'.', '.', '.', 'X','X','.','.','X','.','.'},
	};
	// массив вывод
	for (int i = 0; i < 10; i++) {
		for (int j = 0; j < 10; j++) {
			cout << arr3[i][j] << " ";
		}
		cout << endl;
	}
	int count = 0;
	for (int i = 0; i < 10; ++i) {
		for (int j = 0; j < 10; ++j) {
			if (arr3[i][j] == 'X') {
				if ((i == 0 || arr3[i - 1][j] == '.') && (j == 0 || arr3[i][j - 1] == '.')) {
					count++;
				}
			}
		}
	}
	cout << "Кол-во кораблей: " << count;
}
