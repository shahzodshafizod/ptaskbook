#include <iostream>
using namespace std;

int main()
{
	//Task("For8");
	int A, B;
	cin >> A >> B;
	int zarb = 1;
	for (int i = A; i <= B; i++)
		zarb *= i;
	cout << zarb;
	
	return 0;
}
