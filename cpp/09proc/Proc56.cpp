#include <iostream>
#include <cmath>
using namespace std;

struct Point
{
	double x;
	double y;
};

double Leng(double xA, double yA, double xB, double yB);

int main()
{
	//Task("Proc56");
	Point A, B;
	cin >> A.x >> A.y;
	for (int i = 1; i < 4; i++)
	{
		cin >> B.x >> B.y;
		cout << Leng(A.x, A.y, B.x, B.y);
	}
	
	return 0;
}

double Leng(double xA, double yA, double xB, double yB)
{
	return sqrt( pow(xB - xA, 2) + pow(yB - yA, 2) );
}