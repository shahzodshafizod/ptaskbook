#include <iostream>
using namespace std;

int main()
{
	//Task("If29");
	int number;
	cin >> number;
	string str = "";
	if (number != 0)
	{
		if (number > 0)
			str += "положи";
		else
			str += "отрица";
		
		str += "тельное ";
		if (number % 2 != 0)
			str += "не";
		str += "четное";
	}
	else
		str += "нулевое";
	str += " число";

	cout << str;
	
	return 0;
}
