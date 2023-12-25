#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix80");
	
	int M;
	cin >> M;

	double** A = new double* [M];
	for (int i = 0; i < M; i++)
	{
		A[i] = new double [M];
		for (int j = 0; j < M; j++)
			cin >> A[i][j];
	}

	double sum = 0;
	for (int i = 0; i < M; i++)
		sum += A[i][i];

	cout << sum;

	for (int i = 0; i < M; i++)
		delete [] A[i];
	delete [] A;
	A = 0;
	
	return 0;
}
