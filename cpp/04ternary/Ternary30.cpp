#include <iostream>
using namespace std;

int main()
{
	//Task("If30");
	int number;
	cin >> number;
	string str = "";
	str += number % 2 != 0 ? "не" : "";
	str += "четное ";
	str += number <= 9 ? "одно" : number <= 99 ? "дву" : "трех";
	str += "значное число";
	cout << str;
	
	return 0;
}
