#include <iostream>
using namespace std;

void Swap(double& X, double& Y);

int main()
{
    //Task("Proc10");
	double a, b, c, d;
	cin >> a >> b >> c >> d;

	Swap(a, b);
	Swap(c, d);
	Swap(b, c);

	cout << a << b << c << d;
	
	return 0;
}

void Swap(double& X, double& Y)
{
	/*
	X += Y;
	Y = X - Y;
	X -= Y;
	*/

	double temp = X;
	X = Y;
	Y = temp;
}