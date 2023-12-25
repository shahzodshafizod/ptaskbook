#include <iostream>
using namespace std;

int main()
{
	//Task("If22");
	double x, y;
	cin >> x >> y;
	int quarter = 0;
	if (x > 0)
	{
		if (y > 0)
			quarter = 1;
		else
			quarter = 4;
	}
	else
	{
		if (y > 0)
			quarter = 2;
		else
			quarter = 3;
	}
	cout << quarter;
	
	return 0;
}
