export interface DataNode<T> {
  data: T;
  next: DataNode<T> | null;
  prev?: DataNode<T> | null;
} 
