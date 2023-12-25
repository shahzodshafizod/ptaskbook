#include <iostream>
using namespace std;

int main()
{
	//Task("Boolean33");
	int a, b, c;
	cin >> a >> b >> c;
	bool triangle = (a+b > c) && (a+c > b) && (b+c > a);
	cout << triangle;
	
	return 0;
}
