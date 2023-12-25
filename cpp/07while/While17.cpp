#include <iostream>
using namespace std;

int main()
{
	//Task("While17");
	int N;
	cin >> N;
	while (N > 0)
	{
		cout << N % 10;
		N /= 10;
	}
	
	return 0;
}
