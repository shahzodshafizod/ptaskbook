#include <iostream>
using namespace std;

int main()
{
	//Task("For28");
	double X;
	cin >> X;
	int N;
	cin >> N;
	double sum = 1, degreeX = 1;
	int unevens = 1, evens = 1, factor = 1;
	for (int i = 1; i <= N; )
	{
		degreeX *= X;
		evens *= 2*i;
		sum += factor * unevens * degreeX / evens;
		factor *= -1;
		i++;
		unevens *= 2*i - 3;
	}
	cout << sum;
	
	return 0;
}
