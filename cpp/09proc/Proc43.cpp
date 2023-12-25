#include <iostream>
#include <cmath>
using namespace std;

double Ln1(double x, double e);

int main()
{
	//Task("Proc43");
	double x;
	cin >> x;
	double e;
	for (int i = 1; i < 7; i++)
	{
		cin >> e;
		cout << Ln1(x, e);
	}
	
	return 0;
}

double Ln1(double x, double e)
{
	double degreeX = x;
	int i = 1;
	double jsh = degreeX / i;
	int factor = 1;
	double sum = 0;
	while (abs(jsh) > e)
	{
		sum += factor * jsh;
		degreeX *= x;
		i++;
		jsh = degreeX / i;
		factor *= -1;
	}
	return sum;
}