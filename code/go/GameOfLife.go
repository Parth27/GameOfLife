package main
import "strconv"
import "fmt"

func rules(grid [8][8]int, row int, col int) int {
    var count_neighbors = 0
    neighbors := [...][2]int{{row+1,col},{row-1,col},{row,col+1},{row,col-1},{row-1,col-1},{row+1,col+1},{row-1,col+1},{row+1,col-1}}
    for i := 0; i < len(neighbors); i++ {
        var current_row int = neighbors[i][0]
        var current_column int = neighbors[i][1]
        if current_row < 0 || current_column < 0 || current_row >= len(grid) || current_column >= len(grid[0]){
            continue;
        } else if grid[current_row][current_column] == 1 {
            count_neighbors+=1
        }
    }
    if count_neighbors<=1 && grid[row][col]==1{
        return 0
    } else if count_neighbors>=4 && count_neighbors<=8 && grid[row][col]==1{
        return 0
    } else if count_neighbors>=2 && count_neighbors<=3 && grid[row][col]==1{
        return 1
    } else if count_neighbors==3 && grid[row][col]==0 {
    return 1
    } else{
        return 0
    }
}

func display(grid [8][8]int) {
    for i := 0; i < 8 ; i ++ {
        for j := 0; j < 8 ; j ++ {
            fmt.Print(strconv.Itoa(grid[i][j]) + " ")
        }
        fmt.Println("")
    }
    fmt.Println("\n")
}

func main() {
    generations := 3
    var grid [8][8]int
    grid[3][3] = 1
    grid[3][4] = 1
    grid[3][5] = 1
    display(grid)
    for k := 0 ; k < generations; k ++ {
        var new_grid [8][8]int
        for i := 0 ; i < 8 ; i++ {
            for j:= 0 ; j < 8; j ++{
                new_grid[i][j] = rules(grid, i, j)
            }
        }
        display(new_grid)
        grid = new_grid
    }
}