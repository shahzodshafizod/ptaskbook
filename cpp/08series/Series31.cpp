#include <iostream>
using namespace std;

int main()
{
	//Task("Series31");
	int K, N;
	cin >> K >> N;
	int number, counts = 0;
	for (int i = 1; i <= K; i++)
	{
		bool has2 = false;
		for (int j = 1; j <= N; j++)
		{
			cin >> number;
			if (number == 2)
				has2 = true;
		}
		counts += (int)has2;
	}
	cout << counts;
	
	return 0;
}
