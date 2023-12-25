#include <iostream>
#include <cmath>
using namespace std;

double Rings(double R1, double R2);

int main()
{
	//Task("Proc19");
	double R1, R2;
	for (int i = 1; i <= 3; i++)
	{
		cin >> R1 >> R2;
		cout << Rings(R1, R2);
	}
	
	return 0;
}

double Rings(double R1, double R2)
{
	const double PI = 3.14;
	double S1 = PI * R1 * R1;
	double S2 = PI * R2 * R2;

	return abs(S1 - S2);
}