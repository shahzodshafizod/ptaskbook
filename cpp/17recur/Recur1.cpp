#include <iostream>
using namespace std;

double Fact(int N);

int main()
{
	//Task("Recur1");
	int n;
	for (int i = 1; i < 6; i++)
	{
		cin >> n;
		cout << Fact(n);
	}
	return 0;
}

double Fact(int N)
{
	if (N <= 1)
		return 1;
	return Fact(N-1) * N;
}