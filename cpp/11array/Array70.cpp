#include <iostream>
using namespace std;

int main()
{
	//Task("Array70");
	
	int N;
	cin >> N;

	double* mass = new double [N];
	double temp;
	for (int i = 0; i < N; i++)
	{
		cin >> mass[i];
		if (i >= N/2)
		{
			temp = mass[i];
			mass[i] = mass[i-N/2];
			mass[i-N/2] = temp;
		}
	}

	for (int i = 0; i < N; i++)
		cout << mass[i];

	delete [] mass;
	mass = 0;
	
	return 0;
}
