#include <iostream>
using namespace std;

int main()
{
	//Task("Minmax18");
	int N;
	cin >> N;

	int number, maximal;
	cin >> number;
	maximal = number;

	int firstMaxIndex = 1, lastMaxIndex = 1;

	for (int i = 2; i <= N; i++)
	{
		cin >> number;

		if (number == maximal)
			lastMaxIndex = i;

		if (number > maximal)
		{
			maximal = number;
			firstMaxIndex = lastMaxIndex = i;
		}
	}

	int counts = lastMaxIndex - firstMaxIndex;
	if (firstMaxIndex < lastMaxIndex)
		counts--;
	
	cout << counts;
	
	return 0;
}
