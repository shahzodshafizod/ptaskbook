#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix12");
	
	int M, N;
	cin >> M >> N;

	double** matr = new double* [M];
	for (int i = 0; i < M; i++)
	{
		matr[i] = new double [N];
		for (int j = 0; j < N; j++)
			cin >> matr[i][j];
	}

	for (int j = 0; j < N; j++)
	{
		if (j % 2 == 0)
			for (int i = 0; i < M; i++)
				cout << matr[i][j];
		else
			for (int i = M-1; i >= 0; i--)
				cout << matr[i][j];
	}

	for (int i = 0; i < M; i++)
		delete [] matr[i];

	delete [] matr;
	matr = 0;

	return 0;
}
