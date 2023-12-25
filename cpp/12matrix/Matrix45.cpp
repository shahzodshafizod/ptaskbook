#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix45");
	
	int M, N;
	cin >> M >> N;

	double** matr = new double* [M];
	for (int i = 0; i < M; i++)
	{
		matr[i] = new double [N];
		for (int j = 0; j < N; j++)
			cin >> matr[i][j];
	}

	bool progress, regress, initialized = false;
	double maximal = 0;
	for (int j = 0; j < N; j++)
	{
		progress = regress = true;
		for (int i = 1; i < M; i++)
		{
			if (matr[i-1][j] < matr[i][j])
				regress = false;
			
			if (matr[i-1][j] > matr[i][j])
				progress = false;
		}

		if (progress)
		{
			if (initialized == false)
			{
				maximal = matr[M-1][j];
				initialized = true;
			}
			else if (matr[M-1][j] > maximal)
			{
				maximal = matr[M-1][j];
			}
		}
		else if (regress)
		{
			if (initialized == false)
			{
				maximal = matr[0][j];
				initialized = true;
			}
			else if (matr[0][j] > maximal)
				maximal = matr[0][j];
		}
	}
	cout << maximal;

	for (int i = 0; i < M; i++)
		delete [] matr[i];
	delete [] matr;
	matr = 0;

	return 0;
}
