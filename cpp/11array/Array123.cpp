#include <iostream>
using namespace std;

int main()
{
	//Task("Array123");
	
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
		int* temp = new int [N];
		
		int start0 = 0, end0 = 0;
		element = mass[start0];
		while ( (end0 < N) && (mass[end0] == element) )
			end0++;
		end0--;

		int startK = index, endK = index;
		element = mass[startK];
		while ( (endK < N) && (mass[endK] == element) )
			endK++;
		endK--;

		int tempIndex = 0;
		
		for (int i = startK; i <= endK; i++)
			temp[tempIndex++] = mass[i];

		for (int i = end0+1; i < startK; i++)
			temp[tempIndex++] = mass[i];

		for (int i = start0; i <= end0; i++)
			temp[tempIndex++] = mass[i];

		for (int i = endK+1; i < N; i++)
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
