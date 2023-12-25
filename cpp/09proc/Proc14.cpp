#include <iostream>
using namespace std;

void ShiftRight3(double& A, double& B, double& C);

int main()
{
	//Task("Proc14");
	double a, b, c;
	for (int i = 1; i <= 2; i++)
	{
		cin >> a >> b >> c;
		ShiftRight3(a, b, c);
		cout << a << b << c;
	}
	
	return 0;
}

void ShiftRight3(double& A, double& B, double& C)
{
	double temp = A;
	A = C;
	C = B;
	B = temp;
}