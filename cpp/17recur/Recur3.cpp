#include <iostream>
using namespace std;

double PowerN(double X, int N);

int main()
{
	//Task("Recur3");
	double X;
	cin >> X;
	int N;
	for (int i = 1; i < 6; i++)
	{
		cin >> N;
		cout << PowerN(X, N);
	}
	return 0;
}

double PowerN(double X, int N)
{
	if (N == 0)
		return 1;

	if (N < 0)
		return 1 / PowerN(X, -N);

	else if (N % 2 == 0)
	{
		double half = PowerN(X, N/2);
		return half * half;
	}

	else
		return X * PowerN(X, N-1);
}