#include <iostream>
using namespace std;

int main()
{
	//Task("If27");
	double x;
	int f;
	cin >> x;
	if (x < 0)
		f = 0;
	else if ( (int)x % 2 == 0 )
		f = 1;
	else
		f = -1;
	cout << f;
	
	return 0;
}
