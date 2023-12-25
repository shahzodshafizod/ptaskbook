#include <iostream>
#include <cmath>
using namespace std;

int main()
{
	//Task("Begin19");
	float x1, y1, x2, y2;
	cin >> x1 >> y1 >> x2 >> y2;
	float a = abs(x2 - x1);
	float b = abs(y2 - y1);
	float P = 2 * (a + b);
	float S = a * b;
	cout << P << S;
	
	return 0;
}
