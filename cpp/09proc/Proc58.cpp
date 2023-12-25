#include <iostream>
#include <cmath>
using namespace std;

struct Point
{
	double x, y;
};

double Area(double xA, double yA, double xB, double yB, double xC, double yC);
double Perim(double xA, double yA, double xB, double yB, double xC, double yC,
			 double& a, double& b, double& c);
double Leng(double xA, double yA, double xB, double yB);

int main()
{
	//Task("Proc58");
	Point A, B, C, D;
	
	cin >> A.x >> A.y;
	cin >> B.x >> B.y;
	cin >> C.x >> C.y;
	cin >> D.x >> D.y;

	cout << Area(A.x, A.y, B.x, B.y, C.x, C.y);
	cout << Area(A.x, A.y, B.x, B.y, D.x, D.y);
	cout << Area(A.x, A.y, C.x, C.y, D.x, D.y);
	
	return 0;
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