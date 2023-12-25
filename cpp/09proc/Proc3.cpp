#include <iostream>
#include <cmath>
using namespace std;

void Mean(double X, double Y, double& AMean, double& GMean);

int main()
{
    Task("Proc3");
	double a, b, c, d, am, gm;
	cin >> a >> b >> c >> d;

	Mean(a, b, am, gm);
	cout << am << gm;

	Mean(a, c, am, gm);
	cout << am << gm;

	Mean(a, d, am, gm);
	cout << am << gm;
	
	return 0;
}

void Mean(double X, double Y, double& AMean, double& GMean)
{
	AMean = (Y+X)/2;
	GMean = sqrt(Y*X);
}