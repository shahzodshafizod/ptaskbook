#include <iostream>
using namespace std;

int main()
{
	//Task("While18");
	int N;
	cin >> N;
	int counts = 0, sum = 0;
	while (N > 0)
	{
		sum += N % 10;
		counts++;
		N /= 10;
	}
	cout << counts << sum;
	
	return 0;
}
