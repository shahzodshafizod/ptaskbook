#include <iostream>
using namespace std;

int main()
{
	//Task("While4");
	int N;
	cin >> N;
	int degree3 = 1;
	while (degree3 < N)
		degree3 *= 3;
	bool isDegree3 = degree3 == N;
	cout << isDegree3;
	
	return 0;
}
