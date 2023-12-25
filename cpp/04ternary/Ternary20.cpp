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
	cout << (AB < AC ? B : C);
	cout << (AB < AC ? AB : AC);
	
	return 0;
}
