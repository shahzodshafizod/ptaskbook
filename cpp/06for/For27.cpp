#include <iostream>
#include <cmath>
using namespace std;

int main()
{
	//Task("For27");
	double X;
	cin >> X;
	int N;
	cin >> N;

	double sum = X, degreeX = X;
	double evens = 1, unevens = 1;
	for (int i = 1; i <= N; i++)
	{
		degreeX *= X * X;
		evens *= (2 * i);
		unevens *= (2 * i - 1);
		sum += unevens * degreeX / (evens * (2 * i + 1));
	}
	cout << sum;
	
	return 0;
}
