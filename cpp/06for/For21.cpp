#include <iostream>
using namespace std;

int main()
{
	//Task("For21");
	int N;
	cin >> N;
	double fact = 1, sum = 0;
	for (int i = 0; i <= N; i++)
	{
		sum += 1.0 / fact;
		fact *= i+1;
	}
	cout << sum;
	
	return 0;
}
