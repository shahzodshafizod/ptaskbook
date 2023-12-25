#include <iostream>
using namespace std;

int main()
{
	//Task("If12");
	double a, b, c;
	cin >> a >> b >> c;
	cout << ( (a < b) && (a < c) ? a : b < c ? b : c );
	
	return 0;
}
