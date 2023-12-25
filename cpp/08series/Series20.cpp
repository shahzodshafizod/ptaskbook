#include <iostream>
using namespace std;

int main()
{
	//Task("Series20");
	int N;
	cin >> N;
	int prev, number, K = 0;
	cin >> prev;
	for (int i = 2; i <= N; i++)
	{
		cin >> number;
		if (prev < number)
		{
			cout << prev;
			K++;
		}
		prev = number;
	}
	cout << K;
	
	return 0;
}
