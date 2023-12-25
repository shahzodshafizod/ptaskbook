#include <iostream>
using namespace std;

int main()
{
	//Task("If28");
	int year;
	cin >> year;
	int days;
	if (year % 400 == 0)
		days = 366;
	else if (year % 100 == 0)
		days = 365;
	else if (year % 4 == 0)
		days = 366;
	else
		days = 365;
	cout << days;
	
	return 0;
}
