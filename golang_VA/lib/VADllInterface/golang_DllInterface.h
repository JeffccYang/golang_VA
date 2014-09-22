
#define MATHFUNCSDLL_EXPORTS

#ifdef MATHFUNCSDLL_EXPORTS  
#define MATHFUNCSDLL_API __declspec(dllexport) 
#else 
#define MATHFUNCSDLL_API __declspec(dllimport) 
#endif

 
#ifdef __cplusplus
extern "C"
{
#endif
 
		MATHFUNCSDLL_API  int VA_fd(char* imageName, int w, int h); 
 
#ifdef __cplusplus
}
#endif