func calPoints(operations []string) int {
    points:=[]int{}
    sum:=0

    for _, op:= range operations{
        switch op{
            case "+":
            points=append(points,points[len(points)-1]+points[len(points)-2])

            case "D":
            points=append(points,2*points[len(points)-1])

            case "C":
            points=points[:len(points)-1]

            default:
            num,_:= strconv.Atoi(op)
            points=append(points,num)

        }
    }

    for _, num:= range points{
        sum+=num
    }

    return sum
}