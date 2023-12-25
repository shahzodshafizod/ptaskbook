#include <iostream>
using namespace std;

int main()
{
	//Task("For25");
	double X;
	cin >> X;
	int N;
	cin >> N;
	double degreeX = 1, sum = 0;
	int factor = 1;
	for (int i = 1; i <= N; i++)
	{
		degreeX *= X;
		sum += factor * degreeX / i;
		factor *= -1;
	}
	cout << sum;
	
	return 0;
}
