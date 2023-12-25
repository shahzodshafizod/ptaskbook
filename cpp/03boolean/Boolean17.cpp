#include <iostream>
using namespace std;

int main()
{
	//Task("Boolean17");
	int number;
	cin >> number;
	bool unevenThreeDigital = (number % 2 != 0) && (number > 99) && (number < 1000);
	cout << unevenThreeDigital;
	
	return 0;
}
