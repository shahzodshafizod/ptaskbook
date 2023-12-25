#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix28");
	
	int M, N;
	cin >> M >> N;

	double** matr = new double* [M];
	for (int i = 0; i < M; i++)
	{
		matr[i] = new double [N];
		for (int j = 0; j < N; j++)
			cin >> matr[i][j];
	}

	double minimal, maximal;
	bool initialized = false;
	for (int j = 0; j < N; j++)
	{
		maximal = matr[0][j];
		for (int i = 1; i < M; i++)
			if (matr[i][j] > maximal)
				maximal = matr[i][j];

		if (initialized == false)
		{
			minimal = maximal;
			initialized = true;
		}
		else if (minimal > maximal)
			minimal = maximal;
	}
	cout << minimal;

	for (int i = 0; i < M; i++)
		delete [] matr[i];
	delete [] matr;
	matr = 0;

	return 0;
}
