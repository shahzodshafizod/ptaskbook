#include <iostream>
using namespace std;

int main()
{
	//Task("Series10");
	int N;
	cin >> N;
	int number;
	bool hasPositive = false;
	for (int i = 1; i <= N; i++)
	{
		cin >> number;
		if (number > 0)
			hasPositive = true;
	}
	cout << hasPositive;
	
	return 0;
}
