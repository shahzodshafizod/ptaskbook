#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File65");
	const int chars = 80;
	char fileInName[100], fileOutName[100], str[chars];
	cin.getline(fileInName, 100);
	cin.getline(fileOutName, 100);
	ifstream ifs(fileInName, ios::binary);
	ofstream ofs(fileOutName, ios::binary);
	int pos, length, maxLength = -1, elements = 0;
	int* poses = NULL;
	while (true)
	{
		pos = ifs.tellg();
		ifs.read(reinterpret_cast<char*>(&str), sizeof(str));
		if (ifs.eof())
		{
			ifs.clear();
			break;
		}
		length = strlen(str);
		
		if (maxLength < 0)
		{
			maxLength = length;
			elements = 1;
			poses = new int [elements];
			poses[elements-1] = pos;
		}
		else if (length > maxLength)
		{
			maxLength = length;
			delete [] poses;
			elements = 1;
			poses = new int [elements];
			poses[elements-1] = pos;
		}
		else if (length == maxLength)
		{
			int* temp = new int [elements+1];
			for (int i = 0; i < elements; i++)
				temp[i] = poses[i];
			temp[elements] = pos;
			elements++;
			delete [] poses;
			poses = temp;
			temp = NULL;
		}
	}
	for (int i = elements-1; i >= 0; i--)
	{
		ifs.seekg(poses[i]);
		ifs.read(reinterpret_cast<char*>(&str), sizeof(str));
		ofs.write(reinterpret_cast<const char*>(&str), sizeof(str));
	}
	ifs.close();
	ofs.close();
	
	return 0;
}
