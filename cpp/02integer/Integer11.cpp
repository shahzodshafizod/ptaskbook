#include <iostream>
using namespace std;

int main()
{
	//Task("Integer11");
	int number;
	cin >> number;
	int a = number / 100;
	int b = number / 10 % 10;
	int c = number % 10;
	int sum = a + b + c;
	int zarb = a * b * c;
	cout << sum << zarb;
	
	return 0;
}
