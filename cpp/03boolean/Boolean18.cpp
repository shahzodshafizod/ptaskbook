#include <iostream>
using namespace std;

int main()
{
	//Task("Boolean18");
	int a, b, c;
	cin >> a >> b >> c;
	bool pair = (a == b) || (a == c) || (b == c);
	cout << pair;
	
	return 0;
}
