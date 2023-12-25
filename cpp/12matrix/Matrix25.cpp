#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix25");
	
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
	double sum, maxSum;
	for (int i = 0; i < M; i++)
	{
		sum = 0;
		for (int j = 0; j < N; j++)
			sum += matr[i][j];

		if (index < 0)
		{
			maxSum = sum;
			index = i;
		}
		
		else if (sum > maxSum)
		{
			maxSum = sum;
			index = i;
		}
	}
	cout << index+1;
	cout << maxSum;

	for (int i = 0; i < M; i++)
		delete [] matr[i];
	delete [] matr;
	matr = 0;

	return 0;
}
