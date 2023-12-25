#include <iostream>
using namespace std;

int main()
{
	//Task("For14");
	int N;
	cin >> N;
	int square = 0;
	for (int i = 1; i <= N; i++)
	{
		square += 2 * i - 1;
		cout << square;
	}
	
	return 0;
}
