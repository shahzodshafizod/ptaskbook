#include <iostream>
using namespace std;

double Exp1(double x, double e);

int main()
{
	//Task("Proc40");
	double x;
	cin >> x;
	
	double e;
	for (int i = 1; i <= 6; i++)
	{
		cin >> e;
		cout << Exp1(x, e);
	}
	
	return 0;
}

double Exp1(double x, double e)
{
	double jsh = 0, sum = 0;
	double deg = 1, fact = 1;
	jsh = deg / fact;
	int i = 1;
	while (jsh > e)
	{
		sum += jsh;
		deg *= x;
		fact *= i++;
		jsh = deg / fact;
	}
	
	return sum;
}