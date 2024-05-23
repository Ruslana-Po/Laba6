#include<iomanip>
#include <iostream>
#include <vector>
#include <random>
using namespace std;
const float E = 0.001;
const int n = 4;
// четные варианты – методом LU-разложения
int main() {
	setlocale(LC_ALL, "");
	//float M = 1.08, N = 0.22, P= -1.16;
	float A[n][n] = {
		1.08, -0.04, 0.21, -18.00,
		0.25, -1.23, 0.22, -0.09,
		-0.21, 0.22, 0.8, -0.13,
		0.15, -1.31, 0.06, -1.16};
	float L[n][n], U[n][n];
	float B[n] = { -1.24,  -1.16, 2.56, 1.08 };
	float sum=0;
	for (int i = 0; i < n; i++) {
		for (int j = 0; j < n; j++) {
			if (i <= j) {
				L[i][j] = 0;
				if (i == j) {
					L[i][j] = 1;

				}
				for (int k = 0; k < i; k++) {
					sum += (L[i][k] * U[k][j]);
				}
				U[i][j] = (A[i][j] - sum);
			}

			else if (i > j) {
				U[i][j] = 0;
				for (int k = 0; k < j; k++) {
				sum += L[i][k] * U[k][j];
				}
				//U[j][j] была рпечатка
				L[i][j] = (A[i][j] - sum) / U[j][j];
			}
			sum = 0;
		}
	}
	cout << " L: " << endl;
	for (int i = 0; i < n; i++) {
		for (int j = 0; j < n; j++) {
			cout <<fixed << setprecision(3)<< L[i][j] << "  ";
		}
		cout << endl;
	}
	cout <<endl<< " U: " << endl;
	for (int i = 0; i < n; i++) {
		for (int j = 0; j < n; j++) {
			cout << fixed << setprecision(3) << U[i][j] << "  ";
		}
		cout << endl;
	}
	float x[n], y[n];
	for (int k = 0; k < n; k++) {
		for (int j = 0; j < k; j++) {
			sum += L[k][j] * y[j];
		}
		y[k] = B[k] - sum;
		sum = 0;
	}
	for (int k = n-1; k >=0; k--) {
		for (int j = k+1; j < n; j++) {
			sum += U[k][j] * x[j];
		}
		x[k] = (y[k] - sum)/U[k][k];
		sum = 0;
	}
	cout << endl;
	for (int k = 0; k < n; k++) {
		cout <<" X"<<k+1<<" : " << fixed << setprecision(3) << x[k] << "	Y" << k + 1 << " : " << fixed << setprecision(3) << y[k] << endl;
	}
	cout << endl;
	//Метод простой итерации
	float xIter[n],C[n][n], f[n];
	int iter = 0;
	//0
	for (int k = 0; k < n; k++) {
		xIter[k]= 0;
	}
	float xTime[n];
	//1,2...n
	cout << setw(10) << " N	" << setw(20)<<" x1 " << setw(12) << " x2 " << setw(12) << " x3 " << setw(10) << " x4 " << setw(25) <<" En " << endl;
	while (true) {
		for (int i = 0; i < n; i++) {
			for (int j = 0; j < n; j++) {
				if (i == j) {
					C[i][j] = 0;
				}
				else {
					// каноническому виду
					C[i][j] = -A[i][j] / A[i][i];
					sum += C[i][j];
				}
			}
			xTime[i] = xIter[i];
			f[i] = B[i] / A[i][i];
			xIter[i] = f[i] + sum * xIter[i];
			sum = 0;
		}
		//Если выполнено условие −   (k + 1) (k)x x, то процесс завершить иположить х *= х(к + 1), иначе k = k + 1 перейти к пункту 3.
		
		if (abs(xIter[0]- xTime[0])<E) {
			cout << "Сходится "<<endl;
			cout << endl;
			for (int k = 0; k < n; k++) {
				cout << " X" << k + 1 << " : " << fixed << setprecision(3) << xIter[k] << endl;
			}
			break;
		}if (iter > 16) {
			cout << "Расходится. Причины: Диагональное преобладание: Матрица коэффициентов системы линейных уравнений должна быть диагонально доминирующей. Это означает, что модуль элемента на главной диагонали матрицы должен быть больше суммы модулей всех остальных элементов в соответствующей строке (или столбце)." << endl;
			cout << "Норма матрицы меньше 1: Норма матрицы, обратной к матрице коэффициентов системы, должна быть меньше 1. Это обеспечивает сходимость итерационного процесса.";
			cout << endl;
			break;
		}
		iter++;
		cout << setw(10) << iter  << setw(25) << xIter[0]  << setw(15) << xIter[1] << setw(10) << xIter[2]  << setw(10) << xIter[3]  << setw(25) << xIter[0] - xTime[0] << endl;
	}
	
}
