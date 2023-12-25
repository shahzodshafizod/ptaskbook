#include <iostream>
using namespace std;

int main()
{
	//Task("Series11");
	int K, N, number;
	cin >> K >> N;
	bool lessK = false;
	for (int i = 1; i <= N; i++)
	{
		cin >> number;
		if (number < K)
			lessK = true;
	}
	cout << lessK;
	
	return 0;
}
