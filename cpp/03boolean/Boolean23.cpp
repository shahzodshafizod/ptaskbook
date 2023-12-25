#include <iostream>
using namespace std;

int main()
{
	//Task("Boolean23");
	int number;
	cin >> number;
	int hazori = number / 1000;
	int sadi = number / 100 % 10;
	int dahi = number / 10 % 10;
	int vohid = number % 10;
	bool palindrome = (hazori == vohid) && (sadi == dahi);
	cout << palindrome;
	
	return 0;
}
