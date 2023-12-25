#include <iostream>
using namespace std;

int main()
{
	//Task("Array44");
	int N;
	cin >> N;
	int* mass = new int [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	int i = 0, j = 0;
	bool found = false;
	for (; i < N; i++)
	{
		j = 0;
		for (; j < N; j++)
		{
			if (i == j)
				continue;
			if (mass[i] == mass[j])
			{
				found = true;
				break;
			}
		}
		if (found)
			break;
	}	
	
	if (i < j)
		cout << i+1 << j+1;
	else
		cout << j+1 << i+1;

	delete [] mass;
	mass = 0;
	
	return 0;
}
