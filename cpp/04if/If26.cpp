#include <iostream>
using namespace std;

int main()
{
	//Task("If26");
	double x, f;
	cin >> x;
	if (x <= 0)
		f = -x;
	else if (x >= 2)
		f = 4;
	else
		f = x*x;
	cout << f;
	
	return 0;
}
