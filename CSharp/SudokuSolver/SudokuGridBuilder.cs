namespace SudokuSolver;

public static class SudokuGridBuilder
{
    public static SudokuGrid Build(string input)
    {
        var nodeGrid = LinkNodes(input.Chunk(9).Select(ConvertToRow).ToArray());

        return new SudokuGrid(nodeGrid);
    }
    
    private static Node[] ConvertToRow(char[] input) =>
        input.Select(x => new Node(new Cell(ToInt(x)))).ToArray();

    private static int ToInt(char x) => Convert.ToInt32(x.ToString());

    private static Node[][] LinkNodes(Node[][] nodeGrid)
    {
        var firstNode = nodeGrid[0][0];
        var lastNode = nodeGrid[8][8];
        var previousNode = firstNode;
        
        for (var row = 0; row < 9; row++)
        {
            for (var col = 0; col < 9; col++)
            {
                var currentNode = nodeGrid[row][col];
                
                switch (row)
                {
                    case 0 when col == 0:
                        currentNode.SetPrevious(lastNode);
                        continue;
                    case 8 when col == 8:
                        currentNode.SetNext(firstNode);
                        currentNode.SetPrevious(previousNode);
                        break;
                }

                previousNode.SetNext(currentNode);
                currentNode.SetPrevious(previousNode);
                previousNode = currentNode;
            }
        }

        return nodeGrid;
    }
}