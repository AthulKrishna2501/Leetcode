function sortSentence(s) {
   split=s.split(" ")
   const arr=[]
   for (i=0;i<split.length;i++){
    let num=Number(split[i][split[i].length-1])
    arr[num]= split[i].slice(0,-1)
   }
console.log(arr)
let result = arr.join(" ").trim()
console.log(result)
return result
}