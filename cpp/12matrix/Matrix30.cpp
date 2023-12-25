#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix30");
	
	int M, N;
	cin >> M >> N;

	double* mass = new double [N];
	memset(mass, 0, N*sizeof(double));

	double** matr = new double* [M];
	for (int i = 0; i < M; i++)
	{
		matr[i] = new double [N];
		
		for (int j = 0; j < N; j++)
		{
			cin >> matr[i][j];
			mass[j] += matr[i][j];
		}
	}

	int count;
	for (int j = 0; j < N; j++)
	{
		mass[j] /= M;

		int count = 0;
		for (int i = 0; i < M; i++)
			if (matr[i][j] > mass[j])
				count++;

		cout << count;
	}

	for (int i = 0; i < M; i++)
		delete [] matr[i];
	delete [] mass;
	delete [] matr;
	mass = 0;
	matr = 0;

	return 0;
}
