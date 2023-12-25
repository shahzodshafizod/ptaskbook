#include <iostream>
using namespace std;

int main()
{
	//Task("Minmax19");
	int N;
	cin >> N;

	int number, minimal, mins = 1;
	cin >> number;
	minimal = number;

	for (int i = 2; i <= N; i++)
	{
		cin >> number;

		if (number == minimal)
			mins++;

		if (number < minimal)
		{
			minimal = number;
			mins = 1;
		}
	}

	cout << mins;
	
	return 0;
}
