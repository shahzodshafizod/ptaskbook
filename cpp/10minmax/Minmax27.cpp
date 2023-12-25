#include <iostream>
using namespace std;

int main()
{
	//Task("Minmax27");
	int N;
	cin >> N;

	int prev, curr;
	cin >> prev;

	int counts = 1, maxCounts = 0, index = 1;
	for (int i = 2; i <= N; i++)
	{
		cin >> curr;

		if (curr == prev)
			counts++;
		else
		{
			if (maxCounts == 0)
			{
				maxCounts = counts;
				index = i - maxCounts;
			}
			else if (counts > maxCounts)
			{
				maxCounts = counts;
				index = i - maxCounts;
			}
			counts = 1;
		}

		prev = curr;
	}
	if (maxCounts == 0)
	{
		maxCounts = counts;
		index = N - maxCounts + 1;
	}
	else if (counts > maxCounts)
	{
		maxCounts = counts;
		index = N - maxCounts + 1;
	}
	
	cout << index << maxCounts;
	
	return 0;
}
