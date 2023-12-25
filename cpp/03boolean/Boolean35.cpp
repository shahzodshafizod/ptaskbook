#include <iostream>
using namespace std;

int main()
{
	//Task("Boolean35");
	int x1, y1, x2, y2;
	cin >> x1 >> y1 >> x2 >> y2;
	bool sameColor = ((x1+y1) % 2) == ((x2+y2) % 2);
	cout << sameColor;
	
	return 0;
}
