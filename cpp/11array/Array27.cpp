#include <iostream>
using namespace std;

int main()
{
	//Task("Array27");
	int N;
	cin >> N;
	int* mass = new int [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];
	
	int index = 0;
	for (int i = 1; i < N; i++)
	{
		if (mass[i] * mass[i-1] > 0)
		{
			index = i+1;
			break;
		}
	}
	cout << index;
	delete [] mass;
	mass = 0;
	
	return 0;
}
