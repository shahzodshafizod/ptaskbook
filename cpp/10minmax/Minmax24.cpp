#include <iostream>
using namespace std;

int main()
{
	//Task("Minmax24");
	int N;
	cin >> N;

	double curr, prev, sum, maxSum;
	cin >> prev >> curr;

	maxSum = prev + curr;

	for (int i = 3; i <= N; i++)
	{
		prev = curr;
		cin >> curr;
		sum = prev + curr;
		if (sum > maxSum)
			maxSum = sum;
	}

	cout << maxSum;
	
	return 0;
}
