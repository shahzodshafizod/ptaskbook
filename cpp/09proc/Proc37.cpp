#include <iostream>
#include <cmath>
using namespace std;

double Power1(double A, double B);

int main()
{
	//Task("Proc37");
	double P, number;
	cin >> P;
	for (int i = 1; i < 4; i++)
	{
		cin >> number;
		cout << Power1(number, P);
	}
	
	return 0;
}

double Power1(double A, double B)
{
	if (A <= 0)
		return 0;
	return exp(B * log(A));
}