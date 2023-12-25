#include <iostream>
#include <cmath>
using namespace std;

double TriangleP(double a, double h);

int main()
{
	//Task("Proc20");
	double a, h;
	for (int i = 1; i <= 3; i++)
	{
		cin >> a >> h;
		cout << TriangleP(a, h);
	}
	
	return 0;
}

double TriangleP(double a, double h)
{
	double b = sqrt( pow(a/2, 2) + h*h );

	return 2*b + a;
}