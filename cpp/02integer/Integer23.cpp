#include <iostream>
using namespace std;

int main()
{
	//Task("Integer23");
	int N;
	cin >> N;
	int minutes = N % 3600 / 60;
	cout << minutes;
	
	return 0;
}
