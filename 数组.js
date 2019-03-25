const  arr = ['a','a','a','b','a','c'];
var desc ;
const newArr = arr.map(item=>{
    console.log(desc);
    let a = desc===item ? ++desc : 1;
    console.log(a);
    desc= item;
    return a;
});
//var b = arr.map((v,i,arr)=>i+1-arr.slice(i).indexOf(v)+i-arr[i]);
/*for (let i=0;i<arr.length;i++){
    if(i>0){
        arr[i-1]
    }
}*/
console.log(newArr);