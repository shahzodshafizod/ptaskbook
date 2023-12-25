#include <iostream>
using namespace std;

int Quarter(double x, double y);

int main()
{
	//Task("Proc23");
	double x, y;

	for (int i = 0; i < 3; i++)
	{
		cin >> x >> y;
		cout << Quarter(x, y);
	}
	
	return 0;
}

int Quarter(double x, double y)
{
	return ( x > 0 ? y > 0 ? 1 : 4 : y > 0 ? 2 : 3 );
}