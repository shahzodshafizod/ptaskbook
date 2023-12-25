#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix49");
	
	int M, N;
	cin >> M >> N;

	double** matr = new double* [M];
	double minimal, maximal, temp;
	int minIndex, maxIndex;
	for (int i = 0; i < M; i++)
	{
		matr[i] = new double [N];
		cin >> matr[i][0];
		minimal = maximal = matr[i][0];
		minIndex = maxIndex = 0;
		for (int j = 1; j < N; j++)
		{
			cin >> matr[i][j];

			if (matr[i][j] > maximal)
			{
				maximal = matr[i][j];
				maxIndex = j;
			}

			if (matr[i][j] < minimal)
			{
				minimal = matr[i][j];
				minIndex = j;
			}
		}

		temp = matr[i][minIndex];
		matr[i][minIndex] = matr[i][maxIndex];
		matr[i][maxIndex] = temp;
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
