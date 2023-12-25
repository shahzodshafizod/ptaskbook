#include <iostream>
using namespace std;

int main()
{
	//Task("Array111");
	
	int N;
	cin >> N;

	int* mass = new int [N];
	int unevens = 0;
	for (int i = 0; i < N; i++)
	{
		cin >> mass[i];

		if (mass[i] % 2 != 0)
			unevens++;
	}

	int* temp = new int [N + unevens*2];
	int index = 0;
	for (int i = 0; i < N; i++)
	{
		temp[index++] = mass[i];

		if (mass[i] % 2 != 0)
			temp[index+++1] = temp[index++] = mass[i];
			//temp[1+index++] = temp[index++] = mass[i];
	}

	delete [] mass;
	mass = temp;
	temp = 0;
	N += unevens*2;

	for (int i = 0; i < N; i++)
		cout << mass[i];

	delete [] mass;
	mass = 0;
	
	return 0;
}
