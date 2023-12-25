#include <iostream>
using namespace std;

int main()
{
	//Task("Array83");
	
	int N;
	cin >> N;

	double* mass = new double [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	double temp = mass[N-1];
	for (int i = N-1; i > 0; i--)
		mass[i] = mass[i-1];
	mass[0] = temp;

	for (int i = 0; i < N; i++)
		cout << mass[i];

	delete [] mass;
	mass = 0;
	
	return 0;
}
