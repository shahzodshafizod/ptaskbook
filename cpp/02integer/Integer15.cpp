#include <iostream>
using namespace std;

int main()
{
	//Task("Integer15");
	int number;
	cin >> number;
	int a = number / 100;
	int b = number / 10 % 10;
	int c = number % 10;
	int result = b * 100 + a * 10 + c;
	cout << result;
	
	return 0;
}
