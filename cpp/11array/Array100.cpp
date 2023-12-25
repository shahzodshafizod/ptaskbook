#include <iostream>
using namespace std;

int main()
{
	//Task("Array100");
	
	int N;
	cin >> N;

	int* mass = new int [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	int *temp = 0;
	int number, count = 0;
	for (int i = 0; i < N;)
	{
		count = 0;
		for (int j = 0; j < N; j++)
			if (mass[i] == mass[j])
				count++;
		
		if (count == 2)
		{
			number = mass[i];
			temp = new int [N-count];
			for (int k = 0, j = 0; k < N; k++)
			{
				if (mass[k] == number)
					continue;
				
				temp[j++] = mass[k];
			}
			delete [] mass;
			mass = temp;
			temp = 0;
			N -= count;
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
