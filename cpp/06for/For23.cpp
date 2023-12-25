#include <iostream>
using namespace std;

int main()
{
	//Task("For23");
	double X;
	cin >> X;
	int N;
	cin >> N;
	double sum = 0, fact = 1, degreeX = 1;
	int factor = 1, facting = 1;
	for (int i = 0; i <= N; i++)
	{
		degreeX *= X;
		sum += factor * degreeX / fact;
		degreeX *= X;
		factor *= -1;
		fact *= (facting + 1) * (facting + 2);
		facting += 2;
	}
	cout << sum;
	
	return 0;
}
