#include <iostream>
#include <cmath>
using namespace std;

double Power1(double A, double B);
double Power2(double A, int N);
double Power3(double A, double B);

int main()
{
	//Task("Proc39");
	double P, N;
	cin >> P;
	for (int i = 1; i < 4; i++)
	{
		cin >> N;
		cout << Power3(N, P);
	}
	
	return 0;
}

double Power3(double A, double B)
{
	int N = B;
	
	return ( N == B ? Power2(A, N) : Power1(A, B) );
}

double Power1(double A, double B)
{
	if (A <= 0)
		return 0;
	
	return exp(B * log(A));
}

double Power2(double A, int N)
{
	if (N == 0)
		return 1;
	
	bool isNegative = false;
	if (N < 0)
	{
		isNegative = true;
		N *= -1;
	}
	double degreeN = 1;
	for (int i = 1; i <= N; i++)
		degreeN *= A;

	if (isNegative)
		return 1 / degreeN;
	return degreeN;
}