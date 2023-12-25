#include <iostream>
using namespace std;

struct Rect
{
	double a;
	double b;
};

int main()
{
	//Task("Minmax3");
	int N;
	cin >> N;

	Rect rect;
	cin >> rect.a >> rect.b;

	double P, maxP;
	maxP = 2 * (rect.a + rect.b);

	for (int i = 2; i <= N; i++)
	{
		cin >> rect.a >> rect.b;
		P = 2 * (rect.a + rect.b);
		if (P > maxP)
			maxP = P;
	}

	cout << maxP;
	
	return 0;
}
