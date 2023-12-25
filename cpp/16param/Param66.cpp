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

double Area(TTriangle T);
double Perim(TTriangle T);
double Leng(TPoint A, TPoint B);

int main()
{
	//Task("Param66");
	TPoint A, B, C, D;

	cin >> A.X >> A.Y;
	cin >> B.X >> B.Y;
	cin >> C.X >> C.Y;
	cin >> D.X >> D.Y;

	TTriangle sekunja;

	sekunja.A = A; sekunja.B = B; sekunja.C = C;
	cout << Area(sekunja);

	sekunja.C = D;
	cout << Area(sekunja);

	sekunja.B = C;
	cout << Area(sekunja);
	
	return 0;
}

double Area(TTriangle T)
{
	double a = Leng(T.A, T.B);
	double b = Leng(T.B, T.C);
	double c = Leng(T.A, T.C);

	//double p = Perim(T)/2;
	double p = (a+b+c) / 2;

	return sqrt( p*(p-a)*(p-b)*(p-c) );
}

double Perim(TTriangle T)
{
	double a = Leng(T.A, T.B);
	double b = Leng(T.B, T.C);
	double c = Leng(T.A, T.C);

	return a+b+c;
}

double Leng(TPoint A, TPoint B)
{
	return sqrt( pow(A.X - B.X, 2) + pow(A.Y - B.Y, 2) );
}