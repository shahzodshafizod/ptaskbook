#include <iostream>
using namespace std;

double Power2(double A, int N);

int main()
{
	//Task("Proc38");
	double A;
	cin >> A;
	int N;
	for (int i = 1; i < 4; i++)
	{
		cin >> N;
		cout << Power2(A, N);
	}
	
	return 0;
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