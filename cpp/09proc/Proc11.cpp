#include <iostream>
using namespace std;

void Minmax(double& X, double& Y);
void Swap(double& A, double& B);

int main()
{
	//Task("Proc11");
	double a, b, c, d;
	cin >> a >> b >> c >> d;

	Minmax(a, b);
	Minmax(c, d);
	Minmax(a, c);
	Minmax(b, d);

	cout << a << d;
	
	return 0;
}

void Minmax(double& X, double& Y)
{
	if (X > Y)
		Swap(X, Y);
}

void Swap(double& A, double& B)
{
	A += B;
	B = A - B;
	A -= B;
}