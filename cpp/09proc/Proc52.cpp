#include <iostream>
using namespace std;

bool IsLeapYear(int Y);

int main()
{
	//Task("Proc52");
	int year;
	for (int i = 1; i < 6; i++)
	{
		cin >> year;
		cout << IsLeapYear(year);
	}
	
	return 0;
}

bool IsLeapYear(int Y)
{
	if (Y % 400 == 0)
		return true;
	
	else if (Y % 100 == 0)
		return false;

	else if (Y % 4 == 0)
		return true;

	else
		return false;
}