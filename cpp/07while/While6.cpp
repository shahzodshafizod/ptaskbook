#include <iostream>
using namespace std;

int main()
{
	//Task("While6");
	int N;
	cin >> N;
	double fact2 = 1;
	while (N > 0)
	{
		fact2 *= N;
		N -= 2;
	}
	cout << fact2;
	
	return 0;
}
