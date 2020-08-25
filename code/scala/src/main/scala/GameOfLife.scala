import scalaj.http.{Http, HttpOptions}


object GameOfLife{
  def rules(grid): Int = {
    var count_neighbors = 0

    var i = cell._1
    var j = cell._2

    var neighbors = Array((i+1,j),(i-1,j),(i,j+1),(i,j-1),(i-1,j-1),(i+1,j+1),(i-1,j+1),(i+1,j-1))

    for (i <- 0 until neighbors.length){
      var x = neighbors(i)._1
      var y = neighbors(i)._2

      if(grid(x)(y) == 1)
        count_neighbors+=1
    }


    if(count_neighbors<=1 && grid(i)(j)==1)
      return 0
    else if(count_neighbors>=4 && count_neighbors<=8 && grid(i)(j)==1)
      return 0
    else if(count_neighbors>=2 && count_neighbors<=3 && grid(i)(j)==1)
      return 1
    else
      return 0
  }

  def display(grid): Unit ={
    // Display the Board
    for( i <- 0 to 7) {
      for (j <- 0 to 7)
        print(grid(i)(j) + " ")
      print("\n")
    }
    print("\n")
  }


  def postResults(grid): Unit ={

    var result = ""

    // Display the Board
    for( i <- 0 to 7) {
      for (j <- 0 to 7)
        result += (grid(i)(j)).toString() + ","
      result += ";"
    }

    //val rest_result = Http("https://hot-octopus-15.serverless.social/CheckGrid").postData("{"grid: + result +  ""","language":"scala"}""")
      //.header("Content-Type", "application/json")
      //.header("Charset", "UTF-8")
      //.option(HttpOptions.readTimeout(10000)).asString
    print(result)
    Http("https://hot-octopus-15.serverless.social/CheckGrid").postForm(Seq("grid" -> result, "language" -> "scala")).asString
  }

  def main(args) {
    var grid = Array.ofDim[Int](8,8)

    var i = 0
    var j = 0
    var steps = 0

    // Initialize the grid
    for( i <- 0 to 7)
      for( j <- 0 to 7)
        grid(i)(j) = 0

    grid(1)(0) = 1
    grid(2)(1) = 1
    grid(2)(2) = 1
    grid(1)(2) = 1
    grid(0)(2) = 1

    //Sdisplay(grid)



    while(steps<15) {
      var new_grid = Array.ofDim[Int](8, 8)

      for (i <- 0 to 7)
        for (j <- 0 to 7)
          new_grid(i)(j) = rules(grid, (i, j))

      //display(new_grid)

      grid = new_grid
      steps+=1
    }

    postResults(grid)
  }
}
