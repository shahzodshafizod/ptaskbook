#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File40");
	char fileName[100], tempFileName[] = "tempFileName.txt";
	cin.getline(fileName, 100);
	ifstream ifs(fileName, ios::binary);
	ofstream ofs(tempFileName, ios::binary);
	int number;
	for (int i = 1; true; i++)
	{
		ifs.read(reinterpret_cast<char*>(&number), sizeof(number));
		if (ifs.eof())
		{
			ifs.clear();
			break;
		}
		if (i % 2 != 0)
		{
			ofs.write(reinterpret_cast<const char*>(&number), sizeof(number));
			continue;
		}
		number = 0;
		for (int j = 0; j < 2; j++)
			ofs.write(reinterpret_cast<const char*>(&number), sizeof(number));
	}
	ifs.close();
	ofs.close();
	remove(fileName);
	rename(tempFileName, fileName);
	
	return 0;
}
