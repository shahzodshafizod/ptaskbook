#include <iostream>
using namespace std;

int main()
{
	//Task("Boolean39");
	int x1, y1, x2, y2;
	cin >> x1 >> y1 >> x2 >> y2;
	bool farzin = (x1 == x2) || (y1 == y2) || (abs(x2 - x1) == abs(y2 - y1));
	cout << farzin;
	
	return 0;
}
