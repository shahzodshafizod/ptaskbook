#include <iostream>
using namespace std;

void ShiftLeft3(double& A, double& B, double& C);

int main()
{
	//Task("Proc15");
	double A, B, C;
	for (int i = 2; i > 0; i--)
	{
		cin >> A >> B >> C;
		ShiftLeft3(A, B, C);
		cout << A << B << C;
	}
	
	return 0;
}

void ShiftLeft3(double& A, double& B, double& C)
{
	double temp = A;
	A = B;
	B = C;
	C = temp;
}