#include <iostream>
using namespace std;

int main()
{
	//Task("While25");
	int N;
	cin >> N;
	int prev, curr, next;
	prev = curr = 1;
	do
	{
		next = prev + curr;
		prev = curr;
		curr = next;
	}
	while (curr <= N);
	
	cout << curr;
	
	return 0;
}
