#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix76");
	
	int M, N;
	cin >> M >> N;

	double** matr = new double* [M];
	for (int i = 0; i < M; i++)
	{
		matr[i] = new double [N];
		for (int j = 0; j < N; j++)
			cin >> matr[i][j];
	}

	double* temp = 0;
	for (int i = 0; i < M-1; i++)
	{
		for (int r = 1; r < M-i; r++)
		{
			if (matr[r-1][0] > matr[r][0])
			{
				temp = matr[r-1];
				matr[r-1] = matr[r];
				matr[r] = temp;
			}
		}
	}
	temp = 0;

	for (int i = 0; i < M; i++)
		for (int j = 0; j < N; j++)
			cout << matr[i][j];
	
	for (int i = 0; i < M; i++)
		delete [] matr[i];
	delete [] matr;
	matr = 0;
	
	return 0;
}
