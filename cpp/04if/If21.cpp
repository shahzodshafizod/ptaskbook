#include <iostream>
using namespace std;

int main()
{
	//Task("If21");
	int x, y;
	cin >> x >> y;
	if ( (x == 0) && (y == 0) )
		cout << 0;
	else if (y == 0)
		cout << 1;
	else if (x == 0)
		cout << 2;
	else
		cout << 3;

	return 0;
}
