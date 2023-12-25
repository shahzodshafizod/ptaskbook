#include <iostream>
using namespace std;

int main()
{
	//Task("Array121");
	
	int K;
	cin >> K;

	int N;
	cin >> N;

	int element, series = 0, index = -1;
	int* mass = new int [N];
	cin >> mass[0];
	element = mass[0];
	series++;
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
			temp = new int [N+1];
			for (int i = 0, j = 0; i < N; i++)
			{
				temp[j++] = mass[i];

				if (i == index)
					temp[j++] = mass[i];
			}
			delete [] mass;
			mass = temp;
			temp = 0;
			N++;
			index += 2;
		}
	}

	for (int i = 0; i < N; i++)
		cout << mass[i];

	delete [] mass;
	mass = 0;
	
	return 0;
}
