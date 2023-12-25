#include <iostream>
using namespace std;

int main()
{
	//Task("Case18");
	int number;
	cin >> number;
	int sadi = number / 100;
	int dahi = number / 10 % 10;
	int vohid = number % 10;
	string str = "";
	switch (sadi)
	{
		case 1: str += "сто"; break;
		case 2: str += "двести"; break;
		case 3: str += "триста"; break;
		case 4: str += "четыреста"; break;
		case 5: str += "пятьсот"; break;
		case 6: str += "шестьсот"; break;
		case 7: str += "семьсот"; break;
		case 8: str += "восемьсот"; break;
		case 9: str += "девятьсот"; break;
	}
	if (dahi != 0)
	{
		str += " ";
		switch (dahi)
		{
			case 1:
				switch (vohid)
				{
					case 0: str += "десять"; break;
					case 1: str += "одиннадцать"; break;
					case 2: str += "двенадцать"; break;
					case 3: str += "тринадцать"; break;
					case 4: str += "четырнадцать"; break;
					case 5: str += "пятнадцать"; break;
					case 6: str += "шестнадцать"; break;
					case 7: str += "семнадцать"; break;
					case 8: str += "восемнадцать"; break;
					case 9: str += "девятнадцать"; break;
				}
				break;
			case 2: str += "двадцать"; break;
			case 3: str += "тридцать"; break;
			case 4: str += "сорок"; break;
			case 5: str += "пятьдесят"; break;
			case 6: str += "шестьдесят"; break;
			case 7: str += "семьдесят"; break;
			case 8: str += "восемьдесят"; break;
			case 9: str += "девяносто"; break;
		}
	}
	if ( (vohid != 0) && (dahi != 1) )
	{
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
	}
	cout << str;
	
	return 0;
}
