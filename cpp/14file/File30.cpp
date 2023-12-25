#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File30");
	char fileName[100], tempFileName[] = "tempFileName.txt";
	cin.getline(fileName, 100);
	ifstream ifs(fileName, ios::binary);
	ofstream ofs(tempFileName, ios::binary);
	ifs.seekg(0, ios::end);
	int number, elements = ifs.tellg() / sizeof(int);
	ifs.seekg(0);
	for (int i = 0; i < elements/2; i++)
	{
		ifs.read(reinterpret_cast<char*>(&number), sizeof(number));
		ofs.write(reinterpret_cast<const char*>(&number), sizeof(number));
	}
	ifs.close();
	ofs.close();
	remove(fileName);
	rename(tempFileName, fileName);
	
	return 0;
}
