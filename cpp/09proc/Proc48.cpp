#include <iostream>
using namespace std;

int GCD2(int, int);
int LCM2(int, int);

int main()
{
	//Task("Proc48");
	int A, B;
	cin >> A;
	for (int i = 1; i < 4; i++)
	{
		cin >> B;
		cout << LCM2(A, B);
	}
	
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

int LCM2(int A, int B)
{
	return A * ( B / GCD2(A, B) );
}