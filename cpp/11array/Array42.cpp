#include <iostream>
using namespace std;

int main()
{
	//Task("Array42");
	double R;
	cin >> R;
	int N;
	cin >> N;
	double* mass = new double [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	double closeSum = abs(mass[0] + mass[1] - R);
	int index = 0;
	for (int i = 1; i < N-1; i++)
	{
		if (abs(mass[i] + mass[i+1] - R) < closeSum)
		{
			closeSum = abs(mass[i] + mass[i+1] - R);
			index = i;
		}
	}
	cout << mass[index] << mass[index+1];

	delete [] mass;
	mass = 0;
	
	return 0;
}
