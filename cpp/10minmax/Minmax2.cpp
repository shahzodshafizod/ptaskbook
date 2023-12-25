#include <iostream>
using namespace std;

struct Rect
{
	double a;
	double b;
};

int main()
{
	//Task("Minmax2");
	int N;
	cin >> N;

	Rect rect;
	cin >> rect.a >> rect.b;
	
	double S, minS;
	minS = rect.a * rect.b;
	
	for (int i = 2; i <= N; i++)
	{
		cin >> rect.a >> rect.b;
		S = rect.a * rect.b;
		if (S < minS)
			minS = S;
	}

	cout << minS;
	
	return 0;
}
