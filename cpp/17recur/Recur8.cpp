#include <iostream>
using namespace std;

double RootK(double X, int K, int N);
double PowerN(double X, int N);

int main()
{
	//Task("Recur8");
	double X;
	cin >> X;
	int K;
	cin >> K;
	int N;
	for (int i = 1; i < 7; i++)
	{
		cin >> N;
		cout << RootK(X, K, N);
	}
	return 0;
}

double RootK(double X, int K, int N)
{
	if (N == 0)
		return 1;

	double root = RootK(X, K, N-1);
		return root - ( root - X / PowerN(root, K-1) ) / K;
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