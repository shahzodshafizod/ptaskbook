#include <iostream>
#include <cmath>
using namespace std;

int main()
{
	//Task("If24");
	double x, f;
	cin >> x;
	f = x > 0 ? 2 * sin(x) : 6 - x;
	cout << f;
	
	return 0;
}
