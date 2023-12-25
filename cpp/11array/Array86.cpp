#include <iostream>
using namespace std;

int main()
{
	//Task("Array86");
	
	int N;
	cin >> N;

	double* mass = new double [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	int K;
	cin >> K;
	double* temp = new double [K];
	for (int i = 0; i < K; i++)
		temp[i] = mass[i];

	for (int i = 0; i < N-K; i++)
		mass[i] = mass[i+K];
	for (int i = 0, j = N-K; j < N; j++, i++)
		mass[j] = temp[i];

	for (int i = 0; i < N; i++)
		cout << mass[i];

	delete [] mass;
	mass = 0;
	
	return 0;
}
