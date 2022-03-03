import { BST } from './tree';
import { Stack } from './stack';
import { Queue } from './queue';
import { DoublyLinkedList, SinglyLinkedList, } from './linked_list';

const bst = new BST();
bst.create(20);
bst.create(11);
bst.create(1);
bst.create(22);
bst.create(12);
bst.create(7);
bst.create(0);
bst.create(25);
// console.log('BFS:', bst.bfs());
// console.log('PreOrder:', bst.preOrder());
// console.log('PostOrder:', bst.postOrder());
// console.log('InOrder:', bst.inOrder());

const stack = new Stack<number>();
// console.log(stack);
// stack.push(5);
// stack.push(1);
// stack.push(10);
// console.log(stack.peek());
// stack.pop();
// console.log(stack);

const queue = new Queue<number>();
// console.log(queue);
// queue.enqueue(0);
// queue.enqueue(1);
// console.log(queue);
// queue.enqueue(2);
// queue.enqueue(3);
// console.log(queue);

const dll = new DoublyLinkedList<string>();
dll.push('A');
dll.push('B');
dll.push('C');
dll.removeAt(1);
dll.insertAt(1, 'After A');
// console.log(dll.getAt(1));
// dll.pop();
// dll.dequeue();
dll.print();

const sll = new SinglyLinkedList<number>();
sll.push(1);
sll.enqueue(0);
sll.insertAt(2, 2);
sll.print();
// sll.removeAt(1);
// sll.print();
sll.reverse();

