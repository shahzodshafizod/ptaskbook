#include <iostream>
using namespace std;

int main()
{
	//Task("Series34");
	int K, N;
	cin >> K >> N;
	int number, sum;
	for (int i = 0; i < K; i++)
	{
		sum = 0;
		bool has2 = false;
		for (int j = 1; j <= N; j++)
		{
			cin >> number;
			if (number == 2)
				has2 = true;
			sum += number;
		}
		cout << (has2 ? sum : 0);
	}
	
	return 0;
}
