#include <iostream>
using namespace std;

int main()
{
	//Task("Boolean29");
	float x, y, x1, y1, x2, y2;
	cin >> x >> y >> x1 >> y1 >> x2 >> y2;
	bool pointIn = (x > x1) && (x < x2) && (y > y2) && (y < y1);
	cout << pointIn;
	
	return 0;
}
