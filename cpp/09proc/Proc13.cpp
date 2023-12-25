#include <iostream>
using namespace std;

void SortDec3(double& A, double& B, double& C);
void Maxmin(double& a, double& b);
void Swap(double& x, double& y);

int main()
{
	//Task("Proc13");
	double A, B, C;

	for (int i = 2; i > 0; i--)
	{
		cin >> A >> B >> C;
		SortDec3(A, B, C);
		cout << A << B << C;
	}
	
	return 0;
}

void SortDec3(double& A, double& B, double& C)
{
	Maxmin(A, B);
	Maxmin(A, C);
	Maxmin(B, C);
}

void Maxmin(double& a, double& b)
{
	if (a < b)
		Swap(a, b);
}

void Swap(double& x, double& y)
{
	x += y;
	y = x - y;
	x -= y;
}