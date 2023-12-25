#include <iostream>
using namespace std;

int main()
{
	//Task("Boolean37");
	int x1, y1, x2, y2;
	cin >> x1 >> y1 >> x2 >> y2;
	bool shoh = (abs(x2 - x1) < 2) && (abs(y2 - y1) < 2);
	cout << shoh;
	
	return 0;
}
