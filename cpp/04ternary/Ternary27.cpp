#include <iostream>
using namespace std;

int main()
{
	//Task("If27");
	double x;
	int f;
	cin >> x;
	f = x < 0 ? 0 : (int)x % 2 == 0 ? 1 : -1;
	cout << f;
	
	return 0;
}
