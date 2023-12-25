#include <iostream>
using namespace std;

int main()
{
	//Task("If19");
	int a, b, c, d;
	cin >> a >> b >> c >> d;
	if ( (b == c) && (c == d) )
		cout << 1;
	else if ( (a == c) && (c == d) )
		cout << 2;
	else if ( (a == b) && (b == d) )
		cout << 3;
	else
		cout << 4;

	return 0;
}
