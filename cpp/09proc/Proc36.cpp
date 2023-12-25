#include <iostream>
using namespace std;

int Fib(int N);

int main()
{
	//Task("Proc36");
	int N;
	for (int i = 1; i <= 5; i++)
	{
		cin >> N;
		cout << Fib(N);
	}
	
	return 0;
}

int Fib(int N)
{
	if (N <= 2)
		return 1;
	int prev = 1, curr = 1, next;
	for (int i = 3; i <= N; i++)
	{
		next = prev + curr;
		prev = curr;
		curr = next;
	}
	return curr;
}