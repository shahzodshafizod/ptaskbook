#include <iostream>
using namespace std;

int GCD2(int A, int B);
void Frac1(int a, int b, int& p, int& q);

int main()
{
	//Task("Proc47");
	int a, b, c, d, p, q;
	cin >> a >> b;
	for (int i = 1; i <= 3; i++)
	{
		cin >> c >> d;
		Frac1(d*a+b*c, b*d, p, q);
		cout << p << q;
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

void Frac1(int a, int b, int& p, int& q)
{
	int gcd = GCD2(a, b);
	p = a/gcd;
	q = b/gcd;
	if (q < 0)
	{
		p *= -1;
		q *= -1;
	}
}