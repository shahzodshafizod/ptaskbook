#include <iostream>
using namespace std;

int main()
{
	//Task("While7");
	int N;
	cin >> N;
	int i = 1;
	while (i*i <= N)
		i++;
	cout << i;
	
	return 0;
}
