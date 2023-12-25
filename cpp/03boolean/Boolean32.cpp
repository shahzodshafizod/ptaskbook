#include <iostream>
using namespace std;

int main()
{
	//Task("Boolean32");
	int a, b, c;
	cin >> a >> b >> c;
	bool recTriangle = (a*a == b*b + c*c) ||
						(b*b == a*a + c*c) ||
						(c*c == a*a + b*b);
	cout << recTriangle;
	
	return 0;
}
