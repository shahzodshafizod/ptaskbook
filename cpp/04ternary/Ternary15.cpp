#include <iostream>
using namespace std;

int main()
{
	//Task("If15");
	double a, b, c;
	cin >> a >> b >> c;
	
	double xurd = (a < b) && (a < c) ? a : b < c ? b : c;
	double sum = a + b + c - xurd;
	cout << sum;
	
	return 0;
}
