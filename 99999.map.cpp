#include <iostream>  
#include <cstdio>  
#include <cstring>  
#include <string>  
#include <map>  
using namespace std;  
  
int main()  
{  
    map<string,int> Map;  
    map<string,int> ::iterator it;  
    Map.insert(pair<string,int>("root",12));  
    Map.insert(pair<string,int>("scot",11));  
    for(it=Map.begin();it!=Map.end();it++)  
        cout<<it->first<<"    "<<it->second<<endl;  
    it=Map.begin();  
    Map.erase(it);//通过迭代器删除  
    string key="root";  
    Map.erase(key);//通过key删除  
      
    Map.erase(Map.begin(),Map.end());//一个迭代器，到另一个迭代器  
    //相当于  Map.clear();  
  
    for(it=Map.begin();it!=Map.end();it++)  
        cout<<it->first<<"    "<<it->second<<endl;  
    return 0;  
}