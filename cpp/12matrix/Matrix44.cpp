#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix44");
	
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
	double minimal = 0;
	for (int i = 0; i < M; i++)
	{
		progress = regress = true;
		for (int j = 1; j < N; j++)
		{
			if (matr[i][j-1] > matr[i][j])
				progress = false;
			if (matr[i][j-1] < matr[i][j])
				regress = false;
		}

		if (progress)
		{
			if (initialized == false)
			{
				minimal = matr[i][0];
				initialized = true;
			}
			else if (matr[i][0] < minimal)
				minimal = matr[i][0];
		}
		else if (regress)
		{
			if (initialized == false)
			{
				minimal = matr[i][N-1];
				initialized = true;
			}
			else if (matr[i][N-1] < minimal)
				minimal = matr[i][N-1];
		}
	}
	cout << minimal;

	for (int i = 0; i < M; i++)
		delete [] matr[i];
	delete [] matr;
	matr = 0;

	return 0;
}
