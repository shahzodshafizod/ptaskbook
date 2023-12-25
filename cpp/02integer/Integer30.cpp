#include <iostream>
using namespace std;

int main()
{
	//Task("Integer30");
	int year;
	cin >> year;
	int century = (year - 1) / 100 + 1;
	cout << century;
	
	return 0;
}
