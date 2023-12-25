#include <iostream>
using namespace std;

int main()
{
	//Task("If13");
	double a, b, c;
	cin >> a >> b >> c;
	
	double kalon;
	if ( (a > b) && (a > c) )
		kalon = a;
	else if (b > c)
		kalon = b;
	else
		kalon = c;
	
	double xurd;
	if ( (a < b) && (a < c) )
		xurd = a;
	else if (b < c)
		xurd = b;
	else
		xurd = c;
	
	double bayn = a + b + c - kalon - xurd;
	cout << bayn;
	
	return 0;
}
