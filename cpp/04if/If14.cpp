#include <iostream>
using namespace std;

int main()
{
	//Task("If14");
	double a, b, c;
	cin >> a >> b >> c;
	
	double xurd;
	if ( (a < b) && (a < c) )
		xurd = a;
	else if (b < c)
		xurd = b;
	else
		xurd = c;
	
	double kalon;
	if ( (a > b) && (a > c) )
		kalon = a;
	else if (b > c)
		kalon = b;
	else
		kalon = c;

	cout << xurd << kalon;
	
	return 0;
}
