#include <iostream>
using namespace std;

void IncSerie(int** mass, int& N, int index);

int main()
{
	//Task("Array130");
	
	int N;
	cin >> N;

	int* mass = new int [N];
	cin >> mass[0];
	int element = mass[0];
	int elements = 1;
	int maxElements = 1;
	for (int i = 1; i < N; i++)
	{
		cin >> mass[i];

		if (mass[i] != element)
		{
			if (elements > maxElements)
				maxElements = elements;

			element = mass[i];
			elements = 1;
		}
		else
			elements++;
	}
	if (elements > maxElements)
		maxElements = elements;

	int index = 0;
	element = mass[index];
	elements = 1;
	for (int i = 1; i < N; i++)
	{
		if (mass[i] != element)
		{
			if (elements == maxElements)
			{
				IncSerie(&mass, N, index);
				i++;
			}
		
			index = i;
			element = mass[index];
			elements = 1;
		}
		else
			elements++;
	}
	if (elements == maxElements)
		IncSerie(&mass, N, index);

	for (int i = 0; i < N; i++)
		cout << mass[i];

	delete [] mass;
	mass = 0;
	
	return 0;
}

void IncSerie(int** mass, int& N, int index)
{
	int* temp = new int [N+1];
	int tempIndex = 0;

	for (int i = 0; i < index; i++)
		temp[tempIndex++] = (*mass)[i];

	temp[tempIndex++] = (*mass)[index];

	for (int i = index; i < N; i++)
		temp[tempIndex++] = (*mass)[i];

	delete [] (*mass);
	(*mass) = temp;
	temp = 0;

	N++;
}