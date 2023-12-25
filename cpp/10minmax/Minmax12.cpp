#include <iostream>
using namespace std;

int main()
{
	//Task("Minmax12");
	int N;
	cin >> N;

	double number, minimal = 0;
	for (int i = 1; i <= N; i++)
	{
		cin >> number;

		if (number > 0)
		{
			if (minimal == 0)
				minimal = number;
			else if (number < minimal)
				minimal = number;
		}
	}
	
	cout << minimal;
	
	return 0;
}
