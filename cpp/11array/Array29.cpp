#include <iostream>
using namespace std;

int main()
{
	//Task("Array29");
	int N;
	cin >> N;
	double* mass = new double [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];
	double maximal = mass[0];
	for (int i = 2; i < N; i += 2)
		if (mass[i] > maximal)
			maximal = mass[i];
	cout << maximal;
	delete [] mass;
	mass = 0;
	
	return 0;
}
