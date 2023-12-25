#include <iostream>
using namespace std;

int main()
{
	//Task("Series5");
	int N;
	cin >> N;
	double number, intPart, sum = 0;
	for (int i = 1; i <= N; i++)
	{
		cin >> number;
		intPart = (int)number;
		sum += intPart;
		cout << intPart;
	}
	cout << sum;
	
	return 0;
}
