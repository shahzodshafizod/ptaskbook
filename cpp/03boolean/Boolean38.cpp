#include <iostream>
using namespace std;

int main()
{
	//Task("Boolean38");
	int x1, y1, x2, y2;
	cin >> x1 >> y1 >> x2 >> y2;
	bool fil = abs(x2 - x1) == abs(y2 - y1);
	cout << fil;
	
	return 0;
}
