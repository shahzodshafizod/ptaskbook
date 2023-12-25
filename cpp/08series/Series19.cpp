#include <iostream>
using namespace std;

int main()
{
	//Task("Series19");
	int N;
	cin >> N;
	int number, prev;
	cin >> prev;
	int K = 0;
	for (int i = 2; i <= N; i++)
	{
		cin >> number;
		if (number < prev)
		{
			cout << number;
			K++;
		}
		prev = number;
	}
	cout << K;
	
	return 0;
}
