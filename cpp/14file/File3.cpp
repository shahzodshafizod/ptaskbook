#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File3");
	char fileName[100];
	cin.getline(fileName, 100);
	double A, D;
	cin >> A >> D;
	ofstream file(fileName, ios::binary);
	double progress = A;
	for (int i = 0; i < 10; i++)
	{
		file.write( reinterpret_cast<const char*>(&progress), sizeof(progress) );
		progress += D;
	}
	file.close();
	
	return 0;
}
