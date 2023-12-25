#include <iostream>
using namespace std;

void SortKam(double&, double&, double&);

int main()
{
	//Task("Minmax23");
	int N;
	cin >> N;

	double number, maximal1, maximal2, maximal3;
	cin >> maximal1 >> maximal2 >> maximal3;

	SortKam(maximal1, maximal2, maximal3);

	for (int i = 4; i <= N; i++)
	{
		cin >> number;

		if (number > maximal3)
			maximal3 = number;

		SortKam(maximal1, maximal2, maximal3);
	}

	cout << maximal1 << maximal2 << maximal3;
	
	return 0;
}

void SortKam(double& a, double& b, double& c)
{
	double Small, Great;

	if ( (a > b) && (a > c) )
		Great = a;
	else if (b > c)
		Great = b;
	else
		Great = c;

	if ( (a < b) && (a < c) )
		Small = a;
	else if (b < c)
		Small = b;
	else
		Small = c;

	double sum = a + b + c;
	
	a = Great;
	b = sum - Great - Small;
	c = Small;
}