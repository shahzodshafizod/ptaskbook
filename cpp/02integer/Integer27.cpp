#include <iostream>
using namespace std;

int main()
{
	//Task("Integer27");
	int K;
	cin >> K;
	int weekDay = (K + 4) % 7 + 1;
	cout << weekDay;
	
	return 0;
}
