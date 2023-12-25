#include <iostream>
using namespace std;

struct TDate
{
	int Day;
	int Month;
	int Year;
};

int DaysInMonth(TDate D);
bool LeapYear(TDate D);
int CheckDate(TDate D);
void PrevDate(TDate& D);

int main()
{
	//Task("Param62");
	TDate date;
	for (int i = 1; i < 6; i++)
	{
		cin >> date.Day >> date.Month >> date.Year;
		PrevDate(date);
		cout << date.Day << date.Month << date.Year;
	}
	
	return 0;
}

int DaysInMonth(TDate D)
{
	switch (D.Month)
	{
		case 1: case 3: case 5: case 7: case 8: case 10: case 12:
			return 31;
		case 4: case 6: case 9: case 11:
			return 30;
		case 2:
			return (LeapYear(D) ? 29 : 28);
		default:
			return 0;
	}
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

int CheckDate(TDate D)
{
	if ( (D.Month < 1) || (D.Month > 12) )
		return 1;

	else if ( (D.Day < 1) || (D.Day > DaysInMonth(D)) )
		return 2;

	return 0;
}

void PrevDate(TDate& D)
{
	if (CheckDate(D) != 0)
		return;

	if (D.Day == 1)
	{
		if (D.Month == 1)
		{
			D.Month = 12;
			D.Year--;
		}
		else
			D.Month--;
		D.Day = DaysInMonth(D);
	}
	else
		D.Day--;
}