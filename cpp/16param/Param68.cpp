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

void Altitudes(TTriangle T, double& h1, double& h2, double& h3);
double Dist(TPoint P, TPoint A, TPoint B);
double Area(TTriangle T);
double Leng(TPoint A, TPoint B);
double Perim(TTriangle T);

int main()
{
	//Task("Param68");
	TPoint A, B, C, D;

	cin >> A.X >> A.Y;
	cin >> B.X >> B.Y;
	cin >> C.X >> C.Y;
	cin >> D.X >> D.Y;

	double h1, h2, h3;
	TTriangle sekunja;
	
	sekunja.A = A; sekunja.B = B; sekunja.C = C;
	Altitudes(sekunja, h1, h2, h3);
	cout << h1 << h2 << h3;

	sekunja.C = D;
	Altitudes(sekunja, h1, h2, h3);
	cout << h1 << h2 << h3;

	sekunja.B = C;
	Altitudes(sekunja, h1, h2, h3);
	cout << h1 << h2 << h3;
	
	return 0;
}

void Altitudes(TTriangle T, double& h1, double& h2, double& h3)
{
	h1 = Dist(T.A, T.B, T.C);
	h2 = Dist(T.B, T.A, T.C);
	h3 = Dist(T.C, T.A, T.B);
}

double Dist(TPoint P, TPoint A, TPoint B)
{
	TTriangle tr;
	tr.A = A;
	tr.B = B;
	tr.C = P;

	return 2 * Area(tr) / Leng(A, B);
}

double Area(TTriangle T)
{
	double a = Leng(T.A, T.B);
	double b = Leng(T.B, T.C);
	double c = Leng(T.A, T.C);

	//double p = Perim(T) / 2;
	double p = (a+b+c) / 2;
	
	return sqrt( p * (p - a) * (p - b) * (p - c) );
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

	return a + b + c;
}