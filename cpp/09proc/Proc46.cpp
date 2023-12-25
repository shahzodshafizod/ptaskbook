#include <iostream>
using namespace std;

int GCD2(int A, int B);

int main()
{
	//Task("Proc46");
	int A, B, C, D;
	cin >> A >> B >> C >> D;

	cout << GCD2(A, B);
	cout << GCD2(A, C);
	cout << GCD2(A, D);
	
	return 0;
}

int GCD2(int A, int B)
{
	int temp;
	while (B != 0)
	{
		temp = B;
		B = A % B;
		A = temp;
	}
	return A;
}