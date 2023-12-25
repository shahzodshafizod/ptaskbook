#include <iostream>
#include <cmath>
using namespace std;

struct TPoint
{
	double X;
	double Y;
};

double Leng(TPoint A, TPoint B);

int main()
{
	//Task("Param64");
	TPoint A, B;
	cin >> A.X >> A.Y;
	for (int i = 1; i < 4; i++)
	{
		cin >> B.X >> B.Y;
		cout << Leng(A, B);
	}
	
	return 0;
}

double Leng(TPoint A, TPoint B)
{
	return sqrt( pow(A.X - B.X, 2) + pow(A.Y - B.Y, 2) );
}