#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File18");
	char fileName[100];
	cin.getline(fileName, 100);
	ifstream ifs(fileName, ios::binary);
	double prev, curr, next, firstLocalMin;
	bool initialized = false;
	ifs.read(reinterpret_cast<char*>(&prev), sizeof(prev));
	ifs.read(reinterpret_cast<char*>(&curr), sizeof(curr));
	if (prev < curr)
	{
		firstLocalMin = prev;
		initialized = true;
	}
	else
		while (true)
		{
			ifs.read(reinterpret_cast<char*>(&next), sizeof(next));
			if (ifs.eof())
			{
				ifs.clear();
				break;
			}
			if ( (curr < prev) && (curr < next) )
			{
				firstLocalMin = curr;
				initialized = true;
				break;
			}
			prev = curr;
			curr = next;
		}
	if (!initialized)
		firstLocalMin = curr;

	ifs.close();
	cout << firstLocalMin;
	
	return 0;
}
