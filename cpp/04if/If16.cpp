#include <iostream>
using namespace std;

int main()
{
	//Task("If16");
	double A, B, C;
	cin >> A >> B >> C;
	if ( (A < B) && (B < C) )
	{
		A *= 2;
		B *= 2;
		C *= 2;
	}
	else
	{
		A *= -1;
		B *= -1;
		C *= -1;
	}
	cout << A << B << C;
	
	return 0;
}
