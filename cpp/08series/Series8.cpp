#include <iostream>
using namespace std;

int main()
{
	//Task("Series8");
	int N;
	cin >> N;
	int number, K = 0;
	for (int i = 1; i <= N; i++)
	{
		cin >> number;
		if (number % 2 == 0)
		{
			cout << number;
			K++;
		}
	}
	cout << K;
	
	return 0;
}
