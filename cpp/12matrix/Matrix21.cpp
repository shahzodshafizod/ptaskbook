#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix21");
	
	int M, N;
	cin >> M >> N;

	double** matr = new double* [M];
	for (int i = 0; i < M; i++)
	{
		matr[i] = new double [N];
		for (int j = 0; j < N; j++)
			cin >> matr[i][j];
	}

	double sum;
	for (int i = 0; i < M; i += 2)
	{
		sum = 0;
		for (int j = 0; j < N; j++)
			sum += matr[i][j];
		cout << sum / N;
	}

	for (int i = 0; i < M; i++)
		delete [] matr[i];
	delete [] matr;
	matr = 0;

	return 0;
}
