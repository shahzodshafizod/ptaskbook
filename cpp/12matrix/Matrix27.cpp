#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix27");
	
	int M, N;
	cin >> M >> N;

	double** matr = new double* [M];
	for (int i = 0; i < M; i++)
	{
		matr[i] = new double [N];
		for (int j = 0; j < N; j++)
			cin >> matr[i][j];
	}

	double maximal, minimal;
	bool initialized = false;
	for (int i = 0; i < M; i++)
	{
		minimal = matr[i][0];
		for (int j = 1; j < N; j++)
			if (matr[i][j] < minimal)
				minimal = matr[i][j];

		if (initialized == false)
		{
			maximal = minimal;
			initialized = true;
		}
		else if (minimal > maximal)
			maximal = minimal;
	}
	cout << maximal;

	for (int i = 0; i < M; i++)
		delete [] matr[i];
	delete [] matr;
	matr = 0;

	return 0;
}
