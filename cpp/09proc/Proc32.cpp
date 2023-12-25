#include <iostream>
using namespace std;

double DegToRad(double D);

int main()
{
	//Task("Proc32");
	double D;
	for (int i = 0; i < 5; i++)
	{
		cin >> D;
		cout << DegToRad(D);
	}
	
	return 0;
}

double DegToRad(double D)
{
	return 3.14 * D / 180;
}