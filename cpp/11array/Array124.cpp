#include <iostream>
using namespace std;

int main()
{
	//Task("Array124");
	
	int K, N;
	cin >> K >> N;

	int* mass = new int [N];
	cin >> mass[0];
	int element = mass[0];
	int series = 1;
	int index = -1;
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
		int startK = index, endK = index;
		element = mass[endK];
		while ( (endK < N) && (mass[endK] == element) )
			endK++;
		endK--;

		int startLast = N-1, endLast = N-1;
		element = mass[startLast];
		while ( (startLast >= 0) && (mass[startLast] == element) )
			startLast--;
		startLast++;

		int* temp = new int [N];
		int tempIndex = 0;

		for (int i = 0; i < startK; i++)
			temp[tempIndex++] = mass[i];

		for (int i = startLast; i <= endLast; i++)
			temp[tempIndex++] = mass[i];

		for (int i = endK+1; i < startLast; i++)
			temp[tempIndex++] = mass[i];

		for (int i = startK; i <= endK; i++)
			temp[tempIndex++] = mass[i];

		delete [] mass;
		mass = temp;
		temp = 0;
	}

	for (int i = 0; i < N; i++)
		cout << mass[i];

	delete [] mass;
	mass = 0;
	
	return 0;
}
