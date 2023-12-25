#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix11");
	
	int M, N;
	cin >> M >> N;

	double** matr = new double* [M];
	for (int i = 0; i < M; i++)
	{
		matr[i] = new double [N];
		for (int j = 0; j < N; j++)
			cin >> matr[i][j];
	}

	for (int i = 0; i < M; i++)
	{
		if (i % 2 == 0)
			for (int j = 0; j < N; j++)
				cout << matr[i][j];
		else
			for (int j = N-1; j >= 0; j--)
				cout << matr[i][j];
	}

	for (int i = 0; i < M; i++)
		delete [] matr[i];

	delete [] matr;
	matr = 0;

	return 0;
}
