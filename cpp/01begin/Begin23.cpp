#include <iostream>
using namespace std;

int main()
{
	//Task("Begin23");
	float A, B, C;
	cin >> A >> B >> C;
	float D = A;
	A = C;
	C = B;
	B = D;
	cout << A << B << C;
	
	return 0;
}
