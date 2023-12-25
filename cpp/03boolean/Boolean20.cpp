#include <iostream>
using namespace std;

int main()
{
	//Task("Boolean20");
	int number;
	cin >> number;
	int a = number / 100;
	int b = number / 10 % 10;
	int c = number % 10;
	bool different = (a != b) && (a != c) && (b != c);
	cout << different;
	
	return 0;
}
