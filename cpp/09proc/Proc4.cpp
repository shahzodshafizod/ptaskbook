#include <iostream>
#include <cmath>
using namespace std;

void TrianglePS(double a, double& P, double& S);

int main()
{
    Task("Proc4");
	double a, P, S;
	for (int i = 1; i <= 3; i++)
	{
		cin >> a;
		TrianglePS(a, P, S);
		cout << P << S;
	}
	
	return 0;
}

void TrianglePS(double a, double& P, double& S)
{
	P = 3*a;
	S = a*a * sqrt(3.0) / 4;
}