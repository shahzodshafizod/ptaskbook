#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File35");
	char fileName[100], tempFileName[] = "tempFileName.txt";
	cin.getline(fileName, 100);
	ifstream ifs(fileName, ios::binary);
	ofstream ofs(tempFileName, ios::binary);
	ifs.seekg(0, ios::end);
	int elements = ifs.tellg() / sizeof(int);
	int number = 0;
	for (int i = 50-elements; i > 0; i--)
		ofs.write(reinterpret_cast<const char*>(&number), sizeof(number));
	ifs.seekg(0);
	while (true)
	{
		ifs.read(reinterpret_cast<char*>(&number), sizeof(number));
		if (ifs.eof())
		{
			ifs.clear();
			break;
		}
		ofs.write(reinterpret_cast<const char*>(&number), sizeof(number));
	}
	ifs.close();
	ofs.close();
	remove(fileName);
	rename(tempFileName, fileName);
	
	return 0;
}
