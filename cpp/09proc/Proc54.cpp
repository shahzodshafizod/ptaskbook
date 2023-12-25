#include <iostream>
using namespace std;

bool IsLeapYear(int Y);
int MonthDays(int M, int Y);
void PrevDate(int& D, int& M, int& Y);

int main()
{
	//Task("Proc54");
	int d, m, y;
	for (int i = 1; i <= 3; i++)
	{
		cin >> d >> m >> y;
		PrevDate(d, m, y);
		cout << d << m << y;
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

int MonthDays(int M, int Y)
{
	switch (M)
	{
		case 1: case 3: case 5: case 7: case 8: case 10: case 12:
			return 31;
		case 4: case 6: case 9: case 11:
			return 30;
		case 2:
			return ( IsLeapYear(Y) ? 29 : 28 );
		default:
			return 0;
	}
}

void PrevDate(int& D, int& M, int& Y)
{
	if (D == 1)
	{
		if (M == 1)
		{
			Y--;
			M = 12;
		}
		else
			M--;
		
		D = MonthDays(M, Y);
	}
	else
		D--;
}