#include <iostream>
using namespace std;

double Fact(int N);

int main()
{
	//Task("Proc34");
	int N;
	for (int i = 1; i < 6; i++)
	{
		cin >> N;
		cout << Fact(N);
	}
	
	return 0;
}

double Fact(int N)
{
	double result = 1;
	
	if (N <= 0)
		return result;

	do
	{
		result *= N;
		N--;
	}
	while (N > 0);
	
	return result;
}