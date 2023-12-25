#include <iostream>
using namespace std;

int main()
{
	//Task("Boolean21");
	int number;
	cin >> number;
	int a = number / 100;
	int b = number / 10 % 10;
	int c = number % 10;
	bool progress = (a < b) && (b < c);
	cout << progress;
	
	return 0;
}
