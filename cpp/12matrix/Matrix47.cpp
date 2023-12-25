#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix47");
	
	int M, N;
	cin >> M >> N;

	double** matr = new double* [M];
	for (int i = 0; i < M; i++)
	{
		matr[i] = new double [N];
		for (int j = 0; j < N; j++)
			cin >> matr[i][j];
	}

	int K1, K2;
	cin >> K1 >> K2;
	double temp;
	for (int j = 0; j < N; j++)
	{
		temp = matr[K1-1][j];
		matr[K1-1][j] = matr[K2-1][j];
		matr[K2-1][j] = temp;
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
