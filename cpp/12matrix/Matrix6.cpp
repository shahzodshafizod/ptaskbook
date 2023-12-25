#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix6");
	
	int M, N;
	cin >> M >> N;

	double D;
	cin >> D;

	double* mass = new double [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	double** matr = new double* [M];
	for (int i = 0; i < M; i++)
		matr[i] = new double [N];

	for (int j = 0; j < N; j++)
	{
		matr[0][j] = mass[j];
		for (int i = 1; i < M; i++)
			matr[i][j] = matr[i-1][j] * D;
	}

	for (int i = 0; i < M; i++)
		for (int j = 0; j < N; j++)
			cout << matr[i][j];

	for (int i = 0; i < M; i++)
		delete [] matr[i];

	delete [] mass;
	delete [] matr;
	mass = 0;
	matr = 0;
	
	return 0;
}
