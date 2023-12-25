#include <iostream>
using namespace std;

int main()
{
	//Task("Series24");
	int N;
	cin >> N;
	int number, sum = 0, lastSum = 0;
	for (int i = 1; i <= N; i++)
	{
		cin >> number;
		sum += number;
		if (number == 0)
		{
			lastSum = sum;
			sum = 0;
		}
	}
	cout << lastSum;
	
	return 0;
}
