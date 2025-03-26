import os

files = sorted(os.listdir())

for file in files:
    if file.endswith(".go") and file != "00_index.go":
        print("// " + file + ":")
        with open(file, "r") as fin:
            for line in fin:
                line = line.strip()
                if line == "":
                    break
                if line.startswith("// - ") or line.startswith("//  "):
                    print(line)
        print("//")
