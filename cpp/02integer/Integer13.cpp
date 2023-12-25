#include <iostream>
using namespace std;

int main()
{
	//Task("Integer13");
	int number;
	cin >> number;
	int sadi = number / 100;
	int result = number % 100 * 10 + sadi;
	cout << result;
	
	return 0;
}
