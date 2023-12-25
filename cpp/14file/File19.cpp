#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File19");
	char fileName[100];
	cin.getline(fileName, 100);
	ifstream ifs(fileName, ios::binary);
	double prev, curr, next, lastLocalMax;
	ifs.read(reinterpret_cast<char*>(&prev), sizeof(prev));
	ifs.read(reinterpret_cast<char*>(&curr), sizeof(curr));
	if (prev > curr)
		lastLocalMax = prev;
	while (true)
	{
		ifs.read(reinterpret_cast<char*>(&next), sizeof(next));
		if (ifs.eof())
		{
			ifs.clear();
			break;
		}
		if ( (curr > prev) && (curr > next) )
			lastLocalMax = curr;
		prev = curr;
		curr = next;
	}
	if (curr > prev)
		lastLocalMax = curr;
	
	cout << lastLocalMax;
	
	return 0;
}
