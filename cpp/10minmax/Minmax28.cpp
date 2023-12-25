#include <iostream>
using namespace std;

int main()
{
	//Task("Minmax28");
	int N;
	cin >> N;

	int number, counts = 0, maxCountsIndex = 0, maxCounts = 0;
	for (int i = 1; i <= N; i++)
	{
		cin >> number;

		if (number == 1)
		{
			counts++;
		}
		else
		{
			if (counts == 0)
				continue;
			if (maxCounts == 0)
			{
				maxCounts = counts;
				maxCountsIndex = i - counts;
			}
			else if (counts >= maxCounts)
			{
				maxCounts = counts;
				maxCountsIndex = i - counts;
			}
			counts = 0;
		}
	}
	if (counts != 0)
	{
		if (maxCounts == 0)
		{
			maxCounts = counts;
			maxCountsIndex = N - counts + 1;
		}
		else if (counts >= maxCounts)
		{
			maxCounts = counts;
			maxCountsIndex = N - counts + 1;
		}
	}

	cout << maxCountsIndex << maxCounts;
	
	return 0;
}
