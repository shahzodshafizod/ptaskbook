#include <iostream>
#include <cmath>
using namespace std;

int main()
{
	//Task("Begin39");
	double A, B, C;
	cin >> A >> B >> C;
	double D = B*B - 4*A*C;
	double x1 = (-B - sqrt(D)) / (2*A);
	double x2 = (-B + sqrt(D)) / (2*A);
	cout << x1 << x2;
	
	return 0;
}
