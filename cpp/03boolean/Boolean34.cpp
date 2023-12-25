#include <iostream>
using namespace std;

int main()
{
	//Task("Boolean34");
	int x, y;
	cin >> x >> y;
	bool white = (x+y) % 2 != 0;
	cout << white;
	
	return 0;
}
