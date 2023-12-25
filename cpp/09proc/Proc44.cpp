#include <iostream>
#include <cmath>
using namespace std;

double Arctg1(double x, double e);

int main()
{
	//Task("Proc44");
	double x;
	cin >> x;
	double e;
	for (int i = 1; i < 7; i++)
	{
		cin >> e;
		cout << Arctg1(x, e);
	}
	
	return 0;
}

double Arctg1(double x, double e)
{
	int factor = 1;
	int i = 1;
	double degreeX = x;
	double jsh = degreeX / i;
	double sum = 0;
	while (abs(jsh) > e)
	{
		sum += factor * jsh;
		factor *= -1;
		degreeX *= x*x;
		i += 2;
		jsh = degreeX / i;
	}
	return sum;
}