#include <iostream>
using namespace std;

int main()
{
	//Task("Array81");
	
	int N;
	cin >> N;

	double* mass = new double [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	int K;
	cin >> K;
	for (int i = N-1; i >= K; i--)
		mass[i] = mass[i-K];
	for (int i = 0; i < K; i++)
		mass[i] = 0;

	for (int i = 0; i < N; i++)
		cout << mass[i];

	delete [] mass;
	mass = 0;
	
	return 0;
}
