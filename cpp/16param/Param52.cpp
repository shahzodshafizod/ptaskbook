#include <iostream>
#include <fstream>
using namespace std;

void RemoveLineNumbers(const char* S);

int main()
{
	//Task("Param52");
	char fileName[100];
	cin.getline(fileName, 100);
	RemoveLineNumbers(fileName);
	
	return 0;
}

void RemoveLineNumbers(const char* S)
{
	ifstream file(S);
	if (file.is_open())
	{
		const int stringSize = 100;
		char str[stringSize+1];
		
		int index, len, K, L;
		do
		{
			file.getline(str, stringSize+1);
			len = strlen(str);
			if (len == 0)
				continue;
			
			index = K = L = 0;
			while ( (index < len) && (str[index] == ' ') )
				index++, K++;
			while ( (index < len) && (str[index] >= '0') && (str[index] <= '9') )
				index++, K++;
			while ( (index < len) && (str[index] == ' ') )
				index++, L++;
			
			break;
		}
		while (true);

		if (K > 0)
		{
			char tempFileName[] = "tempFileName.txt";
			ofstream tempFile(tempFileName);

			file.seekg(0);
			while (file.peek() != -1)
			{
				file.getline(str, stringSize+1);
				len = strlen(str);
				if (len == K)
				{
					tempFile << endl;
					continue;
				}
				index = K+L;
				while (str[index] != '\0')
					tempFile << str[index++];
				tempFile << endl;
			}
			file.close();
			tempFile.close();

			remove(S);
			rename(tempFileName, S);
		}
		else
			file.close();
	}
}