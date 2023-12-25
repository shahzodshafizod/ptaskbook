#include <iostream>
using namespace std;

int main()
{
	//Task("For40");
	int A, B;
	cin >> A >> B;
	for (int i = A, k = 1; i <= B; i++, k++)
		for (int j = 1; j <= k; j++)
			cout << i;
	
	return 0;
}
