#include <iostream>
using namespace std;

int main()
{
	//Task("If14");
	double a, b, c;
	cin >> a >> b >> c;
	
	double xurd = (a < b) && (a < c) ? a : b < c ? b : c;
	double kalon = (a > b) && (a > c) ? a : b > c ? b : c;

	cout << xurd << kalon;
	
	return 0;
}
