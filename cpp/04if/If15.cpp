#include <iostream>
using namespace std;

int main()
{
	//Task("If15");
	double a, b, c;
	cin >> a >> b >> c;
	
	double xurd;
	if ( (a < b) && (a < c) )
		xurd = a;
	else if (b < c)
		xurd = b;
	else
		xurd = c;
	
	double sum = a + b + c - xurd;
	cout << sum;
	
	return 0;
}
