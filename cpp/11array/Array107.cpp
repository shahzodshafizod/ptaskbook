#include <iostream>
using namespace std;

int main()
{
	//Task("Array107");
	
	int N;
	cin >> N;

	double* mass = new double [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	double* temp = new double [N + (N/2+N%2)*2];
	
	for (int i = 0, j = 0; i < N; i += 2, j += 4)
		temp[j] = temp[j+1] = temp[j+2] = mass[i];

	for (int i = 1, j = 3; i < N; i += 2, j += 4)
		temp[j] = mass[i];

	delete [] mass;
	mass = temp;
	temp = 0;
	N += (N/2 + N%2) * 2;

	for (int i = 0; i < N; i++)
		cout << mass[i];

	delete [] mass;
	mass = 0;
	
	return 0;
}
