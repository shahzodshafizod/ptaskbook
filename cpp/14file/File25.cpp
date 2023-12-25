#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File25");
	char fileName[100];
	cin.getline(fileName, 100);
	fstream fs(fileName, ios::binary | ios::in | ios::out);
	double number;
	int doubleSize = sizeof(double);
	fs.seekg(0, ios::end);
	int elements = fs.tellg() / doubleSize;
	for (int i = 0; i < elements; i++)
	{
		fs.seekg(i*doubleSize);
		fs.read(reinterpret_cast<char*>(&number), sizeof(number));
		fs.seekg(i*doubleSize);
		number *= number;
		fs.write(reinterpret_cast<const char*>(&number), sizeof(number));
	}
	fs.close();
	
	return 0;
}
