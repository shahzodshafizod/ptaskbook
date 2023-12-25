#include <iostream>
using namespace std;

int main()
{
	//Task("Begin28");
	double A;
	cin >> A;
	
	double deg1, deg2;
	
	deg1 = A * A;// deg1 = 2
	cout << deg1;
	
	deg2 = deg1 * A;// deg2 = 3
	cout << deg2;
	
	deg1 *= deg2;// deg1 = 5
	cout << deg1;
	
	deg2 = deg1 * deg1;// deg2 = 10
	cout << deg2;
	
	deg1 *= deg2;// deg1 = 15
	cout << deg1;
	
	return 0;
}
