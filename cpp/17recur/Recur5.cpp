#include <iostream>
using namespace std;

int Fib2(int N);
int mass[20];

int main()
{
	//Task("Recur5");
	int N;
	for (int i = 1; i < 6; i++)
	{
		cin >> N;
		for (int j = 0; j < N; j++)
			mass[j] = 0;
		
		cout << Fib2(N);
	}
	return 0;
}

int Fib2(int N)
{
	if (N <= 2)
		return 1;
	
	mass[N-1] = ( mass[N-2] > 0 ? mass[N-2] : Fib2(N-1) ) + ( mass[N-3] > 0 ? mass[N-3] : Fib2(N-2) );
	return mass[N-1];
}