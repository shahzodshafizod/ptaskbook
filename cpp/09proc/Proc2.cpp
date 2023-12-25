#include <iostream>
using namespace std;

void PowerA234(double A, double& B, double& C, double& D);

int main()
{
	//Task("Proc2");
	double a, b, c, d;
	for (int i = 0; i < 5; i++)
	{
		cin >> a;
		PowerA234(a, b, c, d);
		cout << b << c << d;
	}
	
	return 0;
}

void PowerA234(double A, double& B, double& C, double& D)
{
	B = A*A;
	C = B*A;
	D = C*A;
}