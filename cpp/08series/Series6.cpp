#include <iostream>
using namespace std;

int main()
{
	//Task("Series6");
	int N;
	cin >> N;
	double number, precisionPart, zarb = 1;
	for (int i = 1; i <= N; i++)
	{
		cin >> number;
		precisionPart = number - (int)number;
		zarb *= precisionPart;
		cout << precisionPart;
	}
	cout << zarb;
	
	return 0;
}
