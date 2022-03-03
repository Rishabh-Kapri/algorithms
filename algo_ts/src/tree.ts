export class TreeNode {
  val: string;
  children: TreeNode[];

  constructor(value: string) {
    this.val = value;
    this.children = [];
  }
}

export class BSTNode {
  val: number;
  left: BSTNode;
  right: BSTNode;

  constructor(value: number) {
    this.val = value;
    this.right = null;
    this.left = null;
  }
}

export class BST {
  root: BSTNode;
  constructor() {
    this.root = null;
  }

  create(val: number) {
    const newNode = new BSTNode(val)

    if (!this.root) {
      this.root = newNode;
      return this;
    }

    let current = this.root;

    const addSide = (side: 'left' | 'right') => {
      if (!current[side]) {
        current[side] = newNode;
        return this;
      }
      current = current[side];
    };

    while (true) {
      if (val === current.val) {
        return this;
      } else if (val < current.val) {
        addSide('left');
      } else {
        addSide('right');
      }
    }
  }

  bfs() {
    const visited: number[] = [];
    const queue: BSTNode[] = [];
    let current = this.root;

    queue.push(current);
    while (queue.length) {
      current = queue.shift();
      visited.push(current.val);

      if (current.left) {
        queue.push(current.left);
      }
      if (current.right) {
        queue.push(current.right)
      }
    };
    return visited;
  }

  // dfs traversal
  preOrder() {
    // root, left, right
    const visited: number[] = [];
    const current = this.root;

    const traverse = (node: BSTNode) => {
      visited.push(node.val);
      if (node.left) {
        traverse(node.left);
      }
      if (node.right) {
        traverse(node.right);
      }
    }

    traverse(current);
    return visited;
  }

  postOrder() {
    // left, right, root
    const visited: number[] = [];
    const current = this.root;

    const traverse = (node: BSTNode) => {
      if (node.left) {
        traverse(node.left);
      }
      if (node.right) {
        traverse(node.right);
      }
      visited.push(node.val);
    }

    traverse(current);
    return visited;
  }

  inOrder() {
    // left, root, right
    const visited: number[] = [];
    const current = this.root;

    const traverse = (node: BSTNode) => {
      if (node.left) {
        traverse(node.left);
      }
      visited.push(node.val);
      if (node.right) {
        traverse(node.right);
      }
    }

    traverse(current);
    return visited;
  }
}


