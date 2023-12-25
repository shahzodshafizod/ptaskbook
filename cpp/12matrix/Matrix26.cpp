#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix26");
	
	int M, N;
	cin >> M >> N;

	double** matr = new double* [M];
	for (int i = 0; i < M; i++)
	{
		matr[i] = new double [N];
		for (int j = 0; j < N; j++)
			cin >> matr[i][j];
	}

	int index = -1;
	double pro, minPro;
	for (int j = 0; j < N; j++)
	{
		pro = 1;
		for (int i = 0; i < M; i++)
			pro *= matr[i][j];

		if (index < 0)
		{
			minPro = pro;
			index = j;
		}
		else if (pro < minPro)
		{
			minPro = pro;
			index = j;
		}
	}
	cout << index+1;
	cout << minPro;

	for (int i = 0; i < M; i++)
		delete [] matr[i];
	delete [] matr;
	matr = 0;

	return 0;
}
