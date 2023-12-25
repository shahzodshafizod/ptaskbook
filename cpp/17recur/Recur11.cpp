#include <iostream>
using namespace std;

int MaxElem(int* A, int N);

int main()
{
	//Task("Recur11");
	int N;
	int* mass = NULL;
	for (int i = 1; i < 4; i++)
	{
		cin >> N;
		mass = new int [N];
		for (int j = 0; j < N; j++)
			cin >> mass[j];
		cout << MaxElem(mass, N);

		delete [] mass;
		mass = NULL;
	}
	return 0;
}

int MaxElem(int* A, int N)
{
	if (N == 0)
		return A[N];

	int elem = MaxElem(A, N-1);
	return ( A[N-1] > elem ? A[N-1] : elem );
}