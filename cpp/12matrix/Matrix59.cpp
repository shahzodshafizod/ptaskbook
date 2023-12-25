#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix59");
	
	int M, N;
	cin >> M >> N;

	double** matr = new double* [M];
	for (int i = 0; i < M; i++)
	{
		matr[i] = new double [N];
		for (int j = 0; j < N; j++)
			cin >> matr[i][j];
	}

	double temp;
	for (int i = 0; i < M/2; i++)
	{
		for (int j = 0; j < N; j++)
		{
			temp = matr[i][j];
			matr[i][j] = matr[M-1-i][j];
			matr[M-1-i][j] = temp;
		}
	}

	for (int i = 0; i < M; i++)
		for (int j = 0; j < N; j++)
			cout << matr[i][j];

	for (int i = 0; i < M; i++)
		delete [] matr[i];
	delete [] matr;
	matr = 0;
	
	return 0;
}
