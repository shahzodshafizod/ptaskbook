#include <iostream>
using namespace std;

int main()
{
	//Task("For18");
	double A;
	cin >> A;
	int N;
	cin >> N;
	int factor = 1;
	double degreeA = 1, sum = 0;
	for (int i = 0; i <= N; i++)
	{
		sum += factor * degreeA;
		degreeA *= A;
		factor *= -1;
	}
	cout << sum;
	
	return 0;
}
