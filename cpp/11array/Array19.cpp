#include <iostream>
using namespace std;

int main()
{
	//Task("Array19");
	int mass[10];
	for (int i = 0; i < 10; i++)
		cin >> mass[i];
	
	int index = 0;
	for (int i = 1; i < 9; i++)
	{
		if ( (mass[i] > mass[0]) && (mass[i] < mass[9]) )
		{
			index = i+1;
		}
	}

	cout << index;
	
	return 0;
}
