import { DataNode } from "./models";

interface IStack<T> {
  push(data: T): void;
  pop(): T | null;
  peek(): T | null;
  size(): number;
}

export class Stack<T> implements IStack<T> {
  private _size: number;
  private pointer: DataNode<T> | null;

  constructor() {
    this._size = 0;
    this.pointer = null;
  }

  push(data: T) {
    const node: DataNode<T> = {
      data,
      next: null
    };
    if (this.pointer === null) {
      this.pointer = node;
    } else {
      node.next = this.pointer;
      this.pointer = node;
    }
    this._size += 1;
  }

  pop(): T | null {
    if (this._size === 0) {
      return null;
    }
    const item = this.pointer.data;
    this.pointer = this.pointer.next;
    this._size -= 1;
    return item;
  }

  peek(): T | null {
    if (this._size === 0) {
      return null;
    }
    return this.pointer.data;
  }

  size(): number {
    return this._size;
  }
}
