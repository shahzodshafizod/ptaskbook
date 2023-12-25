#include <iostream>
#include <cmath>
using namespace std;

int main()
{
	//Task("If20");
	double A, B, C;
	cin >> A >> B >> C;
	double AB = abs(B - A);
	double AC = abs(C - A);
	if (AB < AC)
		cout << B << AB;
	else
		cout << C << AC;
	
	return 0;
}
