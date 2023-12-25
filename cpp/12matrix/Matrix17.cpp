#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix17");
	
	int M, N;
	cin >> M >> N;

	double** matr = new double* [M];
	for (int i = 0; i < M; i++)
	{
		matr[i] = new double [N];
		for (int j = 0; j < N; j++)
			cin >> matr[i][j];
	}

	int K;
	cin >> K;

	double sum = 0, pro = 1;
	for (int j = 0; j < N; j++)
	{
		sum += matr[K-1][j];
		pro *= matr[K-1][j];
	}
	cout << sum << pro;

	for (int i = 0; i < M; i++)
		delete [] matr[i];
	delete [] matr;
	matr = 0;

	return 0;
}
