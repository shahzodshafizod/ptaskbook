#include <iostream>
using namespace std;

int main()
{
	//Task("For24");
	double X;
	cin >> X;
	int N;
	cin >> N;
	double degreeX = 1, fact = 1, sum = 0;
	int factor = 1, facting = 0;
	for (int i = 0; i <= N; i++)
	{
		sum += factor * degreeX / fact;
		factor *= -1;
		degreeX *= X*X;
		fact *= (facting + 1) * (facting + 2);
		facting += 2;
	}
	cout << sum;
	
	return 0;
}
