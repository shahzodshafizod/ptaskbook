#include <iostream>
using namespace std;

struct TDate
{
	int Day;
	int Month;
	int Year;
};

bool LeapYear(TDate D);

int main()
{
	//Task("Param59");
	TDate date;
	for (int i = 0; i < 5; i++)
	{
		cin >> date.Day >> date.Month >> date.Year;
		cout << LeapYear(date);
	}
	
	return 0;
}

bool LeapYear(TDate D)
{
	if (D.Year % 400 == 0)
		return true;

	else if (D.Year % 100 == 0)
		return false;

	else if (D.Year % 4 == 0)
		return true;

	else
		return false;
}