#include <iostream>
using namespace std;

int main()
{
	//Task("If23");
	int x1, y1, x2, y2, x3, y3, x4, y4;
	cin >> x1 >> y1 >> x2 >> y2 >> x3 >> y3;
	x4 = x1 == x2 ? x3 : x1 == x3 ? x2 : x1;
	y4 = y1 == y2 ? y3 : y1 == y3 ? y2 : y1;
	cout << x4 << y4;
	
	return 0;
}
