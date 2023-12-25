#include <iostream>
using namespace std;

void Hill(double A[], int N);

int main()
{
	//Task("Param13");
	int N;
	double* mass = NULL;

	for (int i = 0; i < 3; i++)
	{
		cin >> N;
		mass = new double [N];
		for (int j = 0; j < N; j++)
			cin >> mass[j];
		
		Hill(mass, N);
		for (int j = 0; j < N; j++)
			cout << mass[j];

		delete [] mass;
		mass = NULL;
	}
	
	return 0;
}

void Hill(double A[], int N)
{
	double temp;
	for (int i = 0; i < N/2; i++)
	{
		for (int j = N-1-i; j > i; j--)
			if (A[j] < A[j-1])
			{
				temp = A[j];
				A[j] = A[j-1];
				A[j-1] = temp;
			}
		
		for (int j = i+2; j < N-i; j++)
			if (A[j] > A[j-1])
			{
				temp = A[j];
				A[j] = A[j-1];
				A[j-1] = temp;
			}
	}
}