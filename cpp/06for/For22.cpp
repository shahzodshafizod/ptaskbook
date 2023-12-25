#include <iostream>
using namespace std;

int main()
{
	//Task("For22");
	double X;
	cin >> X;
	int N;
	cin >> N;
	double fact = 1, degreeX = 1, sum = 0;
	for (int i = 0; i <= N; i++)
	{
		sum += degreeX / fact;
		degreeX *= X;
		fact *= i+1;
	}
	cout << sum;
	
	return 0;
}
