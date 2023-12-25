#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File28");
	char fileName[100];
	cin.getline(fileName, 100);
	fstream fs(fileName, ios::binary | ios::in | ios::out);
	double prev, curr, next, AMean;
	fs.read(reinterpret_cast<char*>(&prev), sizeof(prev));
	fs.read(reinterpret_cast<char*>(&curr), sizeof(curr));
	int index = 0;
	while (true)
	{
		fs.read(reinterpret_cast<char*>(&next), sizeof(next));
		if (fs.eof())
		{
			fs.clear();
			break;
		}
		index++;
		fs.seekp(index*sizeof(double));
		AMean = (prev + curr + next) / 3;
		fs.write(reinterpret_cast<const char*>(&AMean), sizeof(AMean));
		fs.seekg(sizeof(double), ios::cur);
		prev = curr;
		curr = next;
	}
	fs.close();
	
	return 0;
}
