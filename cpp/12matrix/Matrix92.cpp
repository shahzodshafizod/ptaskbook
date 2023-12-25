#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix92");
	
	int M;
	cin >> M;

	double** matr = new double* [M];
	for (int i = 0; i < M; i++)
	{
		matr[i] = new double [M];
		for (int j = 0; j < M; j++)
			cin >> matr[i][j];
	}

	for (int i = 0; i < M/2; i++)
		for (int j = i+1; j < M-1-i; j++)
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
