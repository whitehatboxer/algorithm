#include <limits.h>
#include <stdlib.h>
#include <string.h>

char *minWindow(char *s, char *t)
{
    int len_s = strlen(s);
    int len_t = strlen(t);
    if (len_s == 0 || len_t == 0 || len_t > len_s)
    {
        return "";
    }

    int need[128] = {0};
    int have[128] = {0};
    int required = 0;

    // Initialize the frequency map for characters in `t`
    for (int i = 0; i < len_t; i++)
    {
        if (need[t[i]] == 0)
        {
            required++;
        }
        need[t[i]]++;
    }

    int left = 0, right = 0;
    int formed = 0;
    int min_len = INT_MAX;
    int start = 0;

    while (right < len_s)
    {
        char r_char = s[right];
        have[r_char]++;

        if (need[r_char] > 0 && have[r_char] == need[r_char])
        {
            formed++;
        }

        while (left <= right && formed == required)
        {
            char l_char = s[left];

            if (right - left + 1 < min_len)
            {
                min_len = right - left + 1;
                start = left;
            }

            have[l_char]--;
            if (need[l_char] > 0 && have[l_char] < need[l_char])
            {
                formed--;
            }

            left++;
        }

        right++;
    }

    if (min_len == INT_MAX)
    {
        return "";
    }

    char *res = (char *)malloc((min_len + 1) * sizeof(char));
    strncpy(res, &s[start], min_len);
    res[min_len] = '\0';

    return res;
}
