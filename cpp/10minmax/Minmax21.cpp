#include <iostream>
using namespace std;

int main()
{
	//Task("Minmax21");
	int N;
	cin >> N;

	double number, minimal, maximal, sum, mins = 1, maxs = 1;
	cin >> number;
	minimal = maximal = number;
	sum = number;
	for (int i = 2; i <= N; i++)
	{
		cin >> number;
		sum += number;

		if (number > maximal)
			maximal = number;

		if (number < minimal)
			minimal = number;
	}

	sum -= minimal;
	sum -= maximal;

	sum /= (N - 2);
	
	cout << sum;
	
	return 0;
}
