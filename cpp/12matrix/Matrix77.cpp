#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix77");
	
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
	for (int j = 0; j < N-1; j++)
	{
		for (int c = 1; c < N-j; c++)
		{
			if (matr[M-1][c-1] < matr[M-1][c])
			{
				for (int i = 0; i < M; i++)
				{
					temp = matr[i][c-1];
					matr[i][c-1] = matr[i][c];
					matr[i][c] = temp;
				}
			}
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
