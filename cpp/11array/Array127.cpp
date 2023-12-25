#include <iostream>
using namespace std;

void ZeroSerie(int** mass, int& N, int index, int elements);

int main()
{
	//Task("Array127");
	
	int L, N;
	cin >> L >> N;

	int* mass = new int [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	int index = 0;
	int element = mass[index];
	int elements = 1;
	for (int i = 1; i < N; i++)
	{
		if (mass[i] != element)
		{
			if (elements > L)
			{
				ZeroSerie(&mass, N, index, elements);

				i = ++index;
			}
			else
				index = i;
			
			element = mass[index];
			elements = 1;
		}
		else
			elements++;
	}

	if (elements > L)
		ZeroSerie(&mass, N, index, elements);

	for (int i = 0; i < N; i++)
		cout << mass[i];

	delete [] mass;
	mass = 0;
	
	return 0;
}

void ZeroSerie(int** mass, int& N, int index, int elements)
{
	int* temp = new int [N-elements+1];
	int tempIndex = 0;
	
	for (int i = 0; i < index; i++)
		temp[tempIndex++] = (*mass)[i];

	temp[tempIndex++] = 0;

	for (int i = index+elements; i < N; i++)
		temp[tempIndex++] = (*mass)[i];

	delete [] (*mass);
	(*mass) = temp;
	temp = 0;

	N -= elements;
	N++;
}