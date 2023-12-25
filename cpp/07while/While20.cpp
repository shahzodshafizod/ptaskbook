#include <iostream>
using namespace std;

int main()
{
	//Task("While20");
	int N;
	cin >> N;
	bool hasTwo = false;
	while (N > 0)
	{
		if (N % 10 == 2)
			hasTwo = true;
		N /= 10;
	}
	cout << hasTwo;
	
	return 0;
}
