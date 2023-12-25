#include <iostream>
using namespace std;

int main()
{
	//Task("Begin40");
	double A1, B1, C1, A2, B2, C2;
	cin >> A1 >> B1 >> C1 >> A2 >> B2 >> C2;
	double D = A1 * B2 - A2 * B1;
	double x = (C1 * B2 - C2 * B1) / D;
	double y = (A1 * C2 - A2 * C1) / D;
	cout << x << y;
	
	return 0;
}
