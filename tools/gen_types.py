tag_name = "barrier"

class Item:
    def __init__(self, t: str, k: str, d: str):
        self.tag_name = t
        self.key = k
        self.description = d

items: list[Item] = []

# Read the raw data gathered from the documentation
with open("data.raw") as f:
    lines = f.read().splitlines()
    for line in lines:
        # Parse the various fields and make sure there are 3 elements
        line = line.replace("\t\t", "\t").split("\t")
        if len(line) != 3:
            continue
        items.append(Item(line[0].strip(), line[1].strip(), line[2].strip()))

for item in items:
    _const = f"{tag_name.capitalize()}{item.key.replace("_", " ").title().replace(" ", "")}"
    _type = f"{tag_name.capitalize()}Type"
    _description = f"{item.description[0].lower()}{item.description[1:]}"
    print(f"\t// {_const} represents {_description}\n\t{_const} {_type} = \"{item.key}\"")

print("-"*20)

_var = f"valid{tag_name.capitalize()}Types"
_map = f"map[string]{tag_name.capitalize()}Type"
print(f"var {_var} = {_map}{{")
for item in items:
    _const = f"{tag_name.capitalize()}{item.key.replace("_", " ").title().replace(" ", "")}"
    print(f"\t\"{item.key}\": {_const},")
print("}")
