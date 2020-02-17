"""
简单的正则匹配

对于字符 s，使用模式 p 去匹配
其中 s 的范围为空字符串和小写字母 a-z
p 在 s 的基础上新增两种特殊字符 . 和 *
. 可以匹配 s 的任意一个字符
* 匹配前面出现的字符一次或者多次
"""

# 递归解法
# 只包含一个特殊字符 .
# 该场景下问题是比较简单的
def simpleMatch(s: str, p: str) -> bool:
    
    # 递归终止点
    if not p:
        return not s
    
    first_match = bool(s) and p[0] in {s[0], "."}
    
    # 递归规模减少
    return first_match and simpleMatch(s[1:], p[1:])


# 同样是递归解法
# 在 simpleMatch 的基础上扩展 * 字符，考虑这种字符的情况
# 其关键就是如何应用 * 的作用，同时也缩减递归的规模
def isMatch(s: str, p: str) -> bool:
    
    if not p:
        return not s
    
    first_match = bool(s) and p[0] in {s[0], "."}
    
    if len(p) >= 2 and p[1] == "*":
        return first_match and isMatch(s[1:], p) or isMatch(s, p[2:])
    else:
        return first_match and isMatch(s[1:], p[1:])
     

def test():
    print(simpleMatch("aaactgt", "...ct.t"))
    print(isMatch("aab", "c*a*b"))


test()










