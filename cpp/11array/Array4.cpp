#include <iostream>
#include <cmath>
using namespace std;

int main()
{
	//Task("Array4");
	int N;
	cin >> N;
	double A, D;
	cin >> A >> D;
	
	double* mass = new double [N];
	for (int i = 0; i < N; i++)
		mass[i] = A	* pow(D, i);
	
	for (int i = 0; i < N; i++)
		cout << mass[i];
	
	delete [] mass;
	mass = 0;
	
	return 0;
}
