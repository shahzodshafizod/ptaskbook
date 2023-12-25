#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File1");
	char S[100];
	cin.getline(S, 100);
	ofstream ofs(S);
	cout << ofs.is_open();
	
	return 0;
}
