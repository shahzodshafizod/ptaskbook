#include <iostream>
using namespace std;

int main()
{
	//Task("Array93");
	
	int N;
	cin >> N;

	int* mass = new int [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	int newSize = N/2 + N%2;
	int* temp = new int [newSize];
	for (int i = 0, j = 0; i < N; i += 2, j++)
		temp[j] = mass[i];

	delete [] mass;
	mass = temp;
	temp = 0;
	N = newSize;
	for (int i = 0; i < N; i++)
		cout << mass[i];
	
	delete [] mass;
	mass = 0;
	
	return 0;
}
