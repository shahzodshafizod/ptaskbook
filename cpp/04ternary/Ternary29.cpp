#include <iostream>
using namespace std;

int main()
{
	//Task("If29");
	int number;
	cin >> number;
	string str = "";
	str += number != 0 ? number > 0 ? "положительное" : "отрицательное" : "нулевое";
	str += number != 0 ? number % 2 != 0 ? " нечетное" : " четное" : "";
	str += " число";
	cout << str;
	
	return 0;
}
