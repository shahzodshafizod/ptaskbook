#include <iostream>
using namespace std;

int main()
{
	//Task("Array97");
	
	int N;
	cin >> N;

	int* mass = new int [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	bool hast = false;
	for (int i = N-2; i >= 0;)
	{
		hast = false;
		for (int j = i+1; j < N; j++)
			if (mass[j] == mass[i])
				hast = true;
		if (hast)
		{
			N--;
			for (int j = i; j < N; j++)
				mass[j] = mass[j+1];
		}
		else
			i--;
	}

	for (int i = 0; i < N; i++)
		cout << mass[i];

	delete [] mass;
	mass = 0;
	
	return 0;
}
