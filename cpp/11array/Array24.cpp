#include <iostream>
using namespace std;

int main()
{
	//Task("Array24");
	int N;
	cin >> N;
	int* mass = new int [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];
	int farq = mass[1] - mass[0];
	for (int i = 2; i < N; i++)
	{
		if (mass[i] - mass[i-1] != farq)
		{
			farq = 0;
			break;
		}
	}
	cout << farq;
	delete [] mass;
	mass = 0;
	
	return 0;
}
