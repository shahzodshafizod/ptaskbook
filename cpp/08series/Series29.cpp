#include <iostream>
using namespace std;

int main()
{
	//Task("Series29");
	int K, N;
	cin >> K >> N;
	int number, sum = 0;
	for (int i = 1; i <= K; i++)
		for (int j = 1; j <= N; j++)
		{
			cin >> number;
			sum += number;
		}

	cout << sum;
	
	return 0;
}
