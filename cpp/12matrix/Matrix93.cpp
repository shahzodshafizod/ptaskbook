#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix93");
	
	int M;
	cin >> M;

	double** matr = new double* [M];
	for (int i = 0; i < M; i++)
	{
		matr[i] = new double [M];
		for (int j = 0; j < M; j++)
			cin >> matr[i][j];
	}

	for (int j = M-1, k = 1; j > M/2; j--, k++)
		for (int i = k; i < M-k; i++)
			matr[i][j] = 0;

	for (int i = 0; i < M; i++)
		for (int j = 0; j < M; j++)
			cout << matr[i][j];

	for (int i = 0; i < M; i++)
		delete [] matr[i];
	delete [] matr;
	matr = 0;
	
	return 0;
}
