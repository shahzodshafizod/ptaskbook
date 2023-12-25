#include <iostream>
using namespace std;

int main()
{
	//Task("Begin24");
	float A, B, C;
	cin >> A >> B >> C;
	float D = A;
	A = B;
	B = C;
	C = D;
	cout << A << B << C;
	
	return 0;
}
