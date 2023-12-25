#include <iostream>
using namespace std;

int main()
{
	//Task("Array65");
	
	int N;
	cin >> N;

	double* A = new double [N];
	for (int i = 0; i < N; i++)
		cin >> A[i];

	int K;
	cin >> K;
	double number = A[K-1];
	for (int i = 0; i < N; i++)
		A[i] += number;

	for (int i = 0; i < N; i++)
		cout << A[i];

	delete [] A;
	A = 0;
	
	return 0;
}
