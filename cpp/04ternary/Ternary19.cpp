#include <iostream>
using namespace std;

int main()
{
	//Task("If19");
	int a, b, c, d;
	cin >> a >> b >> c >> d;
	cout << ( (b == c) && (c == d) ? 1 : (a == c) && (c == d) ? 2 : (a == b) && (b == d) ? 3 : 4 );
	
	return 0;
}
