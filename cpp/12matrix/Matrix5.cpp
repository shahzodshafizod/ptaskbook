#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix5");
	
	int M, N;
	cin >> M >> N;

	double D;
	cin >> D;

	double* mass = new double [M];
	for (int i = 0; i < M; i++)
		cin >> mass[i];

	double** matr = new double* [M];
	for (int i = 0; i < M; i++)
	{
		matr[i] = new double [N];
		matr[i][0] = mass[i];
		for (int j = 1; j < N; j++)
			matr[i][j] = matr[i][j-1] + D;
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
