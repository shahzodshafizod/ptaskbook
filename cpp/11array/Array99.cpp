#include <iostream>
using namespace std;

int main()
{
	//Task("Array99");
	
	int N;
	cin >> N;

	int* mass = new int [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	int* temp;
	int number, count = 0;
	for (int i = 0; i < N;)
	{
		count = 0;
		for (int j = 0; j < N; j++)
			if (mass[i] == mass[j])
				count++;

		if (count > 2)
		{
			number = mass[i];
			temp = new int [N-count];
			for (int j = 0, k = 0; j < N; j++)
			{
				if (mass[j] == number)
					continue;

				temp[k++] = mass[j];
			}
			N -= count;
			delete [] mass;
			mass = temp;
			temp = 0;
		}
		else
			i++;
	}

	cout << N;
	for (int i = 0; i < N; i++)
		cout << mass[i];

	delete [] mass;
	mass = 0;
	
	return 0;
}
