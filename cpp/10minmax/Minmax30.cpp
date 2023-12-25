#include <iostream>
using namespace std;

int main()
{
	//Task("Minmax30");
	int N;
	cin >> N;

	int number, maximal, counts = 1, maxCounts = 0;
	cin >> number;
	maximal = number;

	for (int i = 2; i <= N; i++)
	{
		cin >> number;

		if (number == maximal)
			counts++;

		else if (number > maximal)
		{
			maximal = number;
			counts = 1;
			maxCounts = 0;
		}

		else
		{
			if (counts == 0)
				continue;
			
			if ( maxCounts == 0 )
				maxCounts = counts;
			else if (counts < maxCounts)
				maxCounts = counts;
			
			counts = 0;
		}
	}
	if ( maxCounts == 0 )
		maxCounts = counts;
	else if ( (counts != 0) && (counts < maxCounts) )
		maxCounts = counts;

	cout << maxCounts;
	
	return 0;
}
