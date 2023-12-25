#include <iostream>
using namespace std;

int main()
{
	//Task("Boolean15");
	int A, B, C;
	cin >> A >> B >> C;
	bool twoPluses = (A > 0) && (B > 0) && (C <= 0) ||
					(A > 0) && (C > 0) && (B <= 0) ||
					(B > 0) && (C > 0) && (A <= 0);
	cout << twoPluses;
	
	return 0;
}
