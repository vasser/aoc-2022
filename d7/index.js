const fs = require("fs");

const input = fs.readFileSync("./input.txt", "utf8").split("\n");
// const input = [
//   "$ cd /",
//   "$ ls",
//   "dir a",
//   "14848514 b.txt",
//   "8504156 c.dat",
//   "dir d",
//   "$ cd a",
//   "$ ls",
//   "dir e",
//   "29116 f",
//   "2557 g",
//   "62596 h.lst",
//   "$ cd e",
//   "$ ls",
//   "584 i",
//   "$ cd ..",
//   "$ cd ..",
//   "$ cd d",
//   "$ ls",
//   "4060174 j",
//   "8033020 d.log",
//   "5626152 d.ext",
//   "7214296 k",
// ];

const createDir = (name) => {
  return {
    type: "dir",
    name,
    children: [],
  };
};

const createFile = (name, size) => {
  return {
    type: "file",
    name,
    size: parseInt(size),
  };
};

function buildTree(input) {
  const paths = [];

  for (const line of input) {
    // blooody hell magic :shrug:
    const parent = paths.at(-1);

    // case 1: create directory
    if (line.startsWith("$ cd")) {
      const name = line.split(" ")[2];

      // go up
      if (name === "..") {
        paths.pop();
        continue;
      }

      // create dir
      const dir = createDir(name);
      if (parent) {
        parent.children.push(dir);
      }
      paths.push(dir);
    }

    // case 2: create file
    const [, size, filename] = line.match(/(\d+) (.+)/) || [];

    if (size && filename) {
      const file = createFile(filename, size);
      parent.children.push(file);
    }
  }

  return paths;
}

function getTreeSizes(tree) {
  const sizes = [];

  function traverse(node) {
    if (node.type === "file") {
      return node.size;
    }

    const dirSize = node.children.reduce((a, b) => a + traverse(b), 0);
    sizes.push(dirSize);
    return dirSize;
  }

  traverse(tree);

  return sizes;
}

function getDirToRemove(tree) {
  const systemSize = 70000000;
  const updateSize = 30000000;

  const sizes = getTreeSizes(tree[0]);

  sizes.sort((a, b) => b - a);

  const freeSpace = systemSize - sizes[0];

  let sizeToRemove = 0;

  for (let i = 0; i < sizes.length; i++) {
    if (freeSpace + sizes[i] < updateSize) {
      sizeToRemove = sizes[i - 1];
      break;
    }
  }

  return sizeToRemove;
}

const t = buildTree(input);

const task1 = getTreeSizes(t[0])
  .filter((s) => s <= 100000)
  .reduce((a, b) => a + b, 0);
console.log("Task 1", task1);

const task2 = getDirToRemove(t);
console.log("Task 1", task2);
