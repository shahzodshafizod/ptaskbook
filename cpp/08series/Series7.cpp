#include <iostream>
using namespace std;

void Solve()
{
	//Task("Series7");
	int N;
	pt >> N;
	double number;
	int rounded, factor, sum = 0;
	for (int i = 1; i <= N; i++)
	{
		pt >> number;
		factor = number < 0 ? -1 : 1;
		number *= factor;
		rounded = number - (int)number >= 0.5 ? (int)number + 1 : (int)number;
		rounded *= factor;
		pt << rounded;
		sum += rounded;
	}
	pt << sum;
	
	return 0;
}