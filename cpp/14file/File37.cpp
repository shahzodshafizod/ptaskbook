#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File37");
	char fileName[100];
	cin.getline(fileName, 100);
	fstream fs(fileName, ios::binary | ios::in | ios::out);
	fs.seekg(0, ios::end);
	int number, elements = fs.tellg() / sizeof(int);
	for (int i = elements-1; i >= 0; i--)
	{
		fs.seekg(i*sizeof(int));
		fs.read(reinterpret_cast<char*>(&number), sizeof(number));
		fs.seekp(0, ios::end);
		fs.write(reinterpret_cast<const char*>(&number), sizeof(number));
	}
	fs.close();
	
	return 0;
}
