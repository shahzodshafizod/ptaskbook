#include <iostream>
using namespace std;

int main()
{
	//Task("Array71");
	
	int N;
	cin >> N;

	double* mass = new double [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	double temp;
	for (int i = 0; i < N/2; i++)
	{
		temp = mass[i];
		mass[i] = mass[N-1-i];
		mass[N-1-i] = temp;
	}

	for (int i = 0; i < N; i++)
		cout << mass[i];

	delete [] mass;
	mass = 0;
	
	return 0;
}
