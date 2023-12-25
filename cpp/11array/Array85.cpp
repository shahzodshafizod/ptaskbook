#include <iostream>
using namespace std;

int main()
{
	//Task("Array85");
	
	int N;
	cin >> N;

	double* mass = new double [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	int K;
	cin >> K;
	double* temp = new double [K];
	for (int i = N-K, j = 0; i < N; i++, j++)
		temp[j] = mass[i];
	
	for (int i = N-1; i >= K; i--)
		mass[i] = mass[i-K];
	for (int i = 0; i < K; i++)
		mass[i] = temp[i];

	for (int i = 0; i < N; i++)
		cout << mass[i];

	delete [] mass;
	delete [] temp;
	mass = temp = 0;
	
	return 0;
}
