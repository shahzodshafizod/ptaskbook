#include <iostream>
using namespace std;

int main()
{
	//Task("Minmax1");
	int N;
	cin >> N;
	
	double number, minimal, maximal;
	cin >> number;
	minimal = maximal = number;
	for (int i = 2; i <= N; i++)
	{
		cin >> number;

		if (number > maximal)
			maximal = number;
		if (number < minimal)
			minimal = number;
	}

	cout << minimal << maximal;
	
	return 0;
}
