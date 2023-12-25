#include <iostream>
using namespace std;

int main()
{
	//Task("Boolean14");
	int A, B, C;
	cin >> A >> B >> C;
	bool onePlus = (A > 0) && (B <= 0) && (C <= 0) ||
					(B > 0) && (A <= 0) && (C <= 0) ||
					(C > 0) && (A <= 0) && (B <= 0);
	cout << onePlus;
	
	return 0;
}
