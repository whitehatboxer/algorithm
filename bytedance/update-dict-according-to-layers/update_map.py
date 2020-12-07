"""
给定一个 map，一个数组，描述关于 map 的层级，一段键值对
依据数组描述的层级，将键值对写入到 map 中
例如：
输入
{"A":"1","B":1,"C":{"D":[{"E":[{"F":"G"}]},{"E":[{"F":"G"}]}]}}
{"C","D","E"}
"X": "Y"

输出
{"A":"1","B":1,"C":{"D":[{"E":[{"F":"G","X":"Y"}]},{"E":[{"F":"G","X":"Y"}]}]}}
"""


def gen_new_map(_map, layers, value):

    if not _map:
        return

    if not layers:
        if isinstance(_map, list):
            for m in _map:
                for k, v in value.items():
                    m[k] = v
        elif isinstance(_map, dict):
            for k, v in value.items():
                _map[k] = v

        return

    layer = layers[0]
    new_layers = layers[1:]

    if isinstance(_map, list):
        for m in _map:
            gen_new_map(m, layers, value)
    elif isinstance(_map, dict):
        new_map = _map.get(layer)
        gen_new_map(new_map, new_layers, value)


a = {"A":"1","B":1,"C":{"D":[{"E":[{"F":"G"}]},{"E":[{"F":"G"}]}]}}
gen_new_map(
    a,
    ["C","D","E"],
    {"X": "Y"}
)
print(a)