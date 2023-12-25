#include <iostream>
using namespace std;

int main()
{
	//Task("Minmax17");
	int N;
	cin >> N;

	int number, maximal;
	cin >> number;
	maximal = number;
	
	int index = 1;
	for (int i = 2; i <= N; i++)
	{
		cin >> number;
		
		if (number >= maximal)
		{
			maximal = number;
			index = i;
		}
	}

	cout << (N - index);
	
	return 0;
}
