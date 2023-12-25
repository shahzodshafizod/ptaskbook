#include <iostream>
using namespace std;

int main()
{
	//Task("If25");
	int x, f;
	cin >> x;
	f = (x < -2) || (x > 2) ? 2 * x : -3 * x;
	cout << f;
	
	return 0;
}
