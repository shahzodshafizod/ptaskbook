#include <iostream>
using namespace std;

int main()
{
	//Task("For39");
	int A, B;
	cin >> A >> B;
	for (int i = A; i <= B; i++)
		for (int j = 1; j <= i; j++)
			cout << i;
	
	return 0;
}
