#include <iostream>
using namespace std;

int main()
{
	//Task("Minmax6");
	int N;
	cin >> N;

	int number, firstMin, lastMax;
	int firstMinIndex, lastMaxIndex;
	cin >> number;
	firstMin = lastMax = number;
	firstMinIndex = lastMaxIndex = 1;
	
	for (int i = 2; i <= N; i++)
	{
		cin >> number;

		if (number < firstMin)
		{
			firstMin = number;
			firstMinIndex = i;
		}

		if (number >= lastMax)
		{
			lastMax = number;
			lastMaxIndex = i;
		}
	}

	cout << firstMinIndex << lastMaxIndex;
	
	return 0;
}
