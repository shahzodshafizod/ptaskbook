#include <iostream>
using namespace std;

int main()
{
	//Task("Array122");
	
	int K;
	cin >> K;

	int N;
	cin >> N;

	int* mass = new int [N];
	cin >> mass[0];
	int element = mass[0];
	int series = 1, index = -1;
	if (series == K)
		index = 0;
	for (int i = 1; i < N; i++)
	{
		cin >> mass[i];

		if (mass[i] != element)
		{
			element = mass[i];
			series++;

			if (series == K)
				index = i;
		}
	}

	if (index >= 0)
	{
		element = mass[index];
		int* temp = 0;
		while ( (index < N) && (mass[index] == element) )
		{
			temp = new int [N-1];
			for (int i = 0, j = 0; i < N; i++)
			{
				if (i == index)
					continue;
				
				temp[j++] = mass[i];
			}

			delete [] mass;
			mass = temp;
			temp = 0;
			N--;
		}
	}

	for (int i = 0; i < N; i++)
		cout << mass[i];

	delete [] mass;
	mass = 0;
	
	return 0;
}
