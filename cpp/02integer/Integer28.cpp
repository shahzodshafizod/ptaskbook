#include <iostream>
using namespace std;

int main()
{
	//Task("Integer28");
	int K, N;
	cin >> K >> N;
	int weekDay = (K + N - 2) % 7 + 1;
	cout << weekDay;
	
	return 0;
}
