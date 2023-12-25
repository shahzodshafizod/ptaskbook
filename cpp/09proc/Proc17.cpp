#include <iostream>
using namespace std;

int RootCount(double A, double B, double C);

int main()
{
	//Task("Proc17");
	double A, B, C;
	for (int i = 0; i < 3; i++)
	{
		cin >> A >> B >> C;
		cout << RootCount(A, B, C);
	}
	
	return 0;
}

int RootCount(double A, double B, double C)
{
	double D = B*B - 4*A*C;

	return ( D > 0 ? 2 : D == 0 ? 1 : 0 );
}