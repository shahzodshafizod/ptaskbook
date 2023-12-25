#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix23");
	
	int M, N;
	cin >> M >> N;

	double** matr = new double* [M];
	for (int i = 0; i < M; i++)
	{
		matr[i] = new double [N];
		for (int j = 0; j < N; j++)
			cin >> matr[i][j];
	}

	double minimal;
	for (int i = 0; i < M; i++)
	{
		minimal = matr[i][0];
		for (int j = 1; j < N; j++)
			if (matr[i][j] < minimal)
				minimal = matr[i][j];
		cout << minimal;
	}

	for (int i = 0; i < M; i++)
		delete [] matr[i];
	delete [] matr;
	matr = 0;

	return 0;
}
