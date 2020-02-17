def longestPalindrome(s: str) -> str:
    sr = s[::-1]
    longest_str = ""
    length = len(s)
    for i in range(length):
        for j in range(i+1, length+1):
            subs = s[i:j]
            index = sr.find(subs)
            # print(subs, index, j - i)
            if index != -1 and len(longest_str) < (j - i):
                if j + index == length:
                    longest_str = subs
    return longest_str


a = "gtecdcelcg"
a1 = "aa"
a2 = "b"
a3 = "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggg"

print(longestPalindrome(a))
print(longestPalindrome(a1))
print(longestPalindrome(a2))
print(longestPalindrome(a3))





