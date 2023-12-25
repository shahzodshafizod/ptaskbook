#include <iostream>
using namespace std;

int main()
{
	//Task("For17");
	double A;
	cin >> A;
	int N;
	cin >> N;
	double degreeA = 1, sum = 0;
	for (int i = 0; i <= N; i++)
	{
		sum += degreeA;
		degreeA *= A;
	}
	cout << sum;
	
	return 0;
}
