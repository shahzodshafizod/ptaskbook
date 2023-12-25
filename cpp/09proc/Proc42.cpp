#include <iostream>
#include <cmath>
using namespace std;

double Cos1(double x, double e);

int main()
{
	//Task("Proc42");
	double x;
	cin >> x;
	double e;
	for (int i = 1; i < 7; i++)
	{
		cin >> e;
		cout << Cos1(x, e);
	}
	
	return 0;
}

double Cos1(double x, double e)
{
	double degreeX = 1;
	double fact = 1;
	int factor = 1;
	double jsh = factor * degreeX / fact;
	double sum = 0;
	int i = 0;
	while (abs(jsh) > e)
	{
		sum += factor * jsh;
		degreeX *= x*x;
		fact *= ++i;
		fact *= ++i;
		factor *= -1;
		jsh = degreeX / fact;
	}
	return sum;
}