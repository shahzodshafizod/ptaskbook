#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix48");
	
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
	for (int i = 0; i < M; i++)
	{
		temp = matr[i][K1-1];
		matr[i][K1-1] = matr[i][K2-1];
		matr[i][K2-1] = temp;
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
