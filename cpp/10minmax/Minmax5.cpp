#include <iostream>
using namespace std;

int main()
{
	//Task("Minmax5");
	int N;
	cin >> N;
	
	double m, v, P, maxP;
	cin >> m >> v;
	maxP = m / v;
	int index = 1;

	for (int i = 2; i <= N; i++)
	{
		cin >> m >> v;
		P = m / v;
		if (P > maxP)
		{
			maxP = P;
			index = i;
		}
	}
	cout << index << maxP;
	
	return 0;
}
