#include <iostream>
#include <cmath>
using namespace std;

int main()
{
	//Task("While29");
	double eps;
	cin >> eps;
	double next, prev = 1, curr = 2;
	int K = 2;
	do
	{
		K++;
		next = (prev + 2 * curr) / 3;
		prev = curr;
		curr = next;
	}
	while (abs(curr - prev) >= eps);
	
	cout << K;
	cout << prev << curr;
	
	return 0;
}
