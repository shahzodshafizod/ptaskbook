#include <iostream>
using namespace std;

int main()
{
	//Task("If28");
	int year;
	cin >> year;
	int days = year % 400 == 0 ? 366 : year % 100 == 0 ? 365 : year % 4 == 0 ? 366 : 365;
	cout << days;
	
	return 0;
}
