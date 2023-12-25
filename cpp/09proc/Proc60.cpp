#include <iostream>
#include <cmath>
using namespace std;

struct Point
{
	double x, y;
};

void Altitudes(double xA, double yA, double xB, double yB, double xC, double yC,
			   double& hA, double& hB, double& hC);
double Dist(double xP, double yP, double xA, double yA, double xB, double yB);
double Area(double xA, double yA, double xB, double yB, double xC, double yC);
double Perim(double xA, double yA, double xB, double yB, double xC, double yC,
			 double& a, double& b, double& c);
double Leng(double xA, double yA, double xB, double yB);

int main()
{
	//Task("Proc60");
	Point A, B, C, D;

	cin >> A.x >> A.y;
	cin >> B.x >> B.y;
	cin >> C.x >> C.y;
	cin >> D.x >> D.y;

	double hA, hB, hC;
	Altitudes(A.x, A.y, B.x, B.y, C.x, C.y, hA, hB, hC);
	cout << hA << hB << hC;

	Altitudes(A.x, A.y, B.x, B.y, D.x, D.y, hA, hB, hC);
	cout << hA << hB << hC;

	Altitudes(A.x, A.y, C.x, C.y, D.x, D.y, hA, hB, hC);
	cout << hA << hB << hC;
	
	return 0;
}

void Altitudes(double xA, double yA, double xB, double yB, double xC, double yC,
			   double& hA, double& hB, double& hC)
{
	hA = Dist(xA, yA, xB, yB, xC, yC);
	hB = Dist(xB, yB, xA, yA, xC, yC);
	hC = Dist(xC, yC, xA, yA, xB, yB);
}

double Dist(double xP, double yP, double xA, double yA, double xB, double yB)
{
	return 2 * Area(xP, yP, xA, yA, xB, yB) / Leng(xA, yA, xB, yB);
}

double Area(double xA, double yA, double xB, double yB, double xC, double yC)
{
	double a, b, c;
	double p = Perim(xA, yA, xB, yB, xC, yC, a, b, c) / 2;

	return sqrt( p*(p-a)*(p-b)*(p-c) );
}

double Perim(double xA, double yA, double xB, double yB, double xC, double yC,
			 double& a, double& b, double& c)
{
	a = Leng(xA, yA, xB, yB);
	b = Leng(xA, yA, xC, yC);
	c = Leng(xB, yB, xC, yC);

	return a+b+c;
}

double Leng(double xA, double yA, double xB, double yB)
{
	return sqrt( pow(xB - xA, 2) + pow(yB - yA, 2) );
}