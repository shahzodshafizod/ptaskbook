#include <iostream>
using namespace std;

int main()
{
	//Task("Integer10");
	int number;
	cin >> number;
	int vohid = number % 10;
	int dahi = number % 100 / 10;
	cout << vohid << dahi;
	
	return 0;
}
