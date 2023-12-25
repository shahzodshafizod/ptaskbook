#include <iostream>
using namespace std;

int main()
{
	//Task("Series2");
	double number, zarb = 1;
	for (int i = 1; i <= 10; i++)
	{
		cin >> number;
		zarb *= number;
	}
	cout << zarb;
	
	return 0;
}
