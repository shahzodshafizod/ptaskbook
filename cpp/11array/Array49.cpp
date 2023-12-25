#include <iostream>
using namespace std;

int main()
{
	//Task("Array49");
	
	int N;
	cin >> N;
	
	int* mass = new int [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	bool alone;
	int index = -1;
	for (int i = 0; i < N; i++)
	{
		alone = true;
		for (int j = 0; j < i; j++)
			if (mass[i] == mass[j])
			{
				alone = false;
				break;
			}
		
		if ( (alone == false) || (mass[i] <= 0) || (mass[i] > N))
		{
			index = i;
			break;
		}
	}
	
	cout << index+1;

	delete [] mass;
	mass = 0;
	
	return 0;
}
