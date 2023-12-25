#include <iostream>
using namespace std;

int main()
{
	//Task("Array28");
	int N;
	cin >> N;
	double* mass = new double [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];
	double minimal = mass[1];
	for (int i = 3; i < N; i+= 2)
		if (mass[i] < minimal)
			minimal = mass[i];
	cout << minimal;
	delete [] mass;
	mass = 0;
	
	return 0;
}
