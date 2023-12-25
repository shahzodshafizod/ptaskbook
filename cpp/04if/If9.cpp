#include <iostream>
using namespace std;

int main()
{
	//Task("If9");
	double A, B;
	cin >> A >> B;
	if (A > B)
	{
		double temp = A;
		A = B;
		B = temp;
	}
	cout << A << B;
	
	return 0;
}
