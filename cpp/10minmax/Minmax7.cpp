#include <iostream>
using namespace std;

int main()
{
	//Task("Minmax7");
	int N;
	cin >> N;

	int number, firstMax, lastMin;
	int firstMaxIndex, lastMinIndex;
	cin >> number;
	firstMax = lastMin = number;
	firstMaxIndex =lastMinIndex = 1;

	for (int i = 2; i <= N; i++)
	{
		cin >> number;

		if (number > firstMax)
		{
			firstMax = number;
			firstMaxIndex = i;
		}

		if (number <= lastMin)
		{
			lastMin = number;
			lastMinIndex = i;
		}
	}

	cout << firstMaxIndex << lastMinIndex;
	
	return 0;
}
