#include <iostream>
using namespace std;

int main()
{
	//Task("Integer8");
	int number;
	cin >> number;
	int dahi = number / 10;
	int vohid = number % 10;
	int chappa = vohid * 10 + dahi;
	cout << chappa;
	
	return 0;
}
