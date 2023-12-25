#include <iostream>
using namespace std;

int main()
{
	//Task("If13");
	double a, b, c;
	cin >> a >> b >> c;
	
	double kalon = (a > b) && (a > c) ? a : b > c ? b : c;
	double xurd = (a < b) && (a < c) ? a : b < c ? b : c;
	double bayn = a + b + c - kalon - xurd;
	cout << bayn;
	
	return 0;
}
