#include <iostream>
using namespace std;

int main()
{
	//Task("Array91");
	
	int N;
	cin >> N;

	double* mass = new double [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	int K, L;
	cin >> K >> L;
	int newSize = N-(L-K+1);
	double* temp = new double [newSize];
	
	int index = 0;
	for (int i = 0; i < K-1; i++)
		temp[index++] = mass[i];
	for (int i = L; i < N; i++)
		temp[index++] = mass[i];

	delete [] mass;
	mass = temp;
	temp = 0;

	N = newSize;
	cout << N;
	for (int i = 0; i < N; i++)
		cout << mass[i];

	delete [] mass;
	mass = 0;
	
	return 0;
}
