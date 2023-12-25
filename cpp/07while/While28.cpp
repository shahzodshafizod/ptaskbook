#include <iostream>
#include <cmath>
using namespace std;

int main()
{
	//Task("While28");
	double eps;
	cin >> eps;
	double curr, prev = 2;
	curr = 2 + 1 / prev;
	int K = 2;
	while (abs(curr - prev) >= eps)
	{
		K++;
		prev = curr;
		curr = 2 + 1 / prev;
	}
	cout << K;
	cout << prev << curr;
	
	return 0;
}
