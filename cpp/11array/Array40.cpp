#include <iostream>
#include <cmath>
using namespace std;

int main()
{
	//Task("Array40");
	
	double R;
	cin >> R;
	
	int N;
	cin >> N;
	
	double* mass = new double [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	double minimal = abs(R - mass[0]);
	double element = mass[0];
	for (int i = 1; i < N; i++)
		if (abs(mass[i]-R) < minimal)
		{
			minimal = abs(mass[i] - R);
			element = mass[i];
		}

	cout << element;
	
	delete [] mass;
	mass = 0;
	
	return 0;
}
