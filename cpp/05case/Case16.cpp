#include <iostream>
using namespace std;

int main()
{
	//Task("Case16");
	int number;
	cin >> number;
	int dahi = number / 10;
	int vohid = number % 10;
	string str = "";
	switch (dahi)
	{
		case 2: str += "двадцать"; break;
		case 3: str += "тридцать"; break;
		case 4: str += "сорок"; break;
		case 5: str += "пятьдесят"; break;
		case 6: str += "шестьдесят"; break;
	}
	if (vohid != 0)
		str += " ";
	switch (vohid)
	{
		case 1: str += "один"; break;
		case 2: str += "два"; break;
		case 3: str += "три"; break;
		case 4: str += "четыре"; break;
		case 5: str += "пять"; break;
		case 6: str += "шесть"; break;
		case 7: str += "семь"; break;
		case 8: str += "восемь"; break;
		case 9: str += "девять"; break;
	}
	str += " ";
	switch (vohid)
	{
		case 1:
			str += "год"; break;
		case 2: case 3: case 4:
			str += "года"; break;
		default:
			str += "лет"; break;
	}
	cout << str;
	
	return 0;
}
