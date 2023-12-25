#include <iostream>
using namespace std;

int main()
{
	//Task("If9");
	double A, B;
	cin >> A >> B;
	double xurd = A < B ? A : B;
	B = A > B ? A : B;
	A = xurd;

	cout << A << B;
	
	return 0;
}
