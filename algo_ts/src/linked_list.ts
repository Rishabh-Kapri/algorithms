import { DataNode } from "./models";

interface ILinkedList<T> {
  enqueue(data: T): void; // DLL: O(1), SLL: O(1)
  dequeue(): T | null; // DLL: O(1), SLL: O(1)
  push(data: T): void; // DLL: O(1), SLL: O(1)
  pop(): T | null; // DLL: O(1), SLL: O(n)
  getAt(index: number): DataNode<T> | null; // DLL: O(n), SLL: O(n)
  insertAt(index: number, data: T): void // DLL: O(n), SLL: O(n)
  removeAt(index: number): T | null; // DLL: O(n), SLL: O(n)
  size(): number;
  print(): void;
}

export class SinglyLinkedList<T> implements ILinkedList<T> {
  private head: DataNode<T> | null;
  private tail: DataNode<T> | null;
  private _size: number;

  constructor() {
    this._size = 0;
    this.head = this.tail = null;
  }

  enqueue(data: T): void {
    const node: DataNode<T> = {
      data,
      next: null
    };
    if (this.tail === null) {
      this.head = this.tail = node;
    } else {
      node.next = this.head;
      this.head = node;
    }
    this._size += 1;
  }

  dequeue(): T | null {
    let item: T | null;
    if (this._size === 0) {
      item = null;
    } else {
      item = this.head.data;
      if (this._size === 1) {
        this.head = this.tail = null;
      } else {
        this.head = this.head.next;
      }
      this._size -= 1;
    }
    return item;
  }

  push(data: T): void {
    const node: DataNode<T> = {
      data,
      next: null
    };
    if (this.tail === null) {
      this.head = this.tail = node;
    } else {
      this.tail.next = node;
      this.tail = node;
    }
    this._size += 1;
  }

  pop(): T | null {
    let item: T | null;
    if (this._size === 0) {
      item = null;
    } else {
      item = this.tail.data;
      if (this._size === 1) {
        this.head = this.tail = null;
      } else {
        // find second to last node
        let currentNode = this.head;
        let secondLastNode = this.head;
        while (currentNode.next) {
          secondLastNode = currentNode;
          currentNode = currentNode.next;
        }
        secondLastNode.next = null;
        this.tail = secondLastNode;
      }
      this._size -= 1;
    }
    return item;
  }

  getAt(index: number): DataNode<T> | null {
    let node: DataNode<T> | null;
    if (this._size === 0 || index < 0 || index >= this._size) {
      node = null;
    } else {
      let counter = 0;
      node = this.head;
      while (counter < index) {
        node = node.next;
        counter += 1;
      }
    }
    return node;
  }

  insertAt(index: number, data: T): void {
    if (index < 0 || index > this._size) {
      return;
    } else if (index === 0) {
      // enqueue
      this.enqueue(data);
    } else if (index === this._size) {
      // push
      this.push(data);
    } else {
      const prevNode = this.getAt(index - 1);
      const nextNode = prevNode.next;
      const node: DataNode<T> = {
        data,
        next: null
      };
      prevNode.next = node;
      node.next = nextNode;

      this._size += 1;
    }
  }

  removeAt(index: number): T {
    let item: T | null;
    if (index < 0 || index > this._size) {
      item = null;
    } else if (index === 0) {
      // dequeue
      item = this.dequeue();
    } else if (index === this._size - 1) {
      // pop
      item = this.pop();
    } else {
      const prevNode = this.getAt(index - 1);
      const nodeToRemove = prevNode.next;
      item = nodeToRemove.data;
      const nextNode = nodeToRemove.next;
      prevNode.next = nextNode;
      nodeToRemove.next = null;

      this._size -= 1;
    }
    return item;
  }

  size(): number {
    return this._size;
  }

  print(): void {
    let current = this.head;
    while (current) {
      console.log(`${current.data} | ${current.next?.data ?? 'null'}`);
      current = current.next;
    }
    console.log();
  }

