#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File64");
	const int chars = 80;
	char fileInName[100], fileOutName[100], str[chars];
	cin.getline(fileInName, 100);
	cin.getline(fileOutName, 100);
	ifstream fileIn(fileInName, ios::binary);
	ofstream fileOut(fileOutName, ios::binary);
	int pos, length, minLength = -1, elements = 0;
	int* poses = NULL;
	while (true)
	{
		pos = fileIn.tellg();
		fileIn.read(reinterpret_cast<char*>(&str), sizeof(str));
		if (fileIn.eof())
		{
			fileIn.clear();
			break;
		}
		length = strlen(str);
		if (minLength < 0)
		{
			minLength = length;
			elements = 1;
			poses = new int [elements];
			poses[0] = pos;
		}
		else if (length == minLength)
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
		else if (length < minLength)
		{
			delete [] poses;
			elements = 1;
			poses = new int [elements];
			poses[0] = pos;
			
			minLength = length;
		}
	}
	fileIn.clear();
	for (int i = 0; i < elements; i++)
	{
		fileIn.seekg(poses[i]);
		fileIn.read(reinterpret_cast<char*>(&str), sizeof(str));
		fileOut.write(reinterpret_cast<const char*>(&str), sizeof(str));
	}

	fileIn.close();
	fileOut.close();
	
	return 0;
}