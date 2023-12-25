#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix81");
	
	int M;
	cin >> M;

	double** A = new double* [M];
	for (int i = 0; i < M; i++)
	{
		A[i] = new double [M];
		for (int j = 0; j < M; j++)
			cin >> A[i][j];
	}

	double AMean = 0;
	for (int i = 0; i < M; i++)
	{
		AMean += A[i][M-1-i];
	}
	AMean /= M;

	cout << AMean;

	for (int i = 0; i < M; i++)
		delete [] A[i];
	delete [] A;
	A = 0;
	
	return 0;
}
