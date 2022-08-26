namespace SudokuSolver;

public class Node
{
    private Node? _next;
    private Node? _previous;

    public readonly Cell Cell;

    public Node(Cell cell)
    {
        Cell = cell;
    }

    public void SetNext(Node nextNode) => _next = nextNode;
    
    public void SetPrevious(Node previousNode) => _previous = previousNode;

    public Node GetNext() => _next ?? throw new InvalidOperationException("Next node is undefined");
    public Node GetPrevious() => _previous ?? throw new InvalidOperationException("Previous node is undefined");
}