#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix29");
	
	int M, N;
	cin >> M >> N;

	double** matr = new double* [M];
	double* mass = new double [M];
	for (int i = 0; i < M; i++)
	{
		matr[i] = new double [N];
		mass[i] = 0;
		for (int j = 0; j < N; j++)
		{
			cin >> matr[i][j];
			mass[i] += matr[i][j];
		}
		mass[i] /= N;
	}

	int count;
	for (int i = 0; i < M; i++)
	{
		count = 0;
		for (int j = 0; j < N; j++)
			if (matr[i][j] < mass[i])
				count++;
		cout << count;
	}

	for (int i = 0; i < M; i++)
		delete [] matr[i];
	delete [] matr;
	delete [] mass;
	matr = 0;
	mass = 0;

	return 0;
}
