#include <iostream>
#include <cmath>
using namespace std;

int main()
{
	//Task("Begin20");
	float x1, y1, x2, y2;
	cin >> x1 >> y1 >> x2 >> y2;
	float S = sqrt(pow(x2 - x1, 2) + pow(y2 - y1, 2));
	cout << S;
	
	return 0;
}
