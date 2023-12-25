#include <iostream>
using namespace std;

int main()
{
	//Task("Integer14");
	int number;
	cin >> number;
	int vohid = number % 10;
	int result = vohid * 100 + number / 10;
	cout << result;
	
	return 0;
}
