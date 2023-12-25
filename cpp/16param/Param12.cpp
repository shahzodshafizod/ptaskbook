#include <iostream>
using namespace std;

void SortIndex(const double A[], int N, int I[]);

int main()
{
	//Task("Param12");
	int N, *index = NULL;
	double* mass = NULL;

	for (int i = 1; i <= 3; i++)
	{
		cin >> N;
		mass = new double [N];
		for (int j = 0; j < N; j++)
			cin >> mass[j];
		index = new int [N];
		SortIndex(mass, N, index);
		for (int j = 0; j < N; j++)
			cout << (index[j] + 1);

		delete [] mass;
		mass = NULL;
		delete [] index;
		index = NULL;
	}
	
	return 0;
}

void SortIndex(const double A[], int N, int I[])
{
	for (int i = 0; i < N; i++)
		I[i] = i;

	int temp;
	for (int i = 0; i < N-1; i++)
		for (int j = 1; j < N; j++)
			if ( A[ I[j] ] < A[ I[j-1] ] )
			{
				temp = I[j];
				I[j] = I[j-1];
				I[j-1] = temp;
			}
}