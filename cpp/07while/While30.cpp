#include <iostream>
using namespace std;

int main()
{
	//Task("While30");
	double A, B, C;
	cin >> A >> B >> C;
	
	int C_in_A = 0;
	while (A >= C)
	{
		A -= C;
		C_in_A++;
	}

	int C_in_B = 0;
	while (B >= C)
	{
		B -= C;
		C_in_B++;
	}

	int kvads = 0;
	int i = 1;
	while (i <= C_in_A)
	{
		int j = 1;
		while (j <= C_in_B)
		{
			kvads++;
			j++;
		}
		i++;
	}

	cout << kvads;
	
	return 0;
}
