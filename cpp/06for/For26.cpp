#include <iostream>
using namespace std;

int main()
{
	//Task("For26");
	double X;
	cin >> X;
	int N;
	cin >> N;
	int factor = 1;
	double degreeX = 1, sum = 0;
	for (int i = 0; i <= N; i++)
	{
		degreeX *= X;
		sum += factor * degreeX / (2*i+1);
		degreeX *= X;
		factor *= -1;
	}
	cout << sum;
	
	return 0;
}
