#include <iostream>
using namespace std;

int main()
{
	//Task("Series25");
	int N;
	cin >> N;
	int number, sum = 0, allSum = 0;
	bool started = false;
	for (int i = 1; i <= N; i++)
	{
		cin >> number;
		sum += number;
		if (number == 0)
		{
			if (started == false)
				started = true;
			else
				allSum += sum;
			sum = 0;
		}
	}
	cout << allSum;
	
	return 0;
}
