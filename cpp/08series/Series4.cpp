#include <iostream>
using namespace std;

int main()
{
	//Task("Series4");
	int N;
	cin >> N;
	double number, sum = 0, zarb = 1;
	for (int i = 1; i <= N; i++)
	{
		cin >> number;
		sum += number;
		zarb *= number;
	}
	cout << sum << zarb;
	
	return 0;
}
