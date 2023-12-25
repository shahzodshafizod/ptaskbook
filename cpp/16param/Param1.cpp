#include <iostream>
using namespace std;

int MinElem(int A[], int N);

int main()
{
	//Task("Param1");
	int N;
	int* mass = NULL;

	for (int i = 1; i <= 3; i++)
	{
		cin >> N;
		mass = new int [N];
		for (int j = 0; j < N; j++)
			cin >> mass[j];
		cout << MinElem(mass, N);
		delete [] mass;
		mass = NULL;
	}
	
	return 0;
}

int MinElem(int A[], int N)
{
	int minimal = A[0];
	for (int i = 1; i < N; i++)
		if (A[i] < minimal)
			minimal = A[i];
	
	return minimal;
}