#include <iostream>
using namespace std;

int main()
{
	//Task("Series18");
	int N;
	cin >> N;
	int prev, number;
	cin >> prev;
	cout << prev;
	for (int i = 2; i <= N; i++)
	{
		cin >> number;
		if (number != prev)
			cout << number;
		prev = number;
	}
	
	return 0;
}
