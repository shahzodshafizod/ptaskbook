#include <iostream>
using namespace std;

int main()
{
	//Task("For11");
	int N;
	cin >> N;
	int sum = 0;
	for (int i = 0; i <= N; i++)
		sum += (N+i) * (N+i);
	cout << sum;
	
	return 0;
}
