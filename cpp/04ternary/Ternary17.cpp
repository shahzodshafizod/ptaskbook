#include <iostream>
using namespace std;

int main()
{
	//Task("If17");
	double A, B, C;
	cin >> A >> B >> C;
	int factor = (A < B) && (B < C) || (A > B) && (B > C) ? 2 : -1;
	A *= factor;
	B *= factor;
	C *= factor;
	cout << A << B << C;
	
	return 0;
}
