#include <iostream>
using namespace std;

int main()
{
	//Task("Case15");
	int N, M;
	cin >> N >> M;
	string str = "";
	switch (N)
	{
		case 6: str += "шестерка"; break;
		case 7: str += "семерка"; break;
		case 8: str += "восьмерка"; break;
		case 9: str += "девятка"; break;
		case 10: str += "десятка"; break;
		case 11: str += "валет"; break;
		case 12: str += "дама"; break;
		case 13: str += "король"; break;
		case 14: str += "туз"; break;
	}
	str += " ";
	switch (M)
	{
		case 1: str += "пик"; break;
		case 2: str += "треф"; break;
		case 3: str += "бубен"; break;
		case 4: str += "червей"; break;
	}
	
	cout << str;
	
	return 0;
}
