#include <iostream>
using namespace std;

int GCD2(int A, int B);
int GCD3(int A, int B, int C);

int main()
{
	//Task("Proc49");
	int A, B, C, D;
	cin >> A >> B >> C >> D;

	cout << GCD3(A, B, C);
	cout << GCD3(A, C, D);
	cout << GCD3(B, C, D);
	
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

int GCD3(int A, int B, int C)
{
	return GCD2( GCD2(A, B), C );
}