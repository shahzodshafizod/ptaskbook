#include <iostream>
using namespace std;

int main()
{
	//Task("Series23");
	int N;
	cin >> N;
	double prev, curr, next;
	cin >> prev >> curr;
	int errorIndex = 0;
	for (int i = 3; i <= N; i++)
	{
		cin >> next;

		if ( (errorIndex == 0) && !((curr > prev) && (curr > next) || (curr < prev) && (curr < next)) )
			errorIndex = i-1;
		prev = curr;
		curr = next;
	}
	cout << errorIndex;
	
	return 0;
}
