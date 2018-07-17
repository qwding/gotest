#include <iostream>
#include <cstring>
using namespace std;

string start(int num[], int k,int lenNum)
{
    int result[lenNum-k];
    int maxLen = lenNum - k;
    int useK = 0;
    int resLen = 0;
    //int useLen = 0
    //while(useK < k || resLen <maxLen){

    //}
    for (int i = 0; i < lenNum; i++)
    {
        if (resLen >= maxLen)
        {
            break;
        }
        //  else if (useLen )
        if (resLen == 0)
        {
            result[0] = num[i];
            resLen++;
            // useLen++;
        }
        else if (result[resLen -1] >= num[i])
        {
            result[resLen] = num[i];
            resLen++;
            //      useLen++;
        }
        else
        {
            while(resLen - 1 >= 0)
            {
                if (result[resLen - 1] < num[i])
                {
                    resLen--;
                    useK++;
                    if (useK >= k)
                    {
                        for(; i<lenNum; i++,resLen++)
                        {
                            result[resLen] = num[i];
                        }
                        break;
                    }
                }
                result[resLen] = num[i];
                resLen++;
                //useLen++;
            }
        }
        for(int t; t<maxLen; t++)
        {
            cout<<result[t];
        }
        cout<<endl;
    }




    string resultInt = "";
    for (int i =0; i < lenNum - k; i ++)
    {
        char tmp  = '0' + result[i];
        resultInt = resultInt + tmp;
        // cout << "tmp is "<< resultInt<< endl;
    }
    // cout << "result num is  "<< resultInt << endl;
    return resultInt;
}
int main()
{
    int l;
    cin >> l;
    string resultArray[l];
    for (int t=0; t<l; t++)
    {
        string numStr;
        cin >> numStr;
        int k;
        cin >> k;
        int lenNum = numStr.length();
        int num[lenNum];
        for (int i = 0 ; i < numStr.length(); i++)
        {
            num[i] = numStr[i] - '0';
        }

        resultArray[t] = start(num,k,lenNum);

    }
    for(int t=0; t<l; t++)
    {
        cout <<resultArray[t]<<endl;
    }
    return 0;
}
