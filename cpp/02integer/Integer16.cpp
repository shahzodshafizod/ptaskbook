#include <iostream>
using namespace std;

int main()
{
	//Task("Integer16");
	int number;
	cin >> number;
	int a = number / 100;
	int b = number / 10 % 10;
	int c = number % 10;
	int result = a * 100 + c * 10 + b;
	cout << result;
	
	return 0;
}
