#include <iostream>
using namespace std;

struct TDate
{
	int Day;
	int Month;
	int Year;
};

void NextDate(TDate& D);
int CheckDate(TDate D);
int  DaysInMonth(TDate D);
bool LeapYear(TDate D);

int main()
{
	//Task("Param63");
	TDate date;
	for (int i = 0; i < 5; i++)
	{
		cin >> date.Day >> date.Month >> date.Year;
		NextDate(date);
		cout << date.Day << date.Month << date.Year;
	}
	
	return 0;
}

void NextDate(TDate& D)
{
	if (CheckDate(D) != 0)
		return;

	if (D.Day == DaysInMonth(D))
	{
		if (D.Month == 12)
		{
			D.Month = 1;
			D.Year++;
		}
		else
			D.Month++;
		D.Day = 1;
	}
	else
		D.Day++;
}

int CheckDate(TDate D)
{
	if ( (D.Month < 1) || (D.Month > 12) )
		return 1;

	else if ( (D.Day < 1) || (D.Day > DaysInMonth(D)) )
		return 2;

	return 0;
}

int  DaysInMonth(TDate D)
{
	switch (D.Month)
	{
		case 1: case 3: case 5: case 7: case 8: case 10: case 12:
			return 31;
		case 4: case 6: case 9: case 11:
			return 30;
		case 2:
			return (LeapYear(D) ? 29 : 28);
	}
}

bool LeapYear(TDate D)
{
	if (D.Year % 400 == 0)
		return true;

	if (D.Year % 100 == 0)
		return false;

	if (D.Year % 4 == 0)
		return true;

	return false;
}