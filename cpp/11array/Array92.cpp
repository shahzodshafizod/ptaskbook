#include <iostream>
using namespace std;

int main()
{
	//Task("Array92");
	
	int N;
	cin >> N;

	int* mass = new int [N];
	int newSize = N;
	for (int i = 0; i < N; i++)
	{
		cin >> mass[i];
		if (mass[i] % 2 != 0)
			newSize--;
	}

	int* temp = new int [newSize];
	for (int i = 0, j = 0; i < N; i++)
	{
		if (mass[i] % 2 != 0)
			continue;
		temp[j++] = mass[i];
	}

	N = newSize;
	delete [] mass;
	mass = temp;
	temp = 0;
	cout << N;
	for (int i = 0; i < N; i++)
		cout << mass[i];

	delete [] mass;
	mass = 0;
	
	return 0;
}
