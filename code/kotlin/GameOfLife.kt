fun rules(grid, row, col): Int{
    var count_neighbors = 0

    var i = row
    var j = col 

    val neighbors = arrayOf(arrayOf(i+1,j),arrayOf(i-1,j),arrayOf(i,j+1),arrayOf(i,j-1),arrayOf(i-1,j-1),arrayOf(i+1,j+1),arrayOf(i-1,j+1),arrayOf(i+1,j-1))
    
    for (pair in 0..7) { 
      var x = neighbors[pair][0]
      var y = neighbors[pair][1]

      if(grid[x][y] == 1){
          count_neighbors += 1
      }
        
    }   
    if(count_neighbors<=1 && grid[i][j]==1)
      return 0
    else if(count_neighbors>=4 && count_neighbors<=8 && grid[i][j]==1)
      return 0
    else if(count_neighbors>=2 && count_neighbors<=3 && grid[i][j]==1)
      return 1
    else
      return 0
}

// to display the grid
fun display(grid){

        for (arr in grid){
            
            for (value in arr){
                print(value)
            }
            println()
        }
        println()
        
    }
    
fun main(args){
        // Initialize the grid to 0 
        var grid = Array(8) {Array(8) {0} }
        var steps = 0
        
        grid[1][0] = 1
        grid[2][1] = 1
        grid[2][2] = 1
        grid[1][2] = 1
        grid[0][2] = 1

        print("Initial State")
        display(grid)


        while(steps < 10){
          var new_grid = Array(8) {Array(8) {0} }

      
          for( i in 0..grid.size-1){
              for( j in 0..grid[0].size-1)
              {
                  new_grid[i][j] = rules(grid,i,j)
              }
          }
          
          display(new_grid)
          grid = new_grid
          steps+=1
      }
    }
