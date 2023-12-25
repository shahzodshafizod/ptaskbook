#include <iostream>
#include <cmath>
using namespace std;

double Sin1(double x, double e);

int main()
{
	//Task("Proc41");
	double x;
	cin >> x;
	double e;
	for (int i = 1; i < 7; i++)
	{
		cin >> e;
		cout << Sin1(x, e);
	}
	
	return 0;
}

double Sin1(double x, double e)
{
	double degreeX = x;
	double fact = 1;
	double sum = 0;
	double factor = 1;
	double jsh = factor * degreeX / fact;
	int i = 1;
	while (abs(jsh) > e)
	{
		sum += factor * jsh;
		degreeX *= x * x;
		fact *= ++i;
		fact *= ++i;
		factor *= -1;
		jsh = degreeX / fact;
	}
	return sum;
}