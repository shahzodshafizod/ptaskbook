#include <iostream>
using namespace std;

int main()
{
	//Task("Array82");
	
	int N;
	cin >> N;

	double* mass = new double [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	int K;
	cin >> K;
	for (int i = 0; i < N-K; i++)
		mass[i] = mass[i+K];
	for (int i = N-K; i < N; i++)
		mass[i] = 0;

	for (int i = 0; i < N; i++)
		cout << mass[i];

	delete [] mass;
	mass = 0;
	
	return 0;
}
