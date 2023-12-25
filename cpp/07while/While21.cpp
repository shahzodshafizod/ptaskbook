#include <iostream>
using namespace std;

int main()
{
	//Task("While21");
	int N;
	cin >> N;
	bool hasUneven = false;
	while (N > 0)
	{
		if (N % 10 % 2 != 0)
			hasUneven = true;
		N /= 10;
	}
	cout << hasUneven;
	
	return 0;
}
