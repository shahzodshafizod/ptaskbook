#include <iostream>
using namespace std;

int main()
{
	//Task("Minmax11");
	int N;
	cin >> N;

	int number, minimal, maximal;
	cin >> number;
	minimal = maximal = number;

	int minIndex, maxIndex;
	minIndex = maxIndex = 1;

	for (int i = 2; i <= N; i++)
	{
		cin >> number;

		if (number <= minimal)
		{
			minimal = number;
			minIndex = i;
		}

		if (number >= maximal)
		{
			maximal = number;
			maxIndex = i;
		}
	}

	cout << ( minIndex > maxIndex ? minIndex : maxIndex );
	
	return 0;
}
