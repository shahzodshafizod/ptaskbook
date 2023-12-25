#include <iostream>
using namespace std;

int Sign(double X);

int main()
{
	//Task("Proc16");
	double A, B;
	cin >> A >> B;
	cout << ( Sign(A) + Sign(B) );
	
	return 0;
}

int Sign(double X)
{
	return ( X < 0 ? -1 : X == 0 ? 0 : 1 );
}