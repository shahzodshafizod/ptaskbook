#include <iostream>
using namespace std;

int main()
{
	//Task("Array73");
	
	int N;
	cin >> N;

	double* A = new double [N];
	for (int i = 0; i < N; i++)
		cin >> A[i];

	int K, L;
	cin >> K >> L;
	double temp;
	int from = K, to = L-2, howMany = to - from + 1;
	for (int i = from, j = 0; j < howMany/2; j++, i++)
	{
		temp = A[i];
		A[i] = A[to -j];
		A[to - j] = temp;
	}

	for (int i = 0; i < N; i++)
		cout << A[i];

	delete [] A;
	A = 0;
	
	return 0;
}
