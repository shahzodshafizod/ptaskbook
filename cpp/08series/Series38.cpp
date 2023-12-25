#include <iostream>
using namespace std;

int main()
{
	//Task("Series38");
	int K;
	cin >> K;
	int prev, curr;
	for (int i = 1; i <= K; i++)
	{
		bool isProgress = true, isRegress = true;
		cin >> prev;
		do
		{
			cin >> curr;
			if (curr == 0)
				break;
			
			if (prev < curr)
				isRegress = false;

			if (prev > curr)
				isProgress = false;

			prev = curr;
		}
		while (true);

		cout << ( isProgress ? 1 : isRegress ? -1 : 0 );
	}
	
	return 0;
}