  /**
   * In place reversal of linked list
   */
  reverse(): void {
    let currentNode = this.head;
    let prevNode = null;
    let nextNode = null;
    while (currentNode) {
      nextNode = currentNode.next;
      currentNode.next = prevNode;
      prevNode = currentNode;
      currentNode = nextNode;
    }
    this.tail = this.head;
    this.head = prevNode;
    console.log('Reversed:');
    this.print();
  }
}


export class DoublyLinkedList<T> implements ILinkedList<T> {
  private head: DataNode<T> | null;
  private tail: DataNode<T> | null;
  private _size: number;

  constructor() {
    this.head = this.tail = null;
    this._size = 0;
  }

  /**
   * Adds a node to the start of the linked list
   */
  enqueue(data: T): void {
    const node: DataNode<T> = {
      data,
      next: null,
      prev: null
    };

    if (this.tail === null) {
      this.head = this.tail = node;
    } else {
      this.head.prev = node;
      node.next = this.head;
      this.head = node;
    }
    this._size += 1;
  }

  /**
   * Removes a node from the start of the linked list
   */
  dequeue(): T | null {
    let item: T | null;
    if (this._size === 0) {
      item = null;
    } else {
      item = this.tail.data;
      if (this._size === 1) {
        this.head = null;
        this.tail = null;
      } else {
        this.head = this.head.next;
        this.head.prev = null;
      }
      this._size -= 1;
    }
    return item;
  }

  /**
   * Adds a node to the end of the linked list
   */
  push(data: T): void {
    const node: DataNode<T> = {
      data,
      next: null,
      prev: null
    };

    if (this.tail === null) {
      this.head = this.tail = node;
    } else {
      this.tail.next = node;
      node.prev = this.tail;
      this.tail = node;
    }
    this._size += 1;
  }

  /**
   * Removes a node from the end of the linked list
   */
  pop(): T | null {
    let item: T | null;
    if (this._size === 0) {
      item = null;
    } else {
      item = this.tail.data;
      if (this._size === 1) {
        this.head = this.tail = null;
      } else {
        this.tail = this.tail.prev;
        this.tail.next = null;
      }
      this._size -= 1;
    }
    console.log(item, this.tail, this.tail.prev);
    return item;
  }

  /**
   * Returns the data at the particular index
   */
  getAt(index: number): DataNode<T> | null {
    let node: DataNode<T> | null;
    if (this._size === 0 || index >= this._size || index < 0) {
      node = null;
    } else {
      let counter = 0;
      node = this.head;
      while (counter < index) {
        node = node.next;
        counter += 1;
      }
    }
    return node;
  }

  /**
   * Insert a node at a particular index
   */
  insertAt(index: number, data: T): void {
    if (index > this._size || index < 0) {
      return;
    } else if (index === 0) {
      // if index is 0 use enqueue
      this.enqueue(data);
    } else if (index === this._size) {
      // if index is equal to size use push
      this.push(data);
    } else {
      const node: DataNode<T> = {
        data,
        next: null,
        prev: null,
      };
      const prevNode = this.getAt(index - 1);
      const nextNode = prevNode.next;
      prevNode.next = node;
      node.prev = prevNode;
      node.next = nextNode;
      nextNode.prev = node;

      this._size += 1;
    }
  }

  /**
   * Removes a node from a particular index
   */
  removeAt(index: number): T | null {
    let item: T;
    if (index >= this._size || index < 0) {
      item = null;
    } else if (index === 0) {
      // removing at start, use dequeue
      item = this.dequeue();
    } else if (index === this._size - 1) {
      // removing at end, use pop
      item = this.pop();
    } else {
      const nodeToRemove = this.getAt(index);
      item = nodeToRemove.data;
      const prevNode = nodeToRemove.prev;
      const nextNode = nodeToRemove.next;
      prevNode.next = nextNode;
      nextNode.prev = prevNode;
      nodeToRemove.next = null;
      nodeToRemove.prev = null;

      this._size -= 1;
    }
    return item;
  }

  /**
   * Returns the size of the linked list
   */
  size(): number {
    return this._size;
  }

  print(): void {
    let current = this.head;
    while (current) {
      console.log(`${current.prev?.data ?? 'null'} | ${current.data} | ${current.next?.data ?? 'null'}`);
      current = current.next;
    }
    console.log();
  }
}
