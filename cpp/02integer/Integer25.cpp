#include <iostream>
using namespace std;

int main()
{
	//Task("Integer25");
	int K;
	cin >> K;
	int weekDay = (K + 3) % 7;
	cout << weekDay;
	
	return 0;
}
