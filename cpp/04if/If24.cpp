#include <iostream>
#include <cmath>
using namespace std;

int main()
{
	//Task("If24");
	double x, f;
	cin >> x;
	if (x > 0)
		f = 2 * sin(x);
	else
		f = 6 - x;
	cout << f;
	
	return 0;
}
