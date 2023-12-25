#include <iostream>
using namespace std;

int main()
{
	//Task("If26");
	double x, f;
	cin >> x;
	f = x <= 0 ? -x : x >= 2 ? 4 : x*x;
	cout << f;
	
	return 0;
}
