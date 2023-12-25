#include <iostream>
using namespace std;

int main()
{
	//Task("Boolean22");
	int number;
	cin >> number;
	int a = number / 100;
	int b = number / 10 % 10;
	int c = number % 10;
	bool progReg = (a < b) && (b < c) || (a > b) && (b > c);
	cout << progReg;
	
	return 0;
}
