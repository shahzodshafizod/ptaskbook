#include <iostream>
using namespace std;

int main()
{
	//Task("Array96");
	
	int N;
	cin >> N;

	int* mass = new int [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	bool hast = false;
	for (int i = 1; i < N;)
	{
		hast = false;
		for (int j = 0; j < i; j++)
			if (mass[i] == mass[j])
				hast = true;

		if (hast)
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
