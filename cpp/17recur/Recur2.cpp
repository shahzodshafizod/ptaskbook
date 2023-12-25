#include <iostream>
using namespace std;

double Fact2(int N);

int main()
{
	//Task("Recur2");
	int n;
	for (int i = 1; i < 6; i++)
	{
		cin >> n;
		cout << Fact2(n);
	}
	return 0;
}

double Fact2(int N)
{
	if (N <= 1)
		return 1;
	return Fact2(N - 2) * N;
}