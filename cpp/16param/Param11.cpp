#include <iostream>
using namespace std;

void SortArray(double A[], int N);

int main()
{
	//Task("Param11");
	int N;
	double* mass = NULL;

	for (int i = 0; i < 3; i++)
	{
		cin >> N;
		mass = new double [N];
		for (int j = 0; j < N; j++)
			cin >> mass[j];
		
		SortArray(mass, N);
		for (int j = 0; j < N; j++)
			cout << mass[j];

		delete [] mass;
		mass = NULL;
	}
	
	return 0;
}

void SortArray(double A[], int N)
{
	double temp;
	for (int i = 0; i < N-1; i++)
		for (int j = 1; j < N-i; j++)
			if (A[j] < A[j-1])
			{
				temp = A[j];
				A[j] = A[j-1];
				A[j-1] = temp;
			}
}