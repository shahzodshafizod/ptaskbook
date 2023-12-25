#include <iostream>
using namespace std;

int main()
{
	//Task("Case17");
	int number;
	cin >> number;
	int dahi = number / 10;
	int vohid = number % 10;
	string str = "";
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
	}
	if ( (vohid != 0) && (dahi != 1) )
	{
		str += " ";
		switch (vohid)
		{
			case 1: str += "одно"; break;
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
	str += " ";
	if (dahi != 1)
	{
		switch (vohid)
		{
			case 1:
				str += "учебное задание"; break;
			case 2: case 3: case 4:
				str += "учебных задания"; break;
			default:
				str += "учебных заданий"; break;
		}
	}
	else str += "учебных заданий";
	
	cout << str;
	
	return 0;
}
