#include <iostream>
using namespace std;

int main()
{
	//Task("Array39");
	
	int N;
	cin >> N;
	
	double* mass = new double [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	int count = 0;
	for (int i = 1; i < N-1; i++)
	{
		if ( (mass[i-1] < mass[i]) && (mass[i] > mass[i+1]) ||
			 (mass[i-1] > mass[i]) && (mass[i] < mass[i+1]) )
			 count++;
	}
	if (mass[N-1] != mass[N-2])
		count++;

	cout << count;

	delete [] mass;
	mass = 0;
	
	return 0;
}
