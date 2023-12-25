#include <iostream>
using namespace std;

int main()
{
	//Task("If3");
	int number;
	cin >> number;
	number += number > 0 ? -8 : number < 0 ? 6 : -number + 10;
	cout << number;
	
	return 0;
}
