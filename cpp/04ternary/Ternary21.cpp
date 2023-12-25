#include <iostream>
using namespace std;

int main()
{
	//Task("If21");
	int x, y;
	cin >> x >> y;
	cout << ( (x == 0) && (y == 0) ? 0 : y == 0 ? 1 : x == 0 ? 2 : 3 );
	
	return 0;
}
