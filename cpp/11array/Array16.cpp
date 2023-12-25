#include <iostream>
using namespace std;

int main()
{
	//Task("Array16");
	int N;
	cin >> N;
	
	double* mass = new double [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];
	
	int i = 0;
	for (; i < N/2; i++)
		cout << mass[i] << mass[N-1-i];
	
	if (N%2 != 0)
		cout << mass[i];

	delete [] mass;
	mass = 0;
	
	return 0;
}
