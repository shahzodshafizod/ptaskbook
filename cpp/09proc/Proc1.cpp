#include <iostream>
using namespace std;

void PowerA3(double A, double& B);

int main()
{
	//Task("Proc1");
	double A, B;
	for (int i = 1; i <= 5; i++)
	{
		cin >> A;
		PowerA3(A, B);
		cout << B;
	}
	
	return 0;
}

void PowerA3(double A, double& B)
{
	B = A*A*A;
}