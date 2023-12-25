#include <iostream>
#include <cmath>
using namespace std;

struct TPoint
{
	double X, Y;
};

struct TTriangle
{
	TPoint A, B, C;
};

double Leng(TPoint A, TPoint B);
double Perim(TTriangle T);

int main()
{
	//Task("Param65");
	TPoint A, B, C, D;
	
	cin >> A.X >> A.Y;
	cin >> B.X >> B.Y;
	cin >> C.X >> C.Y;
	cin >> D.X >> D.Y;
	
	TTriangle sekunja;

	sekunja.A = A; sekunja.B = B; sekunja.C = C;
	cout << Perim(sekunja);

	sekunja.C = D;
	cout << Perim(sekunja);

	sekunja.B = C;
	cout << Perim(sekunja);
	
	return 0;
}

double Leng(TPoint A, TPoint B)
{
	return sqrt( pow(A.X - B.X, 2) + pow(A.Y - B.Y, 2) );
}

double Perim(TTriangle T)
{
	double a = Leng(T.A, T.B);
	double b = Leng(T.B, T.C);
	double c = Leng(T.A, T.C);
	
	return a+b+c;
}