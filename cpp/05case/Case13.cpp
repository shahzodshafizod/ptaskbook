#include <iostream>
#include <cmath>
using namespace std;

int main()
{
	//Task("Case13");
	int index;
	cin >> index;
	double value, a, c, h, S;
	cin >> value;
	switch (index)
	{
		case 1:
			a = value;
			c = a * sqrt(2.0);
			h = c / 2;
			S = c * h / 2;
			break;
		case 2:
			c = value;
			a = c / sqrt(2.0);
			h = c / 2;
			S = c * h / 2;
			break;
		case 3:
			h = value;
			c = 2 * h;
			a = c / sqrt(2.0);
			S = c * h / 2;
			break;
		case 4:
			S = value;
			h = sqrt(S);
			c = 2 * h;
			a = c / sqrt(2.0);
			break;
	}

	if (index != 1)
		cout << a;

	if (index != 2)
		cout << c;

	if (index != 3)
		cout << h;

	if (index != 4)
		cout << S;
	
	return 0;
}
