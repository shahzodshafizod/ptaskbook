#include <iostream>
using namespace std;

void NextDate(int& D, int& M, int& Y);
int MonthDays(int M, int Y);
bool IsLeapYear(int Y);

int main()
{
	//Task("Proc55");
	int d, m, y;
	for (int i = 1; i < 4; i++)
	{
		cin >> d >> m >> y;
		NextDate(d, m, y);
		cout << d << m << y;
	}
	
	return 0;
}

void NextDate(int& D, int& M, int& Y)
{
	if (D == MonthDays(M, Y))
	{
		if (M == 12)
		{
			Y++;
			M = 1;
		}
		else
			M++;
		D = 1;
	}
	else
		D++;
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
	}
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