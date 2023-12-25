#include <iostream>
using namespace std;

int main()
{
	//Task("Series22");
	int N;
	cin >> N;
	double prev, number;
	int errorIndex = 0;
	cin >> prev;
	for (int i = 2; i <= N; i++)
	{
		cin >> number;
		
		if ( (errorIndex == 0) && (prev < number) )
			errorIndex = i;
		
		prev = number;
	}
	cout << errorIndex;
	
	return 0;
}
