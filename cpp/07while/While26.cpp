#include <iostream>
using namespace std;

int main()
{
	//Task("While26");
	int N;
	cin >> N;
	int prev, curr, next;
	prev = curr = 1;
	next = prev + curr;
	while (curr < N)
	{
		prev = curr;
		curr = next;
		next = prev + curr;
	}
	cout << prev << next;
	
	return 0;
}
