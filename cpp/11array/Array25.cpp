#include <iostream>
using namespace std;

int main()
{
	//Task("Array25");
	int N;
	cin >> N;
	int* mass = new int [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];
	double maxraj = (double)mass[1] / mass[0];
	for (int i = 2; i < N; i++)
	{
		if ((double)mass[i] / mass[i-1] != maxraj)
		{
			maxraj = 0;
			break;
		}
	}
	cout << (int)maxraj;
	delete [] mass;
	mass = 0;
	
	return 0;
}
