#include <iostream>
#include <cmath>
using namespace std;

struct Point
{
	double x, y;
};

double Leng(double xA, double yA, double xB, double yB);
double Perim(double xA, double yA, double xB, double yB, double xC, double yC);

int main()
{
	//Task("Proc57");
	Point A, B, C, D;
	
	cin >> A.x >> A.y;
	cin >> B.x >> B.y;
	cin >> C.x >> C.y;
	cin >> D.x >> D.y;

	cout << Perim(A.x, A.y, B.x, B.y, C.x, C.y);
	cout << Perim(A.x, A.y, B.x, B.y, D.x, D.y);
	cout << Perim(A.x, A.y, C.x, C.y, D.x, D.y);
	
	return 0;
}

double Leng(double xA, double yA, double xB, double yB)
{
	return sqrt( pow(xB - xA, 2) + pow(yB - yA, 2) );
}

double Perim(double xA, double yA, double xB, double yB, double xC, double yC)
{
	double a = Leng(xA, yA, xB, yB);
	double b = Leng(xA, yA, xC, yC);
	double c = Leng(xB, yB, xC, yC);

	return a + b + c;
}