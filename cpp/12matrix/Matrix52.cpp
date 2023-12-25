#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix52");
	
	int M, N;
	cin >> M >> N;

	double** matr = new double* [M];
	double maximal, minimal;
	int minCol = -1, maxCol = -1;
	for(int i = 0; i < M; i++)
	{
		matr[i] = new double [N];
		for (int j = 0; j < N; j++)
		{
			cin >> matr[i][j];

			if ( (minCol < 0) || (maxCol < 0) )
			{
				minimal = maximal = matr[i][j];
				minCol = maxCol = j;
			}
			else if (matr[i][j] < minimal)
			{
				minimal = matr[i][j];
				minCol = j;
			}
			else if (matr[i][j] > maximal)
			{
				maximal = matr[i][j];
				maxCol = j;
			}
		}
	}

	if (maxCol != minCol)
	{
		double temp;
		for (int i = 0; i < M; i++)
		{
			temp = matr[i][minCol];
			matr[i][minCol] = matr[i][maxCol];
			matr[i][maxCol] = temp;
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
