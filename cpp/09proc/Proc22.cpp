#include <iostream>
using namespace std;

double Calc(double A, double B, int Op);

int main()
{
	//Task("Proc22");
	double A, B;
	cin >> A >> B;

	int N;
	for (int i = 1; i <= 3; i++)
	{
		cin >> N;
		cout << Calc(A, B, N);
	}
	
	return 0;
}

double Calc(double A, double B, int Op)
{
	switch (Op)
	{
		case 1:		return A-B;
		case 2:		return A*B;
		case 3:		return A/B;
		default:	return A+B;
	}
}