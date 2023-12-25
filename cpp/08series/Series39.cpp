#include <iostream>
using namespace std;

int main()
{
	//Task("Series39");
	int K;
	cin >> K;
	int prev, curr, next, counts = 0;
	for (int i = 1; i <= K; i++)
	{
		bool ok = true;
		cin >> prev >> curr;
		do
		{
			cin >> next;
			if (next == 0)
				break;
			if ( (prev > curr) && (curr > next) || (prev < curr) && (curr < next) )
				ok = false;

			prev = curr;
			curr = next;
		}
		while (true);
		counts += (int)ok;
	}
	cout << counts;
	
	return 0;
}
