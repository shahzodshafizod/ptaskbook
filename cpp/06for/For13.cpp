#include <iostream>
using namespace std;

int main()
{
	//Task("For13");
	int N;
	cin >> N;
	double elem = 1.1, sum = 0;
	int factor = 1;
	for (int i = 1; i <= N; i++)
	{
		sum += factor * elem;
		elem += 0.1;
		factor *= -1;
	}
	cout << sum;
	
	return 0;
}
