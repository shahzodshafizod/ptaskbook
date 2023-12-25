#include <iostream>
using namespace std;

int main()
{
	//Task("Integer12");
	int number;
	cin >> number;
	int sadi = number / 100;
	int dahi = number / 10 % 10;
	int vohid = number % 10;
	int chappa = vohid * 100 + dahi * 10 + sadi;
	cout << chappa;
	
	return 0;
}
