#include <iostream>
using namespace std;

void SortInc3(double& A, double& B, double& C);
void Minmax(double& a, double& b);
void Swap(double& x, double& y);

int main()
{
	//Task("Proc12");
	double A, B, C;

	for (int i = 0; i < 2; i++)
	{
		cin >> A >> B >> C;
		SortInc3(A, B, C);
		cout << A << B << C;
	}
	
	return 0;
}

void SortInc3(double& A, double& B, double& C)
{
	Minmax(A, B);
	Minmax(A, C);
	Minmax(B, C);
}

void Minmax(double& a, double& b)
{
	if (a > b)
		Swap(a, b);
}

void Swap(double& x, double& y)
{
	x += y;
	y = x - y;
	x -= y;
}