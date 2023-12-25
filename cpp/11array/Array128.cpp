#include <iostream>
using namespace std;

int main()
{
	//Task("Array128");
	
	int N;
	cin >> N;

	int* mass = new int [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	int index = 0;
	int element = mass[index];
	int elements = 1, maxElements = 1;
	int maxElementsIndex = 0;
	for (int i = 1; i < N; i++)
	{
		if (mass[i] != element)
		{
			if (elements > maxElements)
			{
				maxElements = elements;
				maxElementsIndex = index;
			}
			
			index = i;
			element = mass[index];
			elements = 1;
		}
		else
			elements++;
	}
	if (elements > maxElements)
		maxElementsIndex = index;

	int* temp = new int [N+1];
	int tempIndex = 0;

	for (int i = 0; i < maxElementsIndex; i++)
		temp[tempIndex++] = mass[i];

	temp[tempIndex++] = mass[maxElementsIndex];

	for (int i = maxElementsIndex; i < N; i++)
		temp[tempIndex++] = mass[i];

	delete [] mass;
	mass = temp;
	temp = 0;

	N++;

	for (int i = 0; i < N; i++)
		cout << mass[i];

	delete [] mass;
	mass = 0;
	
	return 0;
}
