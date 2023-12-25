#include <iostream>
using namespace std;

int main()
{
	//Task("Case19");
	int year;
	cin >> year;
	year -= 4;
	if (year < 0)
		year *= -1;
	
	int n = (year) % 60;
	string str = "год ";
	
	if (n < 12)
		str += "зелен";
	else if (n < 24)
		str += "красн";
	else if (n < 36)
		str += "желт";
	else if (n < 48)
		str += "бел";
	else
		str += "черн";

	n %= 12;
	switch (n)
	{
		case 0: str += "ой крысы"; break;
		case 1: str += "ой коровы"; break;
		case 2: str += "ого тигра"; break;
		case 3: str += "ого зайца"; break;
		case 4: str += "ого дракона"; break;
		case 5: str += "ой змеи"; break;
		case 6: str += "ой лошади"; break;
		case 7: str += "ой овцы"; break;
		case 8: str += "ой обезьяны"; break;
		case 9: str += "ой курицы"; break;
		case 10: str += "ой собаки"; break;
		case 11: str += "ой свиньи"; break;
	}

	cout << str;
	
	return 0;
}
