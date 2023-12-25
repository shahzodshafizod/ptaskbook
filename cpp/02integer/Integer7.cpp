#include <iostream>
using namespace std;

int main()
{
	//Task("Integer7");
	int number;
	cin >> number;
	int a = number / 10;
	int b = number % 10;
	int sum = a+b;
	int zarb = a*b;
	cout << sum << zarb;
	
	return 0;
}
