#include <iostream>
#include <cmath>
using namespace std;

struct Point
{
	double x, y;
};

double Leng(double xA, double yA, double xB, double yB);
double Area(double xA, double yA, double xB, double yB, double xC, double yC);
double Perim(double xA, double yA, double xB, double yB, double xC, double yC,
			 double& a, double& b, double& c);
double Dist(double xP, double yP, double xA, double yA, double xB, double yB);

int main()
{
	//Task("Proc59");
	Point P, A, B, C;

	cin >> P.x >> P.y;
	cin >> A.x >> A.y;
	cin >> B.x >> B.y;
	cin >> C.x >> C.y;

	cout << Dist(P.x, P.y, A.x, A.y, B.x, B.y);
	cout << Dist(P.x, P.y, A.x, A.y, C.x, C.y);
	cout << Dist(P.x, P.y, B.x, B.y, C.x, C.y);
	
	return 0;
}

double Leng(double xA, double yA, double xB, double yB)
{
	return sqrt( pow(xB - xA, 2) + pow(yB - yA, 2) );
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

double Dist(double xP, double yP, double xA, double yA, double xB, double yB)
{
	return 2 * Area(xP, yP, xA, yA, xB, yB) / Leng(xA, yA, xB, yB);
}