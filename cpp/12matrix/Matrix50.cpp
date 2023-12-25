#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix50");
	
	int M, N;
	cin >> M >> N;

	double** matr = new double* [M];
	for (int i = 0; i < M; i++)
	{
		matr[i] = new double [N];
		for (int j = 0; j < N; j++)
			cin >> matr[i][j];
	}

	double minimal, maximal, temp;
	int minIndex, maxIndex;
	for (int j = 0; j < N; j++)
	{
		minimal = maximal = matr[0][j];
		minIndex = maxIndex = 0;
		for (int i = 1; i < M; i++)
		{
			if (matr[i][j] > maximal)
			{
				maximal = matr[i][j];
				maxIndex = i;
			}

			if (matr[i][j] < minimal)
			{
				minimal = matr[i][j];
				minIndex = i;
			}
		}

		temp = matr[minIndex][j];
		matr[minIndex][j] = matr[maxIndex][j];
		matr[maxIndex][j] = temp;
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
