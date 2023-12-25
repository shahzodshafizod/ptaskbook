#include <iostream>
#include <cmath>
using namespace std;

double Power4(double x, double a, double e);

int main()
{
	//Task("Proc45");
	double x, a;
	cin >> x >> a;
	double e;
	for (int i = 1; i <= 6; i++)
	{
		cin >> e;
		cout << Power4(x, a, e);
	}
	
	return 0;
}

double Power4(double x, double a, double e)
{
	int i = 0;
	double A = 1;
	double degreeX = 1;
	double fact = 1;
	double jsh = A * degreeX / fact;
	double sum = 0;
	while (abs(jsh) > e)
	{
		sum += jsh;
		A *= a-i;
		degreeX *= x;
		fact *= ++i;
		jsh = A * degreeX / fact;
	}
	return sum;
}