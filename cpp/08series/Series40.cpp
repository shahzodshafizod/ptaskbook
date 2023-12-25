#include <iostream>
using namespace std;

int main()
{
	//Task("Series40");
	int K;
	cin >> K;
	int prev, curr, next;
	for (int i = 1; i <= K; i++)
	{
		int errorIndex = 0;
		int counts = 2;
		cin >> prev >> curr;
		do
		{
			cin >> next;
			if (next == 0)
				break;
			
			counts++;
			if ( (errorIndex == 0) && ( (prev > curr) && (curr > next) || (prev < curr) && (curr < next) ) )
				errorIndex = counts - 1;

			prev = curr;
			curr = next;
		}
		while (true);
		cout << ( errorIndex == 0 ? counts : errorIndex );
	}
	
	return 0;
}
