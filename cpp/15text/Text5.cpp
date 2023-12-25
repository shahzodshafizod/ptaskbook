#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("Text5");
	const int stringSize = 80;
	char fileName[100], S[stringSize];
	cin.getline(S, stringSize+1);
	cin.getline(fileName, 100);
	ofstream file(fileName, ios::app);
	file << S;
	file.close();
	
	return 0;
}
