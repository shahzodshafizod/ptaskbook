#include <iostream>
using namespace std;

int main()
{
	//Task("Series33");
	int K, N;
	cin >> K >> N;
	int number;
	for (int i = 0; i < K; i++)
	{
		int index = 0;
		for (int j = 1; j <= N; j++)
		{
			cin >> number;
			if (number == 2)
				index = j;
		}
		cout << index;
	}
	
	return 0;
}
