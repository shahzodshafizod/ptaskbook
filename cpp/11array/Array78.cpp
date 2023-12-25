#include <iostream>
using namespace std;

int main()
{
	//Task("Array78");
	
	int N;
	cin >> N;

	double* mass = new double [N];
	double* temp = new double [N];
	for (int i = 0; i < N; i++)
	{
		cin >> mass[i];
		temp[i] = mass[i];
	}

	mass[0] = (temp[0] + temp[1]) / 2;
	mass[N-1] = (temp[N-1] + temp[N-2]) / 2;
	for (int i = 1; i < N-1; i++)
		mass[i] = (temp[i-1] + temp[i] + temp[i+1]) / 3;

	for (int i = 0; i < N; i++)
		cout << mass[i];

	delete [] mass;
	delete [] temp;
	mass = temp = 0;
	
	return 0;
}
