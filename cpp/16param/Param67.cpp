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

double Dist(TPoint P, TPoint A, TPoint B);
double Area(TTriangle T);
double Perim(TPoint A, TPoint B, TPoint C);
double Leng(TPoint A, TPoint B);

int main()
{
	//Task("Param67");
	TPoint P, A, B, C;

	cin >> P.X >> P.Y;
	cin >> A.X >> A.Y;
	cin >> B.X >> B.Y;
	cin >> C.X >> C.Y;

	cout << Dist(P, A, B);
	cout << Dist(P, A, C);
	cout << Dist(P, B, C);
	
	return 0;
}

double Dist(TPoint P, TPoint A, TPoint B)
{
	TTriangle tr;
	tr.A = A;
	tr.B = B;
	tr.C = P;
	
	return 2*Area(tr) / Leng(A, B);
}

double Area(TTriangle T)
{
	double a = Leng(T.A, T.B);
	double b = Leng(T.B, T.C);
	double c = Leng(T.A, T.C);

	//double p = Perim(T) / 2;
	double p = (a+b+c) / 2;
	
	return sqrt( p*(p-a)*(p-b)*(p-c) );
}

double Perim(TPoint A, TPoint B, TPoint C)
{
	double a = Leng(A, B);
	double b = Leng(B, C);
	double c = Leng(A, C);

	return a + b + c;
}

double Leng(TPoint A, TPoint B)
{
	return sqrt( pow(A.X - B.X, 2) + pow(A.Y - B.Y, 2) );
}