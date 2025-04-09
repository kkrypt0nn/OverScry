class Item:
    def __init__(self, k: str, d: str):
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
        d = line[2].strip()
        if not d.endswith("."):
            d = f"{d}."
        items.append(Item(line[1].strip(), d))


print("\n".join([f"* `{item.key}` - {item.description}" for item in items]))
