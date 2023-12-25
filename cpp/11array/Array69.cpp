#include <iostream>
using namespace std;

int main()
{
	//Task("Array69");
	
	int N;
	cin >> N;

	double* mass = new double [N];
	double temp;
	for (int i = 0; i < N; i++)
	{
		cin >> mass[i];
		if (i % 2 != 0)
		{
			temp = mass[i];
			mass[i] = mass[i-1];
			mass[i-1] = temp;
		}
	}

	for (int i = 0; i < N; i++)
		cout << mass[i];

	delete [] mass;
	mass = 0;
	
	return 0;
}
