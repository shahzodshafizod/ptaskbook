#include <iostream>
using namespace std;

int main()
{
	//Task("Series35");
	int K, number, counts, totalCounts = 0;
	cin >> K;
	for (int i = 0; i < K; i++)
	{
		counts = 0;
		do
		{
			cin >> number;
			if (number == 0)
				break;
			counts++;
		}
		while (true);
		cout << counts;
		totalCounts += counts;
	}
	cout << totalCounts;
	
	return 0;
}
