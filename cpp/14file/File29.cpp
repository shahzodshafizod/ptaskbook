#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File29");
	char fileName[100], tempFileName[] = "tempFileName.txt";
	cin.getline(fileName, 100);
	ifstream ifs(fileName, ios::binary | ios::in);
	ofstream ofs(tempFileName, ios::binary | ios::out);
	int number;
	for (int i = 0; i < 50; i++)
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
