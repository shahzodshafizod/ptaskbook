#include <iostream>
using namespace std;

int main()
{
	//Task("Array95");
	
	int N;
	cin >> N;

	int* mass = new int [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	for (int i = 1; i < N;)
	{
		if (mass[i] == mass[i-1])
		{
			N--;
			for (int j = i; j < N; j++)
				mass[j] = mass[j+1];
		}
		else
			i++;
	}

	for (int i = 0; i < N; i++)
		cout << mass[i];
	
	delete [] mass;
	mass = 0;
	
	return 0;
}
