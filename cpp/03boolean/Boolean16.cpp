#include <iostream>
using namespace std;

int main()
{
	//Task("Boolean16");
	int number;
	cin >> number;
	bool evenTwoDigital = (number % 2 == 0) && (number > 9) && (number < 100);
	cout << evenTwoDigital;
	
	return 0;
}
