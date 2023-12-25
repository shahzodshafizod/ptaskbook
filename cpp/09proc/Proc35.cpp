#include <iostream>
using namespace std;

double Fact2(int N);

int main()
{
	//Task("Proc35");
	int N;
	for (int i = 1; i <= 5; i++)
	{
		cin >> N;
		cout << Fact2(N);
	}
	
	return 0;
}

double Fact2(int N)
{
	double result = 1;
	if (N <= 0)
		return 1;
	do
	{
		result *= N;
		N -= 2;
	}
	while (N > 0);

	return result;
}