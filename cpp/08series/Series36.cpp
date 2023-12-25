#include <iostream>
using namespace std;

int main()
{
	//Task("Series36");
	int K;
	cin >> K;
	int prev, number, progresses = 0;
	for (int i = 1; i <= K; i++)
	{
		bool isProgress = true;
		cin >> prev;
		do
		{
			cin >> number;
			if (number == 0)
				break;
			if (number < prev)
				isProgress = false;
			prev = number;
		}
		while (true);
		progresses += (int)isProgress;
	}
	cout << progresses;
	
	return 0;
}
