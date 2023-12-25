#include <iostream>
using namespace std;

int main()
{
	//Task("Boolean31");
	int a, b, c;
	cin >> a >> b >> c;
	bool equalTwoSides = (a == b) || (b == c) || (a == c);
	cout << equalTwoSides;
	
	return 0;
}
