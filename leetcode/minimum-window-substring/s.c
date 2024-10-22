/**
 * Input: s = "ADOBECODEBANC", t = "ABC"
 * Output: "BANC"
 * Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.
 */

#include <stdio.h>
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

    int i, j;
    int need_n;
    int need_m[128];
    int k;
    int min_value = len_s + 1;
    char *res = malloc(2 * len_s * sizeof(char));

    for (i = 0; i < len_s; i++)
    {
        // init
        need_n = len_t;
        for (k = 0; k < 128; k++)
        {
            need_m[k] = 0;
        }
        for (k = 0; k < len_t; k++)
        {
            need_m[t[k]]++;
        }

        j = i;
        while (j < len_s && need_n > 0)
        {
            if (need_m[s[j]] > 0)
            {
                need_m[s[j]]--;
                need_n--;
            }
            j++;
        }
        j = j - 1;
        if (need_n == 0 && j - i < min_value)
        {
            min_value = j - i + 1;
            strncpy(res, &s[i], j - i + 1);
            res[j - i + 1] = '\0';
        }
    }

    if (min_value == len_s + 1)
    {
        return "";
    }

    return res;
}

int main()
{
    char *s = "ADOBECODEBANC";
    char *t = "ABC";

    char *new = minWindow(s, t);
    printf("res is: %s\n", new);

    s = "a";
    t = "aa";
    new = minWindow(s, t);
    printf("res is: %s\n", new);

    s = "a";
    t = "a";
    new = minWindow(s, t);
    printf("res is: %s\n", new);
}
