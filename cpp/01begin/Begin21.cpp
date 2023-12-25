#include <iostream>
#include <cmath>
using namespace std;

int main()
{
	//Task("Begin21");
	double x1, y1, x2, y2, x3, y3;
	cin >> x1 >> y1 >> x2 >> y2 >> x3 >> y3;
	double a = sqrt(pow(x2 - x1, 2) + pow(y2 - y1, 2));
	double b = sqrt(pow(x3 - x2, 2) + pow(y3 - y2, 2));
	double c = sqrt(pow(x3 - x1, 2) + pow(y3 - y1, 2));
	double P = a + b + c;
	double p = P / 2;
	double S = sqrt(p * (p - a) * (p - b) * (p - c));
	cout << P << S;
	
	return 0;
}
