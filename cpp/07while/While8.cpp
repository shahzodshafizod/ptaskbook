#include <iostream>
using namespace std;

int main()
{
	//Task("While8");
	int N;
	cin >> N;
	int i = 1;
	while (i*i <= N)
		i++;
	--i;
	cout << i;
	
	return 0;
}
