import { DataNode } from "./models";

interface IQueue<T> {
  enqueue(data: T): void;
  dequeue(): T | null;
  peek(): T | null;
  size(): number;
}

export class Queue<T> implements IQueue<T> {
  private _size: number;
  private head: DataNode<T> | null;
  private tail: DataNode<T> | null;

  constructor() {
    this._size = 0;
    this.head = this.tail = null;
  }

  size(): number {
    return this._size;
  }

  peek(): T | null {
    if (this._size === 0) {
      return null;
    }
    return this.head.data;
  }

  enqueue(data: T) {
    const node: DataNode<T> = {
      data,
      next: null
    }

    if (this.head === null) {
      this.head = this.tail = node;
    } else {
      this.tail.next = node;
      this.tail = this.tail.next;
      console.log();
    }

    this._size += 1;
  }

  dequeue(): T | null {
    if (this._size === 0) {
      return null;
    }
    const item = this.head.data;
    this.head = this.head.next;
    this._size -= 1;

    if (!this._size) {
      this.tail = null;
    }
    return item;
  }
}
