#include <iostream>
using namespace std;

int main()
{
	//Task("Series30");
	int K, N;
	cin >> K >> N;
	int number, sum;
	for (int i = 1; i <= K; i++)
	{
		sum = 0;
		for (int j = 1; j <= N; j++)
		{
			cin >> number;
			sum += number;
		}
		cout << sum;
	}
	
	return 0;
}
