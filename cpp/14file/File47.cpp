#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File47");
	char fileName1[100], fileName2[100];
	cin.getline(fileName1, 100);
	cin.getline(fileName2, 100);
	fstream one(fileName1, ios::binary | ios::in | ios::out);
	fstream two(fileName2, ios::binary | ios::in | ios::out);
	one.seekg(0, ios::end);
	two.seekg(0, ios::end);
	int oneSize = one.tellg();
	int twoSize = two.tellg();
	one.seekg(0);
	two.seekg(0, ios::end);
	// from one to two
	char C;
	while (true)
	{
		one.read(reinterpret_cast<char*>(&C), sizeof(C));
		if (one.eof())
		{
			one.clear();
			break;
		}
		two.write(reinterpret_cast<const char*>(&C), sizeof(C));
	}
	// from two to one
	two.seekg(0);
	one.seekg(0, ios::end);
	for (int i = 0; i < twoSize; i++)
	{
		two.read(reinterpret_cast<char*>(&C), sizeof(C));
		if (two.eof())
		{
			two.clear();
			break;
		}
		one.write(reinterpret_cast<const char*>(&C), sizeof(C));
	}
	one.close();
	two.close();
	
	return 0;
}
