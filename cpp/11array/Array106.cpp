#include <iostream>
using namespace std;

int main()
{
	//Task("Array106");
	
	int N;
	cin >> N;

	double* mass = new double [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	double* temp = new double [N + N/2];
	memset(temp, 0, sizeof(double)*(N+N/2));
	
	for (int i = 0, j = 0; i < N; i += 2, j += 3)
		temp[j] = mass[i];
	
	for (int i = 1, j = 1; i < N; i += 2, j += 3)
		temp[j] = temp[j+1] = mass[i];

	N += N/2;
	delete [] mass;
	mass = temp;
	temp = 0;

	for (int i = 0; i < N; i++)
		cout << mass[i];

	delete [] mass;
	mass = 0;
	
	return 0;
}
