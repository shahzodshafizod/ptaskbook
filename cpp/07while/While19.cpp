#include <iostream>
using namespace std;

int main()
{
	//Task("While19");
	int N;
	cin >> N;
	int result = 0;
	while (N > 0)
	{
		result = result * 10 + N % 10;
		N /= 10;
	}
	cout << result;
	
	return 0;
}
