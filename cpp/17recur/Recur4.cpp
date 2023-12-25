#include <iostream>
using namespace std;

int Fib1(int N);
int counts;

int main()
{
	//Task("Recur4");
	int N;
	for (int i = 0; i < 5; i++)
	{
		cin >> N;
		counts = 0;
		cout << Fib1(N);
		cout << counts;
	}
	return 0;
}

int Fib1(int N)
{
	counts++;
	if (N <= 2)
		return 1;
	return Fib1(N-1) + Fib1(N-2);
}