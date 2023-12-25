#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix3");
	
	int M, N;
	cin >> M >> N;

	double* mass = new double [M];
	double** matr = new double* [M];
	for (int i = 0; i < M; i++)
	{
		cin >> mass[i];
		matr[i] = new double [N];
		for (int j = 0; j < N; j++)
			matr[i][j] = mass[i];
	}

	for (int i = 0; i < M; i++)
		for (int j = 0; j < N; j++)
			cout << matr[i][j];

	for (int i = 0; i < M; i++)
		delete [] matr[i];

	delete [] matr;
	delete [] mass;
	matr = 0;
	mass = 0;
	
	return 0;
}
