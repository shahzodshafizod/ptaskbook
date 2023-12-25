#include <iostream>
using namespace std;

int main()
{
	//Task("Array80");
	
	int N;
	cin >> N;

	double* mass = new double [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	for (int i = 0; i < N-1; i++)
		mass[i] = mass[i+1];
	mass[N-1] = 0;

	for (int i = 0; i < N; i++)
		cout << mass[i];

	delete [] mass;
	mass = 0;
	
	return 0;
}
