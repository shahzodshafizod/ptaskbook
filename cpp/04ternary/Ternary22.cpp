#include <iostream>
using namespace std;

int main()
{
	//Task("If22");
	double x, y;
	cin >> x >> y;
	int quarter = x > 0 ? y > 0 ? 1 : 4 : y > 0 ? 2 : 3;
	cout << quarter;
	
	return 0;
}
